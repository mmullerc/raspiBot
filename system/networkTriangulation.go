package system

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func CheckNetworkStrngth(c chan string) {
	// docker build current directory
	//cmdType := "bash"
	// /(ESSID)/g

	var result string

	cmdName := "$sudo iwlist wlan0 scan"

	cmd := exec.Command("/bin/sh", "-c", cmdName)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			result += scanner.Text()
		}
		c <- result
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
}
