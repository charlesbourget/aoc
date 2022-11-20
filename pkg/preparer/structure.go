package preparer

import (
	"bufio"
	"embed"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/charlesbourget/aoc/pkg/utils"
)

//go:embed all:templates
var templateFS embed.FS

type config struct {
	Day  int
	Year int
}

type FileStructureConfig struct {
	Structures []Structure `yaml:"structures"`
}

type Structure struct {
	Template  string `yaml:"template"`
	File      string `yaml:"file"`
	Directory string `yaml:"directory"`
}

const templateDir = "templates"
const fileStructureConfigFileName = "structure.toml"

func (p Preparer) createStructure(workspace string) error {
	day := p.Day
	year := p.Year

	name := fmt.Sprintf("day%02d", day)
	path := filepath.Join(workspace, name)

	err := utils.CreateDir(path)
	if err != nil {
		return err
	}

	var fileStructureConfig FileStructureConfig
	readFileStructureConfig(&fileStructureConfig, p.Lang)

	for _, structure := range fileStructureConfig.Structures {
		fileName := structure.File
		if strings.Contains(fileName, "%02d") {
			fileName = fmt.Sprintf(structure.File, day)
		}

		if structure.Directory != "" {
			if err := utils.CreateDir(filepath.Join(path, structure.Directory)); err != nil {
				return err
			}

			fileName = filepath.Join(structure.Directory, fileName)
		}

		err = createFileTemplate(filepath.Join(templateDir, p.Lang, structure.Template), filepath.Join(path, fileName), day, year)
		if err != nil {
			return err
		}
	}

	err = createInputFile(filepath.Join(path, "input.txt"), day, year, p.Session)
	if err != nil {
		return err
	}

	return nil
}

func createFileTemplate(templateFilePath string, filePath string, day int, year int) error {
	file, err := utils.CreateFile(filePath)
	if err != nil {
		return nil
	}

	t, err := template.ParseFS(templateFS, templateFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	err = t.Execute(file, &config{day, year})
	if err != nil {
		return err
	}

	return nil
}

func createInputFile(filePath string, day int, year int, session string) error {
	file, err := utils.CreateFile(filePath)
	if err != nil {
		return nil
	}

	if file == nil {
		fmt.Println("Input already exists. Skipping")
		return nil
	}

	input, err := fetchInput(day, year, session)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	_, err = writer.Write(input)
	if err != nil {
		return err
	}

	return nil
}

func readFileStructureConfig(fileStructureConfig *FileStructureConfig, lang string) error {
	configFile, err := templateFS.ReadFile(filepath.Join(templateDir, lang, fileStructureConfigFileName))
	if err != nil {
		return err
	}

	if err := utils.ReadTomlFile(configFile, &fileStructureConfig); err != nil {
		return err
	}

	return nil
}
