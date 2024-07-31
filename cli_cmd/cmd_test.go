package cli_cmd

import (
	"testing"
)

func Test_Command_Help_Returns_Nil(t *testing.T) {
	is_nil := CommandHelp()
	if is_nil != nil {
		t.Errorf("Function did not return nil like it was supposed to")
	}
}

func Test_GetCommands_Finds_Commands(t *testing.T) {
	keys := GetCommands()

	Good_Inputs := []string{"help", "quit"}

	helpCMD := CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    CommandHelp,
	}

	quitCMD := CliCommand{
		Name:        "quit",
		Description: "Quit the Pokedex",
		Callback:    CommandQuit,
	}

	for _, input := range Good_Inputs {
		cmd, ok := keys[input]
		if !ok {
			t.Errorf("Cannot find command in the map")
		}
		if input == "help" {
			the_same := helpCMD.Name == cmd.Name && helpCMD.Description == cmd.Description
			if !the_same {
				t.Errorf("The %s struct is not correctly specified", input)
				t.Errorf("Name was: %s \n Description was %s\n", cmd.Name, cmd.Description)
			}
		}

		if input == "quit" {
			the_same := quitCMD.Name == cmd.Name && quitCMD.Description == cmd.Description
			if !the_same {
				t.Errorf("The %s struct is not correctly specified", input)
				t.Errorf("Name was: %s \n Description was %s\n", cmd.Name, cmd.Description)
			}
		}

	}

}
