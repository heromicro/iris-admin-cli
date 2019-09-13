package cmd

import (
	"fmt"

	"github.com/wanhello/iris-admin-cli/cmd/generate"
	"github.com/wanhello/iris-admin-cli/cmd/new"

	"github.com/urfave/cli"
)

func NewCommmand() cli.Command {
	return cli.Command{
		Name:                   "new",
		Aliases:                []string{"n"},
		Usage:                  "",
		UsageText:              "",
		Description:            "",
		ArgsUsage:              "",
		Category:               "",
		BashComplete:           nil,
		Before:                 nil,
		After:                  nil,
		Action:                 func(ctx *cli.Context) error {
			cfg := new.Config{
				Dir: c.String("dir"),
			}
		},
		OnUsageError:           nil,
		Subcommands:            nil,
		Flags:                  []cli.Flag{

		},
		SkipFlagParsing:        false,
		SkipArgReorder:         false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}

}




