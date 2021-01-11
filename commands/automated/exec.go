package automated

import (
	"fluxy/utils"
	"fmt"
	"github.com/kyokomi/emoji/v2"
	"github.com/urfave/cli/v2"
)

func exec(project *string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		config, err := utils.LoadConfigFile()
		if err != nil {
			fmt.Printf(emoji.Sprint(":prohibited: ERROR reading config file (.fluxy at your home directory)", err))
			return nil
		}

		projectConfig, exists := config.Projects[*project]
		if !exists {
			fmt.Println(emoji.Sprint(":prohibited: No project configuration setup found in .fluxy"))
			return nil
		}

		err = enableAutomaticDeploys(projectConfig.Annotations)
		if err != nil {
			return nil
		}

		fmt.Println(emoji.Sprint(":coffee: Enabled automatic deploys!"))
		return nil
	}
}

func enableAutomaticDeploys(annotationsPath string) error {
	err := utils.ModifyLine(annotationsPath, "fluxcd.io/automated", "false", "true")
	if err != nil {
		if utils.IsFileNotFound(err) {
			fmt.Println(emoji.Sprintf(":prohibited: Annotations file not found. Please check [%s]", annotationsPath))
			return err
		}

		fmt.Println(emoji.Sprint(":prohibited: Error while activating automatic deploys"))

		return err
	}

	return nil
}
