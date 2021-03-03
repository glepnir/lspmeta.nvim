// Copyright 2021 glepnir. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package lspmeta

import (
	"log"
	"os"

	"github.com/neovim/go-client/nvim"
	"github.com/neovim/go-client/nvim/plugin"
)

func Register(p *plugin.Plugin) error {
	stdin := os.Stdin
	stdout := os.Stdout
	v, err := nvim.New(stdin, stdout, stdout, log.Printf)
	if err != nil {
		return err
	}
	c := NewCommand(v)
	p.HandleFunction(
		&plugin.FunctionOptions{Name: "GetUserLspConfig"},
		func(args []string) ([]string, error) {
			return c.LspInstall(p)
		},
	)
	p.HandleFunction(
		&plugin.FunctionOptions{Name: "TestFunc"},
		func(args []string) (string, error) {
			return "hello", nil
		},
	)
	return nil
}
