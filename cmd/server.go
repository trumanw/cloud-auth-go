package cmd

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	srv "github.com/trumanw/cloud-auth-go/server"
)

// Flags can be setup by command or os.ENV
var host string
var port int
var etcdnsForServers []string

// Add a cobra command to maintain all the gRPC server
func init() {
	// Init flags
	serverCmd.Flags().StringVarP(&host, "host", "", "", "host of the gRPC server.")
	serverCmd.Flags().IntVarP(&port, "port", "", 9090, "port of the gRPC server.")
	serverCmd.Flags().StringArrayVar(&etcdnsForServers, "etcdns", []string{"http://localhost:2379"}, "endpoints of etcd cluster.")

	// Add commands to root cmd
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

		if err := srv.Run(host, port, etcdnsForServers); err != nil {
			glog.Fatal(err)
		}
	},
}
