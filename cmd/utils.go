package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ReadArgs(cmd *cobra.Command, arguments any) error {
	var err error

	err = viper.BindPFlags(cmd.Flags())
	if err != nil {
		return err
	}

	err = viper.MergeInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(arguments)
	if err != nil {
		return err
	}

	return nil
}
