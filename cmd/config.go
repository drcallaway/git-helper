/*
Copyright © 2019 Dustin R. Callaway

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var configCommands = []ghCommand{
	ghCommand{name: "show value", description: "Show config setting for given name."},
	ghCommand{name: "add/change", description: "Add config setting or change existing. \"user.name\" and \"user.email\" are common settings."},
	ghCommand{name: "--list", description: "List all config values. Defaults to --local settings."},
	ghCommand{name: "--show-origin", description: "Source of settings. Requires --list."},
	ghCommand{name: "--add", description: "Add config setting."},
	ghCommand{name: "--replace-all", description: "Replace all rows with given name using new value."},
	ghCommand{name: "--unset", description: "Remove config setting."},
	ghCommand{name: "--global", description: "Read or write only global config settings. Requires --list to read, name/value to set, or name/--unset to remove."},
	ghCommand{name: "--local", description: "Read or write only local config settings. Requires --list to read, name/value to set, or name/--unset to remove."},
}

var configMap = createGhMap(configCommands)

// ConfigCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Set local and global options",
	Long: `You can query/set/replace/unset options with this command. The name is actually the section
and the key separated by a dot, and the value will be escaped.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Println("CONFIG MENU")
		showGhCommands(configMap, 13)

		input := stringChoiceOptions()

		shellArgs := []string{}

		if input == "" {
			return
		}

		for i := 0; i < len(input); i++ {
			ghCommand := configMap[rune(input[i])]

			switch ghCommand.name {
			case "show value":
				name := stringChoice("Enter name: ")
				shellArgs = append(shellArgs, name)
			case "add/change", "--add", "--replace-all":
				name := stringChoice("Enter name: ")
				if ghCommand.name == "add/change" {
					shellArgs = append(shellArgs, name)
				} else {
					shellArgs = append(shellArgs, ghCommand.name, name)
				}
				value := stringChoice("Enter value: ")
				shellArgs = append(shellArgs, value)
			case "--unset":
				unset := stringChoice("Enter name to remove: ")
				shellArgs = append(shellArgs, "--unset", unset)
			case "--global", "--local":
				shellArgs = append([]string{ghCommand.name}, shellArgs...)
			default:
				shellArgs = append(shellArgs, ghCommand.name)
			}
		}

		shellArgs = append([]string{"config"}, shellArgs...)

		fullCommand := strings.Join(shellArgs, " ")

		fmt.Println()

		color.Red("git %v", fullCommand)

		fmt.Print("\nExecute? (Y/n) ")

		executeFlag := readString()

		fmt.Println()

		if executeFlag != "n" && executeFlag != "N" {
			executeShellCommand(shellArgs)
			fmt.Println()
		}

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(ConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
