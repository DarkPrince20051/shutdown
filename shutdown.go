package main

import (
	"fmt"
	"os/exec"
)

func shutdown() {
	cmd := exec.Command("shutdown", "/s", "/t", "0")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Shutdown command failed: %v\n", err)
	} else {
		fmt.Println("Shutdown command executed successfully")
	}
}

func main() {
	shutdown()
}
