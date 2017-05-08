package cmd

import (
	"github.com/seblegall/sigstat/pkg/postgres"
	"github.com/seblegall/sigstat/pkg/server/http"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs an http server that expose a REST API in order to store data send by the sigstat exec command.",
	Long: `The "serve" command launch a web server exposing an REST API used to store data send by the sigstat exec command.

	This API expose some uri used to create new command, send information about a command status (running, stopped, etc.)
	and get all those datas.

	Usage :
	$ sigstat serve

	And then, go to http://localhost:9000 using your prefered REST client.`,
	Run: func(cmd *cobra.Command, args []string) {

		//Create new postgres client. (that does nothing yet)
		c := postgres.NewClient()

		h := http.NewHandler(c)
		s := http.NewServer(h)
		s.Serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

}
