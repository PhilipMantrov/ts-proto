package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"log"
	"math/rand"
	"net"

	"golang.org/x/net/context"

	rpx "./generated/lib/rpx"
)

var (
	enableTls       = flag.Bool("enable_tls", false, "Use TLS - required for HTTP2.")
	tlsCertFilePath = flag.String("tls_cert_file", "../../misc/localhost.crt", "Path to the CRT/PEM file.")
	tlsKeyFilePath  = flag.String("tls_key_file", "../../misc/localhost.key", "Path to the private key file.")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("create server at localhost:9090")
	server := grpc.NewServer()
	rpx.RegisterDashStateServer(server, &stateService{})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type stateService struct{}

func (s *stateService) UserSettings(ctx context.Context, in *rpx.Empty) (*rpx.DashUserSettingsState, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	urls := rpx.DashUserSettingsState_URLs{
		ConnectGoogle: "http://google.com",
		ConnectGithub: "http://github.com",
	}

	flashes := []*rpx.DashFlash{
		&rpx.DashFlash{
			Msg:  "flash1",
			Type: rpx.DashFlash_Warn,
		},
		&rpx.DashFlash{
			Msg:  "flash2",
			Type: rpx.DashFlash_Success,
		},
	}

	val := rpx.DashUserSettingsState{
		Email:   "test-email@example.com",
		LongNumber: 1,
        SmallNumber: 2,
		Urls:    &urls,
		Flashes: flashes,
	}

	return &val, nil
}

func (s *stateService) SendUserSetting(ctx context.Context, in *rpx.DashUserSettingsState) (*rpx.DashUserSettingsState, error) {
    log.Println(in);
    return in, nil;
}

func (s *stateService) ActiveUserSettingsStream(in *rpx.Empty, stream rpx.DashState_ActiveUserSettingsStreamServer) error {
	urls := rpx.DashUserSettingsState_URLs{
		ConnectGoogle: "http://google.com",
		ConnectGithub: "http://github.com",
	}

	flashes := []*rpx.DashFlash{
		&rpx.DashFlash{
			Msg:  "flash1",
			Type: rpx.DashFlash_Warn,
		},
		&rpx.DashFlash{
			Msg:  "flash2",
			Type: rpx.DashFlash_Success,
		},
	}

	val := rpx.DashUserSettingsState{
		Email:   "test-email@example.com",
		Urls:    &urls,
		LongNumber: 1,
        SmallNumber: 2,
		Flashes: flashes,
	}

	val_second := rpx.DashUserSettingsState{
		Email:   "test2-email@example.com",
		Urls:    &urls,
        LongNumber: 1,
        SmallNumber: 2,
		Flashes: flashes,
	}

	val_third := rpx.DashUserSettingsState{
		Email:   "test3-email@example.com",
		Urls:    &urls,
		LongNumber: 1,
        SmallNumber: 2,
		Flashes: flashes,
	}

	stream.Send(&val)
	stream.Send(&val_second)
	stream.Send(&val_third)
	return nil
}

type credsService struct{}

var creds = map[string]rpx.DashCred{}

func (s *credsService) Create(c context.Context, in *rpx.DashAPICredsCreateReq) (*rpx.DashCred, error) {
	cred := rpx.DashCred{
		Description: in.Description,
		Metadata:    in.Metadata,
		Token:       "token123",
		Id:          fmt.Sprintf("id-%d", rand.Int()),
	}

	creds[cred.Id] = cred

	return &cred, nil
}

func (s *credsService) Update(c context.Context, in *rpx.DashAPICredsUpdateReq) (*rpx.DashCred, error) {

	fmt.Println("Update", in.CredSid)
	return nil, grpc.Errorf(codes.NotFound, "not found")
}

func (s *credsService) Delete(c context.Context, in *rpx.DashAPICredsDeleteReq) (*rpx.DashCred, error) {
	grpclog.Printf("DELETE ID: %v", in.Id)

	cred, ok := creds[in.Id]

	grpclog.Printf("cred: %v", creds)

	if !ok {
		return nil, grpc.Errorf(codes.NotFound, "not found")
	}
	delete(creds, in.Id)
	return &cred, nil
}
