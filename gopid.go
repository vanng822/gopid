package gopid

import (
	"log"
	"strconv"
	"os"
	"syscall"
	"fmt"
	"io/ioutil"
)

// Zero means no pid file detected, any other error will cause fatal
func getPid(pidFile string) int {
	if _, err := os.Stat(pidFile); err != nil {
		log.Printf("Could not stat pid file %s", err)
		return 0
	}
	existPidFile, err := os.Open(pidFile)
	if err != nil {
		log.Fatalf("Error when open file pid file %s, error: %v", pidFile, err)
	}
	defer existPidFile.Close()
	existPidB, err := ioutil.ReadAll(existPidFile)
	if err != nil {
		log.Fatalf("Pidfile %s exist but can not read it", pidFile)
	}
	existPid, err := strconv.Atoi(fmt.Sprintf("%s", existPidB))
	if err != nil {
		log.Fatalf("Could not read from pid file %s, error: %v", pidFile, err)
	}
	return existPid
}

// CheckPid will fatal if the pid file exists and not force,
// any error when getting pid from this file will also cause fatal.
// If no pid file live will go on.
// If there is a pid it will try to kill that process and not fatal
// even if it fails
func CheckPid(pidFile string, force bool) {
	log.Printf("Checking pid file %s", pidFile)
	pid := getPid(pidFile)
	if pid == 0 {
		return
	}
	if !force {
		log.Fatalf("Pidfile %s exists with pid %d", pidFile, pid)
	}
	if pid > 0 {
		log.Printf("Try to kill process with pid %d", pid)
		if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
			// if force then we run over it so no fatal here
			// you have to kill it manually if it is still there
			log.Printf("Could not kill pid %d, error: %v", pid, err)
		}
	}
}

// CreatePid creates a pid file with the current pid.
// It will fatal if can not create one
func CreatePid(pidFile string) {
	log.Printf("Creating pid file %s", pidFile)
	pid := syscall.Getpid()
	pidf, err := os.Create(pidFile)
	if err != nil {
		log.Fatalf("Could not create pid file, error: %v", err)
	}
	pidf.WriteString(fmt.Sprintf("%d", pid))
	pidf.Close()
}

// CleanPid will check first if the current process owns the pid in the file.
// If it does then remove the file
func CleanPid(pidFile string) {
	log.Printf("Cleaning up pid file %s", pidFile)
	existPid := getPid(pidFile)
	pid := syscall.Getpid()
	if existPid != pid {
		log.Printf("Not owning pidFile %s, current pid: %d, but pid from file: %d", pidFile, pid, existPid)
		return
	}
	err := os.Remove(pidFile)
	if err != nil {
		log.Printf("Fail to remove pid file %s", pidFile)
	}
}
