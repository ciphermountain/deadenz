package run

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	proto "github.com/ciphermountain/deadenz/pkg/proto/multiverse"
	"github.com/ciphermountain/deadenz/pkg/service/multiverse"
)

var (
	runMultiverse = &cobra.Command{
		Use:   "multiverse",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			startServer(host, port, cmd.OutOrStderr(), func(server grpc.ServiceRegistrar) {
				proto.RegisterMultiverseServer(server, multiverse.NewMultiverseServer())
			})
		},
	}
)
