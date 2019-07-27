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
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

type ghCommand struct {
	name         string
	short        string
	description  string
	cobraCommand *cobra.Command
}

var topCommands = []ghCommand{
	ghCommand{name: "config", description: "Set local and global options", cobraCommand: ConfigCmd},
	ghCommand{name: "init", description: "Create or re-initialize a git repository", cobraCommand: nil},
	ghCommand{name: "clone", description: "", cobraCommand: nil},
	ghCommand{name: "status", description: "", cobraCommand: nil},
	ghCommand{name: "add", description: "", cobraCommand: nil},
	ghCommand{name: "commit", description: "", cobraCommand: nil},
	ghCommand{name: "pull", description: "", cobraCommand: nil},
	ghCommand{name: "push", description: "", cobraCommand: nil},
	ghCommand{name: "branch", description: "", cobraCommand: nil},
	ghCommand{name: "checkout", description: "", cobraCommand: nil},
	ghCommand{name: "stash", description: "", cobraCommand: nil},
	ghCommand{name: "merge", description: "", cobraCommand: nil},
	ghCommand{name: "rebase", description: "", cobraCommand: nil},
	ghCommand{name: "reset", description: "", cobraCommand: nil},
	ghCommand{name: "remote", description: "", cobraCommand: nil},
	ghCommand{name: "fetch", description: "", cobraCommand: nil},
	ghCommand{name: "unstage", description: "", cobraCommand: nil},
	ghCommand{name: "task", description: "Not a git command. Show available tasks.", cobraCommand: nil},
}

var ghMap = createGhMap(topCommands)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh",
	Short: "Git helper",
	Long:  `Git helper assists with learning git by providing documentation and shortcuts.`,
	Run: func(cmd *cobra.Command, args []string) {

		for {
			fmt.Println()
			fmt.Println("MAIN MENU")

			showGhCommands(ghMap, 8)

			// fmt.Println("a. Amend last commit")
			// fmt.Println("b. Pull with rebase")
			// fmt.Println("c. Clear local git cache (git rm -r --cached .)")
			// fmt.Println("d. Reset local branch (git reset --hard origin/some-topic-branch)")
			// fmt.Println("e. Recreate branch from master (ask local and remote)")
			// fmt.Println("f. View origin (git remote -v)")
			input := charChoiceOption()

			if input == 0 {
				return
			}

			ghc, ok := ghMap[input]

			if ok && ghc.cobraCommand != nil {
				cmd.RemoveCommand(ghc.cobraCommand) // remove command from root to allow it to run on its own
				ghc.cobraCommand.Execute()
			}
		}

		// switch input {
		// case 'a':
		// 	configCommand := ghMap[input].cobraCommand
		// 	cmd.RemoveCommand(configCommand) // remove command from root to allow it to run on its own
		// 	configCommand.Execute()
		// }
	},
}

func createGhMap(list []ghCommand) map[rune]ghCommand {
	m := make(map[rune]ghCommand, len(list))
	letter := 'a'

	for _, cmd := range list {
		m[letter] = cmd
		letter++
	}

	return m
}

func showGhCommands(m map[rune]ghCommand, columnWidth int) {
	for i := 0; i < len(m); i++ {
		letter := rune('a' + i)
		ghc := m[letter]
		fmt.Printf("%v. %-*v - %v\n", string(letter), columnWidth, ghc.name, ghc.description)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func charChoiceOption() rune {
	fmt.Print("Option: ")
	return readChar()
}

func stringChoiceOptions() string {
	fmt.Print("Options: ")
	return readString()
}

func stringChoice(prompt string) string {
	fmt.Print(prompt)
	return readString()
}

func readString() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readChar() rune {
	input := readString()

	if len(input) < 1 {
		return 0
	}

	return rune(input[0])
}

func executeShellCommand(shellArgs []string) {
	shellCommand := exec.Command("git", shellArgs...)
	shellCommand.Stdin = os.Stdin
	shellCommand.Stdout = os.Stdout
	shellCommand.Stderr = os.Stderr
	shellCommand.Run()
}

func findGhCommand(commands []ghCommand, name string) (ghCommand, bool) {
	for _, cmd := range commands {
		if cmd.name == name {
			return cmd, true
		}
	}

	return ghCommand{}, false
}
