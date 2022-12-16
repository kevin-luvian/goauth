package cmd

import (
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/spf13/cobra"
)

var InspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp"},
	Short:   "Inspects a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.Infoln("args", args)

		// i := args[0]
		// res, kind := stringer.Inspect(i, false)

		// pluralS := "s"
		// if res == 1 {
		// 	pluralS = ""
		// }
		// fmt.Printf("'%s' has a %d %s%s.\n", i, res, kind, pluralS)
	},
}
