package cli_cmd

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"quit": {
			Name:        "quit",
			Description: "Quit the Pokedex",
			Callback:    CommandQuit,
		},
	}
}
