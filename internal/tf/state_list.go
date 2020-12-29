package tf

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type TfOp int

const (
	NoOp TfOp = iota
	RenameOp
	DeleteOp
)

type StateOp struct {
	OriginalName string
	NewName      string
	Op           TfOp
}

func GetStateList() []StateOp {
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

	tfStateList := strings.Split(string(output), "\n")

	var tfStateOpList []StateOp

	for _, s := range tfStateList {
		if s != "" {
			tfStateOpList = append(tfStateOpList, StateOp{
				OriginalName: s,
				Op:           NoOp,
			})
		}
	}

	return tfStateOpList
}
