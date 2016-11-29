package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

func init() {
    RootCmd.AddCommand(gatewayCmd)
}

var gatewayCmd = &cobra.Command{
    Use:   "gateway",
    Short: "Start gRPC gateway instance",
    Long:  `The gateway can wrap the rpc requests to be http compatible.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Gateway starting...")
    },
}
