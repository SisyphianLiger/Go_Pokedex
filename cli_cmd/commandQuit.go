package cli_cmd

import "os"

func CommandQuit() error {
	os.Exit(0)
	return nil
}
