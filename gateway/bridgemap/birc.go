// +build !noirc

package bridgemap

import (
	birc "github.com/JodyGaggia/matterbridge_fork/bridge/irc"
)

func init() {
	FullMap["irc"] = birc.New
}
