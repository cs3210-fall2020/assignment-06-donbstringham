// Package cmd provides commands for the application.
// Copyright 2020 Don B. Stringham All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.
//
// @author donbstringham <donbstringham@icloud.com>
//
package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/chzyer/readline"
	"github.com/cs3210-fall2020/gsh/pkg"
	"github.com/cs3210-fall2020/gsh/ver"
	"github.com/spf13/cobra"
)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("bye"),
	readline.PcItem("exit"),
	readline.PcItem("help"),
	readline.PcItem("history"),
	readline.PcItem("mode",
		readline.PcItem("vi"),
		readline.PcItem("emacs"),
	),
	readline.PcItem("setprompt"),
)
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts the interactive shell",
	Long:  `Starts the interactive shell and waits for user input`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			rl, err := readline.NewEx(&readline.Config{
				Prompt:              "\033[31mgsh»\033[0m ",
				HistoryFile:         "history.tmp",
				AutoComplete:        completer,
				InterruptPrompt:     "^C",
				EOFPrompt:           "exit",
				HistorySearchFold:   true,
				FuncFilterInputRune: filterInput,
				VimMode:             false,
			})
			if err != nil {
				panic(err)
			}
			defer rl.Close()

			cmdStr, err := readCmd(rl)
			if err != nil {
				fmt.Println(err)
			}
			err = execCmd(rl, cmdStr)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

}

func filterInput(r rune) (rune, bool) {
	switch r {
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func execCmd(rl *readline.Instance, cmdStr *string) error {
	cmdNew := strings.TrimSuffix(*cmdStr, "\n")
	cmdNew = strings.TrimSpace(cmdNew)
	cmdArr := strings.Fields(cmdNew)
	switch {
	case cmdNew == "bye":
		os.Exit(0)
	case cmdNew == "exit":
		os.Exit(0)
	case cmdNew == "help":
		usage(rl.Stderr())
		return nil
	case strings.HasPrefix(cmdNew, "mode "):
		switch cmdNew[5:] {
		case "vim":
			rl.SetVimMode(true)
			fmt.Println("new mode: vim")
		case "emacs":
			rl.SetVimMode(false)
			fmt.Println("new mode: emacs")
		default:
			fmt.Println("invalid mode:", cmdNew[5:])
		}
		return nil
	case cmdNew == "history":
		hf, err := ioutil.ReadFile("history.tmp")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(hf))
		}
		return nil
	case cmdNew == "mode":
		if rl.IsVimMode() {
			fmt.Println("current mode: vim")
		} else {
			fmt.Println("current mode: emacs")
		}
		return nil
	case strings.HasPrefix(cmdNew, "setprompt"):
		if len(cmdNew) <= 10 {
			log.Println("setprompt <prompt>")
			break
		}
		rl.SetPrompt(cmdNew[10:])
		return nil
	case strings.HasPrefix(cmdNew, "test"):
		if len(cmdArr) <= 1 {
			cmdArr = append(cmdArr, "./")
		}
		files, err := pkg.GetFiles(cmdArr[1])
		if err != nil {
			log.Panicln(err)
			return nil
		}
		for _, file := range files {
			fmt.Println("|" + file)
		}
		return nil
	case cmdNew == "ver":
		fmt.Printf("%s\n", ver.Version)
		return nil
	}
	comd := exec.Command(cmdArr[0], cmdArr[1:]...)
	comd.Stderr = os.Stderr
	comd.Stdout = os.Stdout

	return comd.Run()
}

func readCmd(rl *readline.Instance) (*string, error) {
	for {
		buf, err := rl.Readline()
		if err == readline.ErrInterrupt {
			if len(buf) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		return &buf, nil
	}
	return nil, nil
}

func usage(w io.Writer) {
	io.WriteString(w, "commands:\n")
	io.WriteString(w, completer.Tree(" .   "))
}
