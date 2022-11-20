package prepare

import (
	"fmt"

	"github.com/charlesbourget/aoc/cmd"
	"github.com/charlesbourget/aoc/pkg/preparer"
	"github.com/charlesbourget/aoc/pkg/utils"
	"github.com/spf13/cobra"
)

type Arguments struct {
	Day     int            `mapstructure:"day"`
	Year    int            `mapstructure:"year"`
	Today   bool           `mapstructure:"today"`
	Lang    utils.Language `mapstructure:"lang"`
	Session string         `mapstructure:"session"`
}

func Cmd() *cobra.Command {
	prepareCmd := cobra.Command{
		Use:   "prepare",
		Short: "Prepare the environment for a new day",
		RunE:  run,
	}

	var language = utils.Go

	prepareCmd.Flags().IntP("day", "d", 0, "Day to prepare")
	prepareCmd.Flags().IntP("year", "y", 0, "Year to prepare")
	prepareCmd.Flags().BoolP("today", "t", false, "Prepare today's day")
	prepareCmd.Flags().Var(&language, "lang", "Language to use")

	return &prepareCmd
}

func run(cmdr *cobra.Command, argsv []string) error {
	args := Arguments{}
	if err := cmd.ReadArgs(cmdr, &args); err != nil {
		return err
	}

	var day = args.Day
	var year = args.Year

	if args.Today {
		day = utils.GetCurrentDay()
		year = utils.GetCurrentYear()
	}

	if !utils.ValidateDayYearInput(day, year) {
		return fmt.Errorf("invalid day or year")
	}

	preparer := preparer.Create(day, year, args.Lang, args.Session)
	if err := preparer.PrepareDay(); err != nil {
		return err
	}

	return nil
}
