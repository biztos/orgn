// Package orgn runs an Origin Server for CDN-backed content.
package orgn

import (
	"github.com/biztos/orgn/cmd"
)

// Run exposes the cmd submodule's Execute function.
var Run = cmd.Execute
