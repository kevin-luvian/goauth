package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/spf13/cobra"
)

func MigrateUp(m *migrate.Migrate) *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Migrate up to latest",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if err := m.Up(); err != nil {
				logging.Errorln(err)
			}

			logging.Infoln("migrated up")
		},
	}
}

func MigrateDown(m *migrate.Migrate) *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "Migrate down to earliest",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if err := m.Down(); err != nil {
				logging.Errorln(err)
			}

			logging.Infoln("migrated down")
		},
	}
}
