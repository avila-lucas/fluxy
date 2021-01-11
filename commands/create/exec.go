package create

import (
	"fluxy/models"
	"fluxy/utils"
	"fmt"
	"github.com/kyokomi/emoji/v2"
	"github.com/urfave/cli/v2"
)

func exec(project, annotationsPath, fluxPatchPath *string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		config, err := utils.LoadConfigFile()
		if err != nil {
			if utils.IsFileNotFound(err) {
				config = &models.Config{Projects: make(map[string]models.Project, 1)}
			} else {
				fmt.Printf(emoji.Sprint(":prohibited: Problem parsing config file", err))
				return err
			}
		}

		config.Projects[*project] = models.Project{
			Annotations: *annotationsPath,
			FluxPatch:   *fluxPatchPath,
		}

		err = utils.SaveConfigFile(config)
		if err != nil {
			if utils.IsParsingError(err) {
				fmt.Println(emoji.Sprint(":prohibited: Problem while saving new data at parsing stage"))
				return err
			}

			fmt.Println(emoji.Sprint(":prohibited: Error while saving new content"))
			return err
		}

		fmt.Println(emoji.Sprint(":star: New project added!"))
		return nil
	}
}
