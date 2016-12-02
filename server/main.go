package server

import (
	"net"

	gw "github.com/trumanw/cloud-auth-go/pb"
	it "github.com/trumanw/cloud-auth-go/interceptor"
	"google.golang.org/grpc"
)

func Run() error {
    l ,err := net.Listen("tcp", ":9090")
    if err != nil {
        return err
    }

	var opts []grpc.ServerOption
	// add the handlers as a server option
	opts = append(opts, grpc.UnaryInterceptor(it.IdemUnaryInterceptor))
    s := grpc.NewServer(opts...)
    gw.RegisterCilentCredentialsServiceServer(s, newClientCredentialsServer())

    s.Serve(l)
    return nil
}
