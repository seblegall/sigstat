// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/seblegall/sigstat/pkg/http"
	"github.com/seblegall/sigstat/pkg/sigstat"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute the command passed as argument and send monitoring data to the sigstat server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		//get the current working dir
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal("Error when getting the working dir : ", err)
		}

		command := sigstat.Command{
			Command: args,
			Path:    wd,
		}
		c := http.NewClient()

		//Execute the command
		command.Exec(c)

		//Print the exit code and the process status at the end
		fmt.Println(command.ExitCode)
		c.CommandService().UpdateStatus(&command)

	},
}

func init() {
	RootCmd.AddCommand(execCmd)
}
