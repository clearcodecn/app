package main

import (
	"github.com/clearcodecn/app"
	"github.com/clearcodecn/app/cmd/create"
	"github.com/spf13/cobra"
	"log"
)

var (
	rootCmd = &cobra.Command{
		Use:                        "cli",
		Aliases:                    nil,
		SuggestFor:                 nil,
	}
)

func init()  {
	rootCmd.Flags().BoolVar(&app.IsDebug,"d,debug",true,"debug mode")
	rootCmd.AddCommand(
		create.CreateCmd,
		)
}

func main()  {
	if err := rootCmd.Execute() ; err != nil {
		log.Fatal(err)
	}
}
