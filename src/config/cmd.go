package config

import (
	"github.com/biFebriansyah/gobackend/src/database/orm"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "simple backend with golang",
}

func init() {
	initCommand.AddCommand(ServeCmd)
	initCommand.AddCommand(orm.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
