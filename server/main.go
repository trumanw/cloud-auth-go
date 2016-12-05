package server

import (
	"net"
	"strconv"

	gw "github.com/trumanw/cloud-auth-go/pb"
	it "github.com/trumanw/cloud-auth-go/server/unary"
	chain "github.com/mwitkow/go-grpc-middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"golang.org/x/net/context"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
)

func Run(host string, port int) error {
	addr := host + ":" + strconv.Itoa(port)
    l ,err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }

	// register the service endpoint
	cli, cerr := clientv3.NewFromURL("http://localhost:2379")
	if cerr != nil {
		return cerr
	}
	defer cli.Close()

	r := &etcdnaming.GRPCResolver{Client: cli}
	r.Update(context.TODO(), "cloud-auth-go", naming.Update{Op: naming.Add, Addr: addr, Metadata: "..."})

	// add the handlers as a server option
	unaryChain := chain.ChainUnaryServer(it.BasicAuthUnary)
    s := grpc.NewServer(grpc.UnaryInterceptor(unaryChain))
    gw.RegisterCilentCredentialsServiceServer(s, newClientCredentialsServer())

    s.Serve(l)
    return nil
}
