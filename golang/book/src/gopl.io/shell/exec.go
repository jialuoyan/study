package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	// "github.com/fatih/color"
	"runtime"
)

var bashname string

func Execute(script string) string {
	//bash_path, err := exec.LookPath("sh")
	if  runtime.GOOS == "windows"{
		bashname = "bash"
	} else {
		bashname = "sh"
	}
	output, _ := exec.Command(bashname, "-c", "echo "+strconv.Quote(script)).CombinedOutput()

	// fmt.Printf("> %s", color.GreenString(string(output)))
	fmt.Printf("> %s", string(output))

	cmd := exec.Command(bashname, "-c", script)

	stdoutReader, err := cmd.StdoutPipe()
	defer stdoutReader.Close()
	if err != nil {
		fmt.Printf("Error creating StdoutPipe for Cmd %s", err)
		os.Exit(1)
	}

	var stdout string

	stdoutScanner := bufio.NewScanner(stdoutReader)
	go func() {
		for stdoutScanner.Scan() {
			text := stdoutScanner.Bytes()
			stdout += string(text) + "\n"
			fmt.Fprintf(os.Stdout, "> %s \n", text)
		}
	}()

	stderrReader, err := cmd.StderrPipe()
	defer stderrReader.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	stderrScanner := bufio.NewScanner(stderrReader)
	go func() {
		for stderrScanner.Scan() {
			// fmt.Printf("> %s \n", color.RedString(stderrScanner.Text()))
			fmt.Printf("> %s \n", stderrScanner.Text())
		}
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

	return string(stdout)
}
