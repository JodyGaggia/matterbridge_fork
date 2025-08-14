// +build !nosteam

package bridgemap

import (
	bsteam "github.com/JodyGaggia/matterbridge_fork/bridge/steam"
)

func init() {
	FullMap["steam"] = bsteam.New
}
