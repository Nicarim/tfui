package tf

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

func GetStateList() string {
	cmd := exec.Command("terraform", "state", "list")
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmd.Dir = path.Join(cwd, "example_terraform")
	output, cmdErr := cmd.CombinedOutput()
	if cmdErr != nil {
		log.Println(cmdErr)
	}

	return string(output)
}
