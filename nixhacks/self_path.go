package nixhacks

import (
	"fmt"
	"os"
)

// SelfPath represents the about path to self.
var SelfPath string

func init() {
	var err error
	SelfPath, err = os.Executable()
	if err != nil {
		panic(fmt.Sprintf("error getting the path to the executable: %s", err))
	}
}
