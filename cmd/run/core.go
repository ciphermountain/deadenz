package run

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	proto "github.com/ciphermountain/deadenz/pkg/proto/core"
	"github.com/ciphermountain/deadenz/pkg/service/core"
)

var (
	runCore = &cobra.Command{
		Use:   "core",
		Short: "Core system service",
		Long:  "Core system service",
		Run: func(cmd *cobra.Command, args []string) {
			startServer(host, port, cmd.OutOrStderr(), func(server grpc.ServiceRegistrar) {
				proto.RegisterDeadenzServer(server, core.NewServer())
			})
		},
	}
)
