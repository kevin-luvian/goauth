package main

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kevin-luvian/goauth/server/cmd/migrator/cmd"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "migrator",
	Short: "migrator - a simple CLI to run database migrations",
	Long: `migrator is a super fancy CLI (kidding) 
		One can use migrator to modify database tables`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	setting.Setup()
}

func main() {
	m, err := migrate.New("file://migrations", setting.Database.LocalURL)
	if err != nil {
		logging.Errorf("Error making migration '%s'", err)
		os.Exit(1)
	}

	rootCmd.AddCommand(cmd.InspectCmd)
	rootCmd.AddCommand(cmd.MigrateUp(m))
	rootCmd.AddCommand(cmd.MigrateDown(m))

	if err := rootCmd.Execute(); err != nil {
		logging.Errorf("Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
