package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGHUP)

	haProc, err := StartHAProxy("", "")

	if err != nil {
		panic("")
	}

	templProc, err := StartConsulTemplate()

	if err != nil {
		panic("")
	}

	go TrackProcess(haProc)
	go TrackProcess(templProc)
	go RespondToHub(signals)

	for {
	}
}

func RespondToHub(signals chan os.Signal) {
	for range signals {
		fmt.Println("Got a SIGHUP! Now restarting haproxy...")
		RestartHAProxy()
	}
}

func RestartHAProxy() error {
	// /opt/bin/haproxy -p $pid_path -f $config_path -D -sf $(cat $pid_path)
	fmt.Println("Restarting HAProxy")
	return nil
}

func StartHAProxy(pidPath string, configPath string) (*os.Process, error) {
	fmt.Println(pidPath, configPath)
	// args := []string{
	//  "haproxy",
	// 	"-p",
	// 	pidPath,
	// 	"-f",
	// 	configPath,
	// 	"-D",
	// }
	//
	// proc, err := os.StartProcess("/opt/bin/haproxy", args, nil)
	args := []string{"tail", "-f", "/dev/null"}
	return os.StartProcess("/usr/bin/tail", args, &os.ProcAttr{})
}

func StartConsulTemplate() (*os.Process, error) {
	args := []string{"tail", "-f", "/dev/null"}
	return os.StartProcess("/usr/bin/tail", args, &os.ProcAttr{})
}

func TrackProcess(proc *os.Process) {
	proc.Wait()
	panic("No")
}
