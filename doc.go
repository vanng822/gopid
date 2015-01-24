// Package gopid is for checking, creating and cleaning up pid file
// 
//	import (
//		"github.com/vanng822/gopid"
//	)
//	func main() {
//		// code for getting pidFile
//		gopid.CheckPid(pidFile, force)
//		gopid.CreatePid(pidFile)
//		defer gopid.CleanPid(pidFile)
//		// running your code
//	}
package gopid
