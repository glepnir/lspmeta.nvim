// Copyright 2021 glepnir. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package pkg

import (
	"github.com/glepnir/lspmeta.nvim/pkg/logger"
	"github.com/neovim/go-client/nvim"
)

type Command struct {
	nvim   *nvim.Nvim
	logger *logger.Logger
}

func NewCommand(v *nvim.Nvim, l *logger.Logger) *Command {
	return &Command{
		nvim:   v,
		logger: l,
	}
}

type userLspServers []string

func (c *Command) LspInstall(args []string, bang bool) {
	userConfig := userLspServers{}
	getUserConfig := "require('lspmeta').get_user_config_server()"
	if err := c.nvim.ExecLua(getUserConfig, userConfig); err != nil {
		c.logger.Print(err)
	}
}
