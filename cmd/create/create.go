package create

import (
	"github.com/clearcodecn/app"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:                        "create",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Create(cmd,args)
	},
}

func init()  {
	CreateCmd.Flags().BoolVar(&app.Gopath,"gopath",false,"create in $GOPATH")
}