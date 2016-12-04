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

    opts := []grpc.DialOption{grpc.WithInsecure()}
    err := pb.RegisterCilentCredentialsServiceHandlerFromEndpoint(ctx, mux, *ccEndpoint, opts)
    if err != nil {
        return err
    }

    http.ListenAndServe(":8080", n)
    return nil
}
