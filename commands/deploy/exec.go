package deploy

import (
	"fluxy/models"
	"fluxy/utils"
	"fmt"
	"github.com/kyokomi/emoji/v2"
	"github.com/urfave/cli/v2"
	"strings"
)

func exec(project, container, tag *string) func(c *cli.Context) error {
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

		err = disableAutomatedDeploy(projectConfig.Annotations)
		if err != nil {
			return nil
		}

		err = updateContainerTag(projectConfig.FluxPatch, *container, *tag)
		if err != nil {
			deployDisplayError(err)
			return err
		}

		fmt.Println(emoji.Sprint(":beer: Files updated"))
		return nil
	}
}

func disableAutomatedDeploy(annotationsPath string) error {
	err := utils.ModifyLine(annotationsPath, "fluxcd.io/automated", "true", "false")
	if err != nil {
		if utils.IsFileNotFound(err) {
			fmt.Println(emoji.Sprintf(":prohibited: Annotations file not found. Please check [%s]", annotationsPath))
			return err
		}

		fmt.Println(emoji.Sprint(":prohibited: Error while disabling automatic deploys"))

		return err
	}

	return nil
}

func updateContainerTag(patchPath string, containerName string, tag string) error {
	fluxPatch := &models.FluxPatch{}
	err := utils.ReadYaml(patchPath, fluxPatch)
	if err != nil {
		return err
	}

	containerFound, previousImage, newImage := modifyImageTag(fluxPatch.Spec.Template.Spec.Containers, containerName, tag)
	if !containerFound {
		return ContainerNotFound
	}

	err = utils.ModifyLine(patchPath, previousImage, previousImage, newImage)
	if err != nil {
		return err
	}

	return nil
}

func modifyImageTag(containers []models.Containers, containerName string, tag string) (bool, string, string) {
	var previousImage string
	var newImage string
	var containerFound bool
	for _, container := range containers {
		if container.Name == containerName {
			containerFound = true
			previousImage = container.Image
			segmentedImage := strings.Split(container.Image, ":")
			segmentedImage[len(segmentedImage)-1] = tag
			newImage = strings.Join(segmentedImage, ":")
		}
	}

	return containerFound, previousImage, newImage
}

func deployDisplayError(err error) {
	if utils.IsFileNotFound(err) {
		fmt.Println(emoji.Sprint(":prohibited: Annotations file not found"))
		return
	}

	if IsContainerNotFound(err) {
		fmt.Println(emoji.Sprint(":prohibited: No container found with that name in the flux patch file"))
		return
	}

	fmt.Println(emoji.Sprint(":prohibited: Not able to change tag"))
	return
}
