package cmd

import (
    "fmt"
    "flag"

    "github.com/golang/glog"
    "github.com/spf13/cobra"

    gw "github.com/trumanw/cloud-auth-go/gateway"
)

// Add a cobra command to maintain all the gateway registration
func init() {
    RootCmd.AddCommand(gatewayCmd)
}

var gatewayCmd = &cobra.Command{
    Use:   "gateway",
    Short: "Start gRPC gateway instance",
    Long:  `The gateway can wrap the rpc requests to be http compatible.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Gateway starting...")

        flag.Parse()
        defer glog.Flush()

        if err := gw.Run(); err != nil {
            glog.Fatal(err)
        }
    },
}
