/*
 * Copyright (c) 2022-2023 Arm Limited. All rights reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
	"os"

	"github.com/Open-CMSIS-Pack/cbuild/v2/cmd/cbuild/commands"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(new(LogFormatter))
	log.SetOutput(os.Stdout)

	commands.Version = version
	commands.CopyrightNotice = copyrightNotice

	cmd := commands.NewRootCmd()
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

type LogFormatter struct{}

func (s *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	msg := fmt.Sprintf("%s cbuild: %s\n", entry.Level.String(), entry.Message)
	return []byte(msg), nil
}
