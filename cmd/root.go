package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/chroma/quick"
	"github.com/gohxs/readline"
	"github.com/slashbaseide/slashbase/internal/config"
	"github.com/slashbaseide/slashbase/internal/models"
	"github.com/spf13/cobra"
)

func display(input string) string {
	if cliApp.CurrentDB == nil {
		return input
	} else if cliApp.CurrentDB.Type == models.DBTYPE_POSTGRES {
		buf := bytes.NewBuffer([]byte{})
		err := quick.Highlight(buf, input, "postgres", "terminal16", "monokai")
		if err != nil {
			return input
		}
		return buf.String()
	} else if cliApp.CurrentDB.Type == models.DBTYPE_MONGO {
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
	Short: "slashbase is a cli & an api server for quering databases",
	Long:  `slashbase is a cli & an api server for quering databases`,
	Run: func(cmd *cobra.Command, args []string) {
		term, err := readline.NewEx(&readline.Config{
			Prompt: "slashbase > ",
			Output: display,
		})
		if err != nil {
			log.Fatal(err)
		}
		for {
			line, err := term.Readline()
			if err != nil {
				log.Fatal(err)
			}
			handleCmd(line)
			if cliApp.CurrentDB == nil {
				term.SetPrompt("slashbase > ")
			} else {
				term.SetPrompt(fmt.Sprintf("%s > ", cliApp.CurrentDB.Name))
			}
		}
	},
}

func Execute() {
	if config.GetConfig().BuildName == config.BUILD_DOCKER_PROD {
		return
	}
	fmt.Println("Type 'help' for more info on cli.")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
