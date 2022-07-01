// Code generated by protoc-gen-isle-rpc. DO NOT EDIT.
//
// Source: timesync.proto

package timesync

import (
	context "context"
	errors "errors"
	pbstream "github.com/lab47/isle/pkg/pbstream"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = pbstream.IsAtLeastVersion0_1_0

const (
	// TimeSyncName is the fully-qualified name of the TimeSync service.
	TimeSyncName = "dev.lab47.isle.timesync.TimeSync"
)

// PBSTimeSyncClient is a client for the dev.lab47.isle.timesync.TimeSync service.
type PBSTimeSyncClient interface {
	TimeSync(context.Context) *pbstream.BidiStreamForClient[NTPTimePacket, NTPTimePacket]
}

// PBSNewTimeSyncClient constructs a client for the dev.lab47.isle.timesync.TimeSync service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func PBSNewTimeSyncClient(client pbstream.Client, opts ...pbstream.ClientOption) PBSTimeSyncClient {
	return &pBSTimeSyncClient{
		timeSync: pbstream.NewEndpoint[NTPTimePacket, NTPTimePacket](
			client,
			"/dev.lab47.isle.timesync.TimeSync/TimeSync",
		),
	}
}

// pBSTimeSyncClient implements PBSTimeSyncClient.
type pBSTimeSyncClient struct {
	timeSync *pbstream.Endpoint[NTPTimePacket, NTPTimePacket]
}

// TimeSync calls dev.lab47.isle.timesync.TimeSync.TimeSync.
func (c *pBSTimeSyncClient) TimeSync(ctx context.Context) *pbstream.BidiStreamForClient[NTPTimePacket, NTPTimePacket] {
	return c.timeSync.CallBidiStream(ctx)
}

// PBSTimeSyncHandler is an implementation of the dev.lab47.isle.timesync.TimeSync service.
type PBSTimeSyncHandler interface {
	TimeSync(context.Context, *pbstream.BidiStream[NTPTimePacket, NTPTimePacket]) error
}

// PBSNewTimeSyncHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func PBSNewTimeSyncHandler(svc PBSTimeSyncHandler, opts ...pbstream.HandlerOption) (string, pbstream.Handler) {
	mux := pbstream.NewMux()
	mux.Handle("/dev.lab47.isle.timesync.TimeSync/TimeSync", pbstream.NewBidiStreamHandler(
		"/dev.lab47.isle.timesync.TimeSync/TimeSync",
		svc.TimeSync,
	))
	return "/dev.lab47.isle.timesync.TimeSync/", mux
}

// PBSUnimplementedTimeSyncHandler returns CodeUnimplemented from all methods.
type PBSUnimplementedTimeSyncHandler struct{}

func (PBSUnimplementedTimeSyncHandler) TimeSync(context.Context, *pbstream.BidiStream[NTPTimePacket, NTPTimePacket]) error {
	return pbstream.NewError(pbstream.CodeUnimplemented, errors.New("dev.lab47.isle.timesync.TimeSync.TimeSync is not implemented"))
}
