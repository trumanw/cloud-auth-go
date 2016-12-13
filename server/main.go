package server

import (
	"fmt"
	"net"
	"strconv"
	"time"

	// chain "github.com/mwitkow/go-grpc-middleware"
	gw "github.com/trumanw/cloud-auth-go/pb"
	// it "github.com/trumanw/cloud-auth-go/server/unary"

	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
)

func Run(host string, port int, etcdns []string) error {
	fmt.Println("Server etcd nodes: ")
	fmt.Println(etcdns)
	addr := host + ":" + strconv.Itoa(port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// register the service endpoint
	// cli, cerr := clientv3.NewFromURL("http://localhost:2379")
	cli, cerr := clientv3.New(clientv3.Config{
		Endpoints:   etcdns,
		DialTimeout: 5 * time.Second,
	})

	if cerr != nil {
		return cerr
	}
	defer cli.Close()

	r := &etcdnaming.GRPCResolver{Client: cli}
	r.Update(context.TODO(), "cloud-auth-go", naming.Update{Op: naming.Add, Addr: addr, Metadata: "..."})

	// add the handlers as a server option
	// unaryChain := chain.ChainUnaryServer(it.BasicAuthUnary)
	// s := grpc.NewServer(grpc.UnaryInterceptor(unaryChain))
	s := grpc.NewServer()
	gw.RegisterCilentCredentialsServiceServer(s, newClientCredentialsServer())
	gw.RegisterTokenServiceServer(s, newTokenServer())

	s.Serve(l)
	return nil
}
