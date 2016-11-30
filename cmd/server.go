package cmd

import (
    "fmt"
    "flag"

    "github.com/golang/glog"
    "github.com/spf13/cobra"

    srv "github.com/trumanw/cloud-auth-go/server"
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

        if err := srv.Run(); err != nil {
            glog.Fatal(err)
        }
    },
}
