// +build !norocketchat

package bridgemap

import (
	brocketchat "github.com/JodyGaggia/matterbridge_fork/bridge/rocketchat"
)

func init() {
	FullMap["rocketchat"] = brocketchat.New
}
