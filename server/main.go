package server

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	gw "github.com/trumanw/cloud-auth-go/pb"
	// it "github.com/trumanw/cloud-auth-go/server/unary"

	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
)

// Run is the main function to lanuch the gRPC servers.
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

	// initialize your gRPC server's interceptor.
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	// register your gRPC service implementations.
	gw.RegisterCilentCredentialsServiceServer(s, newClientCredentialsServer())
	gw.RegisterTokenServiceServer(s, newTokenServer())
	// After all your registrations, make sure all of the Prometheus metrics are initialized.
	grpc_prometheus.Register(s)
	// Register Prometheus metrics handler.
	http.Handle("/metrics", prometheus.Handler())

	s.Serve(l)
	return nil
}
