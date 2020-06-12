package cmd

import (
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/yetibob/i80/emu/op"
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

			rom, err := ioutil.ReadFile(romFile)
			panicErr(err)

			pc := 0

			for {
				if pc == len(rom) {
					break
				}
				adv := op.HandleOp(rom[pc : pc+3])
				pc += adv
			}
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
