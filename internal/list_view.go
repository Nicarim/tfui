package internal

import (
	"github.com/rivo/tview"
	"strings"
	"tfui/internal/tf"
)

func RenderList() {
	tfState := tf.GetStateList()
	tfStateList := strings.Split(tfState, "\n")
	app := tview.NewApplication()
	list := tview.NewList()
	for _, el := range tfStateList {
		if el != "" {
			list.AddItem(el, "TF State", 0, nil)
		}
	}
	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})
	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
