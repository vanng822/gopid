package gopid

import (
	"fmt"
	"os"
	"testing"
)
var pidFile = "gopid.pid"


func testCleanPid() {
	os.Remove(pidFile)
}

func testCreatePid() {
	pidf, err := os.Create(pidFile)
	if err != nil {
		panic("Could not create pid file")
	}
	defer pidf.Close()
	// hope not killing any process on build server
	pidf.WriteString("24252254242545")
}

func TestMain(m *testing.M) {
	fmt.Println("Test starting")
	testCleanPid()
	retCode := m.Run()
	testCleanPid()
	fmt.Println("Test ending")
	os.Exit(retCode)
}
