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
	"strings"

	"github.com/spf13/cobra"
)

var configCommands = []ghCommand{
	ghCommand{"name", "Name of config to add/list/remove.", nil},
	ghCommand{"value", "Value of config to add. Must follow name and requires --add.", nil},
	ghCommand{"--list", "List all config values.", nil},
	ghCommand{"--show-origin", "Source of settings. Requires --list.", nil},
	ghCommand{"--add", "Add config. Requires name and value.", nil},
	ghCommand{"--unset", "Remove config. Requires name.", nil},
	ghCommand{"--global", "Global config. Requires --list or --add.", nil},
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
		showGhCommands(configMap, 12)

		input := stringChoice()

		shellArgs := []string{"config"}

		for i := 0; i < len(input); i++ {
			ghCommand := configMap[rune(input[i])]
			shellArgs = append(shellArgs, ghCommand.name)
		}

		fullCommand := strings.Join(shellArgs, " ")

		fmt.Printf("\ngit %v\n\nExecute? (Y/n)", fullCommand)

		executeFlag := readString()

		if executeFlag != "n" && executeFlag != "N" {
			executeShellCommand(shellArgs)
		}
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
