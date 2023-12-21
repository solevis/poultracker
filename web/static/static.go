package static

import "embed"

//go:embed *
var files embed.FS

// Files returns a filesystem with static files.
func Files() embed.FS {
	return files
}
