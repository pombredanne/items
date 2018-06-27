package cterminal

import (
	"os/exec"
	"strings"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/controller"
	"github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

var PathToSiac string

type terminalController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(cmd string) string `slot:"command"`
}

func (c *terminalController) init() {
	c.ConnectCommand(c.command)
}

func (c *terminalController) command(cmd string) string {
	if cmd == "wallet unlock" {
		if controller.Controller.IsLocked() {
			cdialog.Controller.Show("unlock")
			return ""
		}
		return "Wallet already unlocked"
	} else {
		ecmd := exec.Command(PathToSiac, strings.Split(cmd, " ")...)
		out, err := ecmd.CombinedOutput()
		if err != nil {
			println(err.Error())
		}
		return string(out)
	}
}
