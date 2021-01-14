package vm

import "os"

// I80 (Intel 8080) emulator
type I80 struct {
	// once we know the hardware specs, we will specify the available mem size
	memory []byte
}

// Load the passed romfile and init the vm state
func (vm *I80) Load(romFile string) error {
	rom, err := os.Open(romFile)
	if err != nil {
		return err
	}

	rom.Read(vm.memory)
	return nil
}

// Reset will reset the vm to a clean state
func (vm *I80) Reset() {}

// Start the emulator
func (vm *I80) Start() error {
	return nil
}
