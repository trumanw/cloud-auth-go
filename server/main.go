package server

import (
	"net"

	gw "github.com/trumanw/cloud-auth-go/pb"
	it "github.com/trumanw/cloud-auth-go/server/unary"
	chain "github.com/mwitkow/go-grpc-middleware"

	"google.golang.org/grpc"
)

func Run() error {
    l ,err := net.Listen("tcp", ":9090")
    if err != nil {
        return err
    }

	// add the handlers as a server option
	unaryChain := chain.ChainUnaryServer(it.IdemUnary, it.BasicAuthUnary)
    s := grpc.NewServer(grpc.UnaryInterceptor(unaryChain))
    gw.RegisterCilentCredentialsServiceServer(s, newClientCredentialsServer())

    s.Serve(l)
    return nil
}
