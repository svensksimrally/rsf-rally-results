package cmd

import (
	"github.com/caarlos0/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tobier/rsf-rally-results/export"
)

var (
	rootCmd = &cobra.Command{
		Use:           "rsf-rally-results",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			log.SetLevelFromString(viper.GetString("level"))

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info("Welcome to RSF Rally Results!")

			return export.Do()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().String("level", "INFO", "log level")
	//rootCmd.PersistentFlags().String("series", "series/", "where series are fetched from")
	rootCmd.PersistentFlags().String("output", "dist/", "where to publish the resulting files")

	viper.BindPFlag("level", rootCmd.PersistentFlags().Lookup("level"))
	//viper.BindPFlag("series", rootCmd.PersistentFlags().Lookup("series"))
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
}

func Execute() error {
	return rootCmd.Execute()
}
