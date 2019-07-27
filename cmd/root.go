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
	description  string
	cobraCommand *cobra.Command
}

var topCommands = []ghCommand{
	ghCommand{"config", "Set local and global options", ConfigCmd},
	ghCommand{"init", "Create or re-initialize a git repository", nil},
	ghCommand{"clone", "", nil},
	ghCommand{"status", "", nil},
	ghCommand{"add", "", nil},
	ghCommand{"commit", "", nil},
	ghCommand{"pull", "", nil},
	ghCommand{"push", "", nil},
	ghCommand{"branch", "", nil},
	ghCommand{"checkout", "", nil},
	ghCommand{"stash", "", nil},
	ghCommand{"merge", "", nil},
	ghCommand{"rebase", "", nil},
	ghCommand{"reset", "", nil},
	ghCommand{"remote", "", nil},
	ghCommand{"fetch", "", nil},
	ghCommand{"unstage", "", nil},
	ghCommand{"task", "Not a git command. Show available tasks.", nil},
}

var ghMap = createGhMap(topCommands)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh",
	Short: "Git helper",
	Long:  `Git helper assists with learning git by providing documentation and shortcuts.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Println("MAIN MENU")

		showGhCommands(ghMap, 8)

		// fmt.Println("a. Amend last commit")
		// fmt.Println("b. Pull with rebase")
		// fmt.Println("c. Clear local git cache (git rm -r --cached .)")
		// fmt.Println("d. Reset local branch (git reset --hard origin/some-topic-branch)")
		// fmt.Println("e. Recreate branch from master (ask local and remote)")
		// fmt.Println("f. View origin (git remote -v)")
		input := charChoice()

		ghc, ok := ghMap[input]

		if ok && ghc.cobraCommand != nil {
			cmd.RemoveCommand(ghc.cobraCommand) // remove command from root to allow it to run on its own
			ghc.cobraCommand.Execute()
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

func charChoice() rune {
	fmt.Print("Option: ")
	return readChar()
}

func stringChoice() string {
	fmt.Print("Options: ")
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
	shellCommand.Run()
}
