package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/s16rv/s16/app"
	"github.com/s16rv/s16/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		"s16",
		"cosmos",
		app.DefaultNodeHome,
		"s16-1",
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(cmd.AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
