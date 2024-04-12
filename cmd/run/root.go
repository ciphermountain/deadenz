package run

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&host, "listen", "l", "", "address for service to listen on")
	RootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "port for service to listen on")

	RootCmd.AddCommand(runCore)
	RootCmd.AddCommand(runMultiverse)
	RootCmd.AddCommand(runClient)
}

var (
	host string
	port int

	RootCmd = &cobra.Command{
		Use:       "run",
		Short:     "Run services for the deadenz game",
		Long:      "Run services for the deadenz game",
		ValidArgs: []string{"core", "multiverse", "client"},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			log.SetOutput(cmd.OutOrStdout())
			log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		},
	}
)

type protoRegisterFunc func(grpc.ServiceRegistrar)

func startServer(host string, port int, writer io.Writer, f protoRegisterFunc) {
	listen := fmt.Sprintf("%s:%d", host, port)

	lis, err := net.Listen("tcp", listen)
	if err != nil {
		fmt.Fprintf(writer, "failed to listen on %s: %s", listen, err.Error())
		os.Exit(1)
	}

	log.Printf("grpc server listening on %s", listen)

	opts := []grpc.ServerOption{
		// grpc.Creds(insecure.NewCredentials()),
	}

	grpcServer := grpc.NewServer(opts...)
	c := make(chan os.Signal, 1)

	f(grpcServer)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Println("received stop notification")

			grpcServer.GracefulStop()
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Fprintf(writer, "server closed with err: %s", err.Error())
		os.Exit(1)
	}
}
