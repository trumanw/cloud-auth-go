package gateway

import (
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/rs/cors"
	hnd "github.com/trumanw/cloud-auth-go/gateway/handler"
	pb "github.com/trumanw/cloud-auth-go/pb"
	ng "github.com/urfave/negroni"

	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
)

func Run(etcdns []string) error {
	fmt.Printf("Gateway etcd nodes: %v ...\n", etcdns)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	// Add middlewares
	n := ng.New()
	n.Use(hnd.NewLogger())
	n.Use(cors.New(cors.Options{}))
	n.Use(hnd.NewIdempotentHandler())
	n.Use(hnd.NewCacheHandler())
	n.UseHandler(mux)

	// resolve connections through etcd
	cli, cerr := clientv3.NewFromURL("http://localhost:2379")
	if cerr != nil {
		return cerr
	}
	defer cli.Close()
	r := &etcdnaming.GRPCResolver{Client: cli}
	b := grpc.RoundRobin(r)
	conn, gerr := grpc.Dial("cloud-auth-go", grpc.WithInsecure(), grpc.WithBalancer(b))
	if gerr != nil {
		return gerr
	}

	// register client with connection
	err := pb.RegisterCilentCredentialsServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	http.ListenAndServe(":8080", n)
	return nil
}
