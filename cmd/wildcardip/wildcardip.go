package wildcardip

import (
	"github.com/codysk/wildcard-ip/pkg/common"
	"github.com/codysk/wildcard-ip/pkg/wildcardserver"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "wildcard-ip",
	Short: "a simple wildcard dns",
	Run: func(cmd *cobra.Command, args []string) {

		config := common.Config()
		if false == config.TCPServerConfig.Enabled &&
			false == config.UDPServerConfig.Enabled {
			config.TCPServerConfig.Enabled = true
			config.UDPServerConfig.Enabled = true
		}

		exitChan := make(chan bool, 0)
		if config.TCPServerConfig.Enabled {
			go func() {
				log.Infof("Start and listening tcp dns server at %s.", common.Config().TCPServerConfig.Addr)
				if err := wildcardserver.NewServer("tcp", config.TCPServerConfig.Addr).ListenAndServe(); err != nil {
					log.Errorln(err)
					exitChan <- true
				}
			}()
		}

		if config.UDPServerConfig.Enabled {
			go func() {
				log.Infof("Start and listening udp dns server at %s.", common.Config().UDPServerConfig.Addr)
				if err := wildcardserver.NewServer("udp", config.UDPServerConfig.Addr).ListenAndServe(); err != nil {
					log.Errorln(err)
					exitChan <- true
				}
			}()
		}

		<-exitChan
		log.Infoln("wildcard dns server exited.")
	},
}

func init() {
	Command.Flags().StringVar(
		&common.Config().TCPServerConfig.Addr,
		"tcp-addr",
		":53",
		"nameserver TCP binding address",
	)
	Command.Flags().StringVar(
		&common.Config().UDPServerConfig.Addr,
		"udp-addr",
		":53",
		"nameserver UDP binding address",
	)

	Command.Flags().BoolVarP(
		&common.Config().TCPServerConfig.Enabled,
		"tcp",
		"t",
		false,
		"enable tcp nameserver "+
			"(If neither tcp server nor udp server is enabled, both of them will be enabled)",
	)
	Command.Flags().BoolVarP(
		&common.Config().UDPServerConfig.Enabled,
		"udp",
		"u",
		false,
		"enable udp nameserver "+
			"(If neither tcp server nor udp server is enabled, both of them will be enabled)",
	)
}
