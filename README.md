# gopid

Simple handling of pid file, like checking, creating and cleaning up pid file


## GoDoc

[![GoDoc](https://godoc.org/github.com/vanng822/gopid?status.svg)](https://godoc.org/github.com/vanng822/gopid)


## Example

	import (
		"github.com/vanng822/gopid"
	)
	
	func main() {
		// code for getting pidFile and force
		gopid.CheckPid(pidFile, force)
		gopid.CreatePid(pidFile)
		defer gopid.CleanPid(pidFile)
		// running your code
	}