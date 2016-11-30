package cmd

import (
    "fmt"
    "flag"
    "net"

    "github.com/golang/glog"
    "google.golang.org/grpc"
    "github.com/spf13/cobra"

    sr "github.com/trumanw/cloud-auth-go/server"
    gw "github.com/trumanw/cloud-auth-go/pb"
)

// Add a cobra command to maintain all the gRPC server
func init() {
    RootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Start server gRPC instance",
    Long:  `Launch the gRPC server to receive rpc connections from gateway or other endpoints`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Server starting...")

        flag.Parse()
        defer glog.Flush()

        if err := runServer(); err != nil {
            glog.Fatal(err)
        }
    },
}

// Launch the gRPC servers
func runServer() error {
    l ,err := net.Listen("tcp", ":9090")
    if err != nil {
        return err
    }
    s := grpc.NewServer()
    gw.RegisterCilentCredentialsServiceServer(s, sr.newClientCredentialsServer())

    s.Serve(l)
    return nil
}
