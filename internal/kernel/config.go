package kernel

import (
	"fmt"
	"path/filepath"

	"github.com/ithinkiborkedit.git/Ouroboros/utils"
)

func (kb *KernelBuilder) ConfigureKernel(configfile string) error {
	kb.ConfigFile = configfile
	if configfile != "" {
		fmt.Printf("Using custom kernel config: %s\n", configfile)
		configPath := filepath.Join(kb.SourceDir, ".config")
		if err := utils.CopyFile(configfile, configPath); err != nil {
			return fmt.Errorf("error copying config file: %v", err)
		}
	}
}
