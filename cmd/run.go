// Package cmd provides commands for the application.
// Copyright 2020 Don B. Stringham All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.
//
// @author donbstringham <donbstringham@icloud.com>
//
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cs3210-fall2020/gsh/ver"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts the interactive shell",
	Long:  `Starts the interactive shell and waits for user input`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			cmdStr, err := readCmd()
			if err != nil {
				fmt.Println(err)
			}
			err = execCmd(cmdStr)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}

func readCmd() (string, error) {
	fmt.Print("gsh> ")
	reader := bufio.NewReader(os.Stdin)
	cmdStr, err := reader.ReadString('\n')
	if err != nil {
		return cmdStr, err
	}
	return cmdStr, nil
}

func execCmd(cmdStr string) error {
	cmdStr = strings.TrimSuffix(cmdStr, "\n")
	cmdArr := strings.Fields(cmdStr)
	switch cmdArr[0] {
	case "exit":
		os.Exit(0)
	case "ver":
		fmt.Printf("%s\n", ver.Version)
		return nil
	}

	comd := exec.Command(cmdArr[0], cmdArr[1:]...)
	comd.Stderr = os.Stderr
	comd.Stdout = os.Stdout

	return comd.Run()
}
