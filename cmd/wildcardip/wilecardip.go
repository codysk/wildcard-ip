package wildcardip

import (
	"github.com/codysk/wildcard-ip/pkg/wildcardserver"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command {
	Use: "wildcard-ip",
	Short: "simple wildcard dns server",
	Run: func(cmd *cobra.Command, args []string) {

		exitChan := make(chan bool, 0)
		go func() {
			log.Infoln("Start and listening tcp dns server.")
			if err := wildcardserver.NewServer("tcp", ":53").ListenAndServe(); err != nil {
				log.Errorln(err)
				exitChan <- true
			}
		}()

		go func() {
			log.Infoln("Start and listening udp dns server.")
			if err := wildcardserver.NewServer("tcp", ":54").ListenAndServe(); err != nil {
				log.Errorln(err)
				exitChan <- true
			}
		}()

		<-exitChan
		log.Infoln("wildcard dns server exited.")
	},
}


