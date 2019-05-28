package app

import (
	"errors"
	"github.com/clearcodecn/app/binddata"
	"github.com/spf13/cobra"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var (
	Gopath bool
)

func Create(cmd *cobra.Command, args []string) error {

	if len(args) == 0 {
		return errors.New(`no name provide`)
	}

	var appPath = ""
	if Gopath {
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			home, _ := os.UserHomeDir()
			if home != "" {
				appPath = filepath.Join(home, "go", "src")
			}
		} else {
			appPath = filepath.Join(gopath, "src")
		}
	} else {
		appPath = args[0]
	}

	os.MkdirAll(appPath, 0777)
	for _, v := range binddata.AssetNames() {
		b, err := binddata.Asset(v)
		if err != nil {
			return err
		}
		name := strings.Replace(v, "apptpl", "", -1)
		file := filepath.Join(appPath, name)

		tpl, err := template.New(name).Parse(string(b))
		if err != nil {
			return err
		}

		os.MkdirAll(filepath.Dir(file),0755)

		f, err := os.Create(file)
		if err != nil {
			return err
		}
		err = tpl.Execute(f, map[string]interface{}{
			"pkgName": args[0],
		})

		if err != nil {
			return err
		}
	}

	return nil
}
