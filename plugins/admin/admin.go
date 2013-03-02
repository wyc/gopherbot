// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package admin

import (
	"github.com/jteeuwen/ircb/cmd"
	"github.com/jteeuwen/ircb/plugin"
	"github.com/jteeuwen/ircb/proto"
)

func init() { plugin.Register(New) }

type Plugin struct {
	*plugin.Base
}

func New(profile string) plugin.Plugin {
	p := new(Plugin)
	p.Base = plugin.New(profile, "admin")
	return p
}

func (p *Plugin) Load(c *proto.Client) (err error) {
	err = p.Base.Load(c)
	if err != nil {
		return
	}

	quit := new(cmd.Command)
	quit.Name = "quit"
	quit.Restricted = true
	quit.Execute = func(cmd *cmd.Command, c *proto.Client, m *proto.Message) {
		c.Quit("")
	}
	cmd.Register(quit)

	return
}