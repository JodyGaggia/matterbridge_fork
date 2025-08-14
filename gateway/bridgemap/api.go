// +build !noapi

package bridgemap

import (
	"github.com/JodyGaggia/matterbridge_fork/bridge/api"
)

func init() {
	FullMap["api"] = api.New
}
