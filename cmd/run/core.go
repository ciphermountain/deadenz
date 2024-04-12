package run

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	proto "github.com/ciphermountain/deadenz/pkg/proto/core"
	"github.com/ciphermountain/deadenz/pkg/service/core"
	"github.com/ciphermountain/deadenz/pkg/service/multiverse"
)

func init() {
	runCore.Flags().BoolVar(&withMultiverse, "with-multiverse", false, "optionally connect to multiverse service")
	runCore.Flags().StringVar(&multiverseHost, "multiverse-host", "127.0.0.1:8080", "host address to multiverse service")
}

var (
	withMultiverse bool
	multiverseHost string

	runCore = &cobra.Command{
		Use:   "core",
		Short: "Core system service",
		Long:  "Core system service",
		Run: func(cmd *cobra.Command, args []string) {
			var (
				client *multiverse.Client
				err    error
			)

			if withMultiverse {
				if client, err = multiverse.NewClient(multiverseHost); err != nil {
					fmt.Fprintln(cmd.ErrOrStderr(), "could not connect to multiverse service")
					os.Exit(1)
				}
			}

			log.Println("starting core service")

			startServer(host, port, cmd.ErrOrStderr(), func(server grpc.ServiceRegistrar) {
				proto.RegisterDeadenzServer(server, core.NewServer(client))
			})
		},
	}
)
