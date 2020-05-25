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
			if err != nil {
				panic(err)
			}
			rom, err := ioutil.ReadFile(romFile)
			if err != nil {
				panic(err)
			}

			pc := 0

			for {
				if pc == len(rom) {
					break
				}
				code := rom[pc]
				adv := op.HandleOp(code, rom[pc+1:pc+3])
				pc += adv
			}
		},
	}
)

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
