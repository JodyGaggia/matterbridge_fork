// +build !nomatrix

package bridgemap

import (
	bmatrix "github.com/JodyGaggia/matterbridge_fork/bridge/matrix"
)

func init() {
	FullMap["matrix"] = bmatrix.New
}
