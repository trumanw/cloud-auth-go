package gateway

import (
    "flag"
    "net/http"

    "golang.org/x/net/context"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "google.golang.org/grpc"

    ng "github.com/urfave/negroni"
    hnd "github.com/trumanw/cloud-auth-go/gateway/handler"
    pb "github.com/trumanw/cloud-auth-go/pb"
    "github.com/rs/cors"

    "github.com/coreos/etcd/clientv3"
    etcdnaming "github.com/coreos/etcd/clientv3/naming"
)

// gRPC gateway registration
var (
    ccEndpoint = flag.String("client_credentials_endpoint", "localhost:9090", "endpoint of client crendentials")
)

func Run() error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()
    // Add middlewares
    n := ng.New()
    n.Use(hnd.NewLogger())
    n.Use(hnd.NewIdempotent())
    n.Use(cors.New(cors.Options{}))
    n.UseHandler(mux)

    // opts := []grpc.DialOption{grpc.WithInsecure()}
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

    // err := pb.RegisterCilentCredentialsServiceHandlerFromEndpoint(ctx, mux, *ccEndpoint, opts)
    // register client with connection
    err := pb.RegisterCilentCredentialsServiceHandler(ctx, mux, conn)
    if err != nil {
        return err
    }

    http.ListenAndServe(":8080", n)
    return nil
}
