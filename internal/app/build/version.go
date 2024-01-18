package build

import "runtime"

const defaultValue = "unknown"

//nolint:gochecknoglobals // these variables is set during the build
var (
	Version       = defaultValue
	VersionCommit = defaultValue
	BuildDate     = defaultValue
	GoVersion     = runtime.Version()
)
