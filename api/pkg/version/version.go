package version

import (
	"fmt"
)

// Version is the current version of the APIs
const Version = "v0.0.1"

// LongVersion returns the current version of the APIs
func LongVersion() string {
	return fmt.Sprintf("MarktHub APIs %s", Version)
}
