package main

// HACK: Access internal client.setMaxBodySize.
// Remove this once setMaxBodySize becomes exported, see issue #732.

// go:linkname setMaxBodySize git.corp.com/client.setMaxBodySize
func setMaxBodySize(size int64)

func main() {
	setMaxBodySize(1024)
	//c := client.New()
}
