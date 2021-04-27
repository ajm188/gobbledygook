package main

import (
	"log"
	"math/rand"
	"os"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/ajm188/gobbledygook"
)

var (
	fm      template.FuncMap
	tmplStr string
	root    = &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			tmpl, err := template.New("").Funcs(map[string]interface{}{
				"iterate": func(count int) []int {
					ints := make([]int, count)
					for i := range ints {
						ints[i] = i
					}

					return ints
				},
				"randint": rand.Intn,
				"word":    gobbledygook.Word,
				"words":   gobbledygook.Words,
			}).Parse(tmplStr)
			if err != nil {
				return err
			}

			if err := gobbledygook.InitWords(); err != nil {
				return err
			}

			cmd.SilenceUsage = true
			return tmpl.Execute(os.Stdout, nil)
		},
	}
)

func init() {
	root.Flags().StringVar(&tmplStr, "template", "", "")

	root.MarkFlagRequired("template")
}

func main() {
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
