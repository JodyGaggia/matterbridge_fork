// +build !nonctalk

package bridgemap

import (
	btalk "github.com/JodyGaggia/matterbridge_fork/bridge/nctalk"
)

func init() {
	FullMap["nctalk"] = btalk.New
}
