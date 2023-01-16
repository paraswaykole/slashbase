package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alecthomas/chroma/quick"
	"github.com/slashbaseide/slashbase/internal/config"
	"github.com/slashbaseide/slashbase/internal/server"
	"github.com/slashbaseide/slashbase/internal/setup"
	"github.com/slashbaseide/slashbase/internal/tasks"
	"github.com/slashbaseide/slashbase/pkg/queryengines"
	qemodels "github.com/slashbaseide/slashbase/pkg/queryengines/models"
	"github.com/spf13/cobra"
)

var longDescription = `Slashbase is a modern in-browser database IDE & CLI for your dev/data workflows. 
Use Slashbase to connect to your database, browse data and schema, write, 
run and save queries, create charts, right from your browser. 
Connects to Slashbase IDE at https://local.slashbase.com`

func display(input string) string {
	if cliApp.CurrentDB == nil {
		return input
	} else if cliApp.CurrentDB.Type == qemodels.DBTYPE_POSTGRES {
		buf := bytes.NewBuffer([]byte{})
		err := quick.Highlight(buf, input, "postgres", "terminal16", "monokai")
		if err != nil {
			return input
		}
		return buf.String()
	} else if cliApp.CurrentDB.Type == qemodels.DBTYPE_MYSQL {
		buf := bytes.NewBuffer([]byte{})
		err := quick.Highlight(buf, input, "mysql", "terminal16", "monokai")
		if err != nil {
			return input
		}
		return buf.String()
	} else if cliApp.CurrentDB.Type == qemodels.DBTYPE_MONGO {
		buf := bytes.NewBuffer([]byte{})
		err := quick.Highlight(buf, input, "javascript", "terminal16", "monokai")
		if err != nil {
			return input
		}
		return buf.String()
	}
	return input
}

var rootCmd = &cobra.Command{
	Use:   "slashbase",
	Short: "Slashbase is a modern in-browser database IDE & CLI",
	Long:  longDescription,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Access Slashbase IDE at http://localhost:" + config.GetConfig().Port)
		fmt.Println("Type 'help' for more info on cli.")
		setup.SetupApp()
		queryengines.Init()
		tasks.InitCron()
		server.Init()
		startCLI()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Slashbase",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.GetConfig().Version)
	},
}

func Execute() {
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
