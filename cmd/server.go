package cmd

import (
    "fmt"
    
    "github.com/spf13/cobra"
)

func init() {
    RootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Start server gRPC instance",
    Long:  `Launch the gRPC server to receive rpc connections from gateway or other endpoints`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Server starting...")
    },
}
