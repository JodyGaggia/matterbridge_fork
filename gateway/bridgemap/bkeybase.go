// +build !nokeybase

package bridgemap

import (
	bkeybase "github.com/JodyGaggia/matterbridge_fork/bridge/keybase"
)

func init() {
	FullMap["keybase"] = bkeybase.New
}
