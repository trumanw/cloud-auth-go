package cmd

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	gw "github.com/trumanw/cloud-auth-go/gateway"
)

var etcdnsForGateway []string

// Add a cobra command to maintain all the gateway registration
func init() {
	// Init flags
	gatewayCmd.Flags().StringArrayVar(&etcdnsForGateway, "etcdns", []string{"http://localhost:2379"}, "endpoints of etcd cluster.")

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

		if err := gw.Run(etcdnsForGateway); err != nil {
			glog.Fatal(err)
		}
	},
}
