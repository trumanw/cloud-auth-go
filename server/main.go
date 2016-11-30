package server

import (
	"net"

	gw "github.com/trumanw/cloud-auth-go/pb"
	"google.golang.org/grpc"
)

func Run() error {
    l ,err := net.Listen("tcp", ":9090")
    if err != nil {
        return err
    }
    s := grpc.NewServer()
    gw.RegisterCilentCredentialsServiceServer(s, newClientCredentialsServer())

    s.Serve(l)
    return nil
}
