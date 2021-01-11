package automated

import "github.com/urfave/cli/v2"

func NewAutomatedCommand(project *string) *cli.Command {
	return &cli.Command{
		Name:  "automated",
		Usage: "enable automated deploys",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "project",
				Aliases:     []string{"p"},
				Usage:       "project name",
				Required:    true,
				Destination: project,
			},
		},
		Action: exec(project),
	}
}
