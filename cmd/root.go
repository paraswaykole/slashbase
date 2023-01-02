package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/chroma/quick"
	"github.com/gohxs/readline"
	qemodels "github.com/slashbaseide/slashbase/pkg/queryengines/models"
	"github.com/spf13/cobra"
)

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
	Short: "slashbase is a cli & an api server for quering databases",
	Long:  `slashbase is a cli & an api server for quering databases`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Connect to Slashbase IDE at https://app.slashbase.com")
		fmt.Println("Type 'help' for more info on cli.")
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
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
