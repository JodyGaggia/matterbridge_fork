// +build !nogitter

package bridgemap

import (
	bgitter "github.com/JodyGaggia/matterbridge_fork/bridge/gitter"
)

func init() {
	FullMap["gitter"] = bgitter.New
}
