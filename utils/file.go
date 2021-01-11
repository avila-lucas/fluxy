package utils

import (
	"fluxy/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

const configFile = ".fluxy"

func ReadYaml(path string, data interface{}) error {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return FileNotFound
	}

	err = yaml.Unmarshal(fileContent, data)

	return err
}

func LoadConfigFile() (*models.Config, error) {
	config := &models.Config{}
	home, err := os.UserHomeDir()
	if err != nil {
		return config, err
	}

	err = ReadYaml(home+"/"+configFile, config)

	return config, err
}

func SaveConfigFile(config *models.Config) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	newFileContent, err := yaml.Marshal(config)
	if err != nil {
		return ParsingError
	}

	return ioutil.WriteFile(home+"/"+configFile, newFileContent, 0644)
}

func ModifyLine(path string, targetLine string, toReplace string, newValue string) error {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return FileNotFound
	}

	lines := strings.Split(string(fileContent), "\n")

	for idx, line := range lines {
		if strings.Contains(line, targetLine) {
			lines[idx] = strings.Replace(line, toReplace, newValue, 1)
		}
	}

	newFileContent := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(newFileContent), 0644)

	return err
}
