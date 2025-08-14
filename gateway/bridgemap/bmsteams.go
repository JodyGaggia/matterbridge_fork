// +build !nomsteams

package bridgemap

import (
	bmsteams "github.com/JodyGaggia/matterbridge_fork/bridge/msteams"
)

func init() {
	FullMap["msteams"] = bmsteams.New
}
