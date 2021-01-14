package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yetibob/i80/emu/vm"
)

var (
	// Used for flags.
	rootCmd = &cobra.Command{
		Use:   "i80",
		Short: "i80 runs the intel 8080 emulator against the given rom file",
		Run: func(cmd *cobra.Command, args []string) {
			// read rom into buffer
			romFile, err := cmd.PersistentFlags().GetString("rom")
			panicErr(err)

			var vm vm.I80
			err = vm.Load(romFile)
			panicErr(err)

			err = vm.Start()
			panicErr(err)
		},
	}
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Execute t
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("rom", "r", "", "rom file to regen")
	rootCmd.MarkPersistentFlagRequired("rom")
}

func initConfig() {
	return
}
