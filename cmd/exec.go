package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/seblegall/sigstat/pkg/client/http"
	"github.com/seblegall/sigstat/pkg/sigstat"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute the command passed as argument and send monitoring data to the sigstat server",
	Long: `The "exec" command execute the cli command passed as argument and send informations about the underlying process regularly.
	Simple command can be passed as easily as :

	$ sigstat exec ls

	For more complexe command, you just need to add quotes arround the cli command :

	$ sigstat exec "ls -la | wc -l"

	`,
	Run: func(cmd *cobra.Command, args []string) {

		//get the current working dir
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal("Error when getting the working dir : ", err)
		}

		command := sigstat.Command{
			Command: strings.Join(args, " "),
			Path:    wd,
			Group:   "test",
		}
		c := http.NewClient()

		id, err := c.CommandService().CreateCommand(command)
		if err != nil {
			log.Println("test")
			log.Fatal(err)
			log.Println("test2")
		}

		fmt.Println("The return id is : ", id)
		command.ID = id

		//Execute the command
		command.Exec(c)

		//Print the exit code at the end
		fmt.Println("Exit Code : ", command.ExitCode)

	},
}

func init() {
	RootCmd.AddCommand(execCmd)
}
