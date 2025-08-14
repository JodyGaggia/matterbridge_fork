// +build !nomumble

package bridgemap

import (
	bmumble "github.com/JodyGaggia/matterbridge_fork/bridge/mumble"
)

func init() {
	FullMap["mumble"] = bmumble.New
}
