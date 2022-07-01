package testutils

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/lab47/isle/guest"
	"github.com/lab47/isle/guestapi"
	"github.com/stretchr/testify/require"
	"go.etcd.io/bbolt"
)

type SetupData struct {
	ResourceContext *guest.ResourceContext
	Logger          hclog.Logger
	IPNetwork       *guestapi.ResourceId
	Containers      *guest.ContainerManager
}

func Setup(t *testing.T, f func(*SetupData)) {
	log := hclog.New(&hclog.LoggerOptions{
		Name:  "testspawn",
		Level: hclog.Trace,
	})

	home, err := os.UserHomeDir()
	require.NoError(t, err)

	homeTmp := filepath.Join(home, "tmp")
	os.MkdirAll(homeTmp, 0755)

	dir, err := os.MkdirTemp(homeTmp, "cont")
	require.NoError(t, err)

	os.MkdirAll(dir, 0755)

	defer os.RemoveAll(dir)

	tf, err := os.CreateTemp("", "rs")
	require.NoError(t, err)

	defer os.Remove(tf.Name())

	tf.Close()

	db, err := bbolt.Open(tf.Name(), 0644, bbolt.DefaultOptions)
	require.NoError(t, err)

	defer db.Close()

	var r guest.ResourceStorage

	err = r.Init(db)
	require.NoError(t, err)

	top, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx := &guest.ResourceContext{
		Context:         top,
		ResourceStorage: &r,
	}

	runDir, err := os.MkdirTemp("/tmp", "testc")
	require.NoError(t, err)

	defer os.RemoveAll(runDir)

	runcDir := filepath.Join(runDir, "runc")

	var ipm guest.IPNetworkManager

	err = ipm.Init(ctx)
	require.NoError(t, err)

	v6net, err := guest.IPv6Base(guest.NewHexId(6), guest.NewHexId(2))
	require.NoError(t, err)

	ipn := &guestapi.IPNetwork{
		Name:        "default",
		Ipv4Block:   guestapi.MustParseCIDR("10.100.0.0/24"),
		Ipv4Gateway: guestapi.MustParseCIDR("10.100.0.1/32"),
		Ipv6Block:   guestapi.ToIPAddress(v6net),
		Ipv6Gateway: guestapi.FromNetIP(guest.FirstAddress(v6net)),
	}

	resip, err := ipm.Create(ctx, ipn)
	require.NoError(t, err)

	var cm guest.ContainerManager

	layerCache := filepath.Join(homeTmp, "layer-cache")
	os.MkdirAll(layerCache, 0755)

	err = cm.Init(ctx, &guest.ContainerConfig{
		Logger:   log,
		BaseDir:  dir,
		HomeDir:  runDir,
		NodeId:   guest.NewUniqueId(),
		User:     "test",
		RuncRoot: runcDir,
		RunDir:   runDir,

		LayerCacheDir: layerCache,

		BridgeID: 2,

		HelperPath:     "/usr/bin/isle",
		NetworkManager: &ipm,
	})
	require.NoError(t, err)

	defer cm.Close()

	go cm.StartSSHAgent(ctx)

	f(&SetupData{
		ResourceContext: ctx,
		Logger:          log,
		IPNetwork:       resip.Id,
		Containers:      &cm,
	})
}
