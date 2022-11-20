package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charlesbourget/aoc/cmd/prepare"
	"github.com/charlesbourget/aoc/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Errorf("could not get user config dir: %w", err)
		os.Exit(1)
	}

	err = utils.CreateDir(filepath.Join(userConfigDir, "aoc"))
	if err != nil {
		log.Errorf("could not create config dir: %w", err)
		os.Exit(1)
	}

	viper.KeyDelimiter("-")
	viper.SetConfigName("aoc")
	viper.SetConfigType("toml")
	viper.AddConfigPath(filepath.Join(userConfigDir, "aoc"))
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.ReadInConfig()

	rootCmd.AddCommand(prepare.Cmd())
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug mode")
}

type Arguments struct {
	Debug bool `mapstructure:"debug"`
}

var rootCmd = cobra.Command{
	Short: "Advent of Code CLI tool",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetBool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
