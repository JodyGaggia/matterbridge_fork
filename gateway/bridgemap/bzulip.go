// +build !nozulip

package bridgemap

import (
	bzulip "github.com/JodyGaggia/matterbridge_fork/bridge/zulip"
)

func init() {
	FullMap["zulip"] = bzulip.New
}
