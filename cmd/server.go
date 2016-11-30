package cmd

import (
    "fmt"
    "flag"
    "net/http"

    "github.com/golang/glog"
    "golang.org/x/net/context"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "google.golang.org/grpc"

    "github.com/spf13/cobra"

    gw "github.com/trumanw/cloud-auth-go/pb"
)

func init() {
    RootCmd.AddCommand(serverCmd)
}

var (
    ccEndpoint = flag.String("client_credentials_endpoint", "localhost:9090", "endpoint of client crendentials")
)

func run() error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}
    err := gw.RegisterCilentCredentialsServiceHandlerFromEndpoint(ctx, mux, *ccEndpoint, opts)
    if err != nil {
        return err
    }

    http.ListenAndServe(":8080", mux)
    return nil
}

var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Start server gRPC instance",
    Long:  `Launch the gRPC server to receive rpc connections from gateway or other endpoints`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Server starting...")

        flag.Parse()
        defer glog.Flush()

        if err := run(); err != nil {
            glog.Fatal(err)
        }
    },
}
