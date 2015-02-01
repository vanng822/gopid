package gopid

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"syscall"
)

func TestGetPidNoPidFile(t *testing.T) {
	assert.Equal(t, 0, GetPid("somepid.pid"))
}

func TestCreatePidOK(t *testing.T) {	
	CreatePid(pidFile)
	fi, err := os.Stat(pidFile)
	assert.Nil(t, err)
	assert.NotNil(t, fi)
	assert.Equal(t, pidFile, fi.Name())
	testCleanPid()
}

func TestCheckForce(t *testing.T) {
	testCreatePid()
	CheckPid(pidFile, true)
	testCleanPid()
}

func TestCleanPid(t *testing.T) {
	CreatePid(pidFile)
	pid := GetPid(pidFile)
	assert.Equal(t, syscall.Getpid(), pid)
	CleanPid(pidFile)
	fi, err := os.Stat(pidFile)
	assert.NotNil(t, err)
	assert.Nil(t, fi)
}