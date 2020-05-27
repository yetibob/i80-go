package cmd

import (
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/yetibob/i80/regen/op"
)

var (
	// Used for flags.
	rootCmd = &cobra.Command{
		Use:   "regen",
		Short: "regen runs through an 8080 program and spits out assembly ",
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
	rootCmd.PersistentFlags().StringP("rom", "r", "./roms/invaders.h", "rom file to regen")
}

func initConfig() {
	return
}
