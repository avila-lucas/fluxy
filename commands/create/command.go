package create

import (
	"github.com/urfave/cli/v2"
)

func NewProjectCommand(project, annotations, fluxPatch *string) *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "creates a new project settings",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "project",
				Aliases:     []string{"p"},
				Usage:       "project name",
				Required:    true,
				Destination: project,
			},
			&cli.StringFlag{
				Name:        "annotations",
				Aliases:     []string{"a"},
				Usage:       "absolute file path where flux annotations are present",
				Required:    true,
				Destination: annotations,
			},
			&cli.StringFlag{
				Name:        "fpatch",
				Aliases:     []string{"f"},
				Usage:       "absolute file path where flux patch is present",
				Required:    true,
				Destination: fluxPatch,
			},
		},
		Action: exec(project, annotations, fluxPatch),
	}
}
