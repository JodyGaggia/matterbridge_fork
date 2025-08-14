// +build !noxmpp

package bridgemap

import (
	bxmpp "github.com/JodyGaggia/matterbridge_fork/bridge/xmpp"
)

func init() {
	FullMap["xmpp"] = bxmpp.New
}
