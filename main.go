package main

import (
	"fluxy/commands/automated"
	"fluxy/commands/create"
	"fluxy/commands/deploy"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var tag string
	var project string
	var container string
	var fluxPatch string
	var annotations string

	app := cli.NewApp()
	app.Usage = "A simple CLI tool to manage deploys versions working with FluxCD and GIT-repositories"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		deploy.NewDeployCommand(&project, &container, &tag),
		create.NewProjectCommand(&project, &annotations, &fluxPatch),
		automated.NewAutomatedCommand(&project),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
