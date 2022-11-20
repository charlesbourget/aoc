package preparer

import (
	"fmt"
	"os"

	"github.com/charlesbourget/aoc/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type Preparer struct {
	Day     int
	Year    int
	Lang    string
	Session string
}

func Create(day int, year int, lang utils.Language, session string) *Preparer {
	return &Preparer{
		Day:     day,
		Year:    year,
		Lang:    lang.String(),
		Session: session,
	}
}

func (p Preparer) PrepareDay() error {
	log.Infoln("Preparing day", p.Day, "for year", p.Year)

	if !utils.ValidateDayYearInput(p.Day, p.Year) {
		return fmt.Errorf("invalid day or year")
	}

	workspace, err := os.Getwd()
	if err != nil {
		return err
	}

	err = p.createStructure(workspace)
	if err != nil {
		return err
	}

	return nil
}
