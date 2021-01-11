package deploy

import "github.com/urfave/cli/v2"

func NewDeployCommand(project, container, tag *string) *cli.Command {
	return &cli.Command{
		Name:    "deploy",
		Usage:   "changes tag version in configuration files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "project",
				Aliases:     []string{"p"},
				Usage:       "project name to retrieve config with file paths",
				Required:    true,
				Destination: project,
			},
			&cli.StringFlag{
				Name:        "container",
				Aliases:     []string{"c"},
				Usage:       "container name to deploy",
				Required:    true,
				Destination: container,
			},
			&cli.StringFlag{
				Name:        "tag",
				Aliases:     []string{"t"},
				Usage:       "tag name of the release to deploy",
				Required:    true,
				Destination: tag,
			},
		},
		Action: exec(project, container, tag),
	}
}
