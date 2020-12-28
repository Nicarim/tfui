package internal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
	"tfui/internal/tf"
)

var app = tview.NewApplication()
var pages = tview.NewPages()
var list = tview.NewList()
var tfStateList []string = nil

func GetListFrame() *tview.Frame {
	tfState := tf.GetStateList()
	tfStateList = strings.Split(tfState, "\n")
	list.ShowSecondaryText(false).
		SetSelectedFunc(listSelectHandler)
	for _, el := range tfStateList {
		if el != "" {
			list.AddItem(el, "", 0, nil)
		}
	}
	list.AddItem("Quit", "Press to exit", 'q', nil)

	frame := tview.NewFrame(list).
		AddText("Rename State", true, tview.AlignCenter, tcell.ColorWhite)
	return frame
}

func GetRenameView(s string) *tview.Flex {
	modalFunc := func(p tview.Primitive, width, height int) *tview.Flex {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

	form := tview.NewForm().
		AddInputField("Rename To", s, 30, nil, nil).
		AddButton("Accept", func() {
			pages.RemovePage("modal")
		}).
		AddButton("Cancel", func() {
			pages.RemovePage("modal")
		}).
		SetButtonsAlign(tview.AlignCenter)

	form.SetBorder(true).SetTitle("Rename resource").SetTitleAlign(tview.AlignCenter)
	modal := modalFunc(form, 40, 0)
	return modal
}

func RenderList() {
	frame := GetListFrame()
	pages.AddPage("rename", frame, true, true)
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func listSelectHandler(index int, _ string, _ string, shortcut rune) {
	if shortcut == 'q' {
		app.Stop()
		return
	}
	if tfStateList == nil {
		panic("tfstatelist is empty despite using handler")
	}
	tfStateSelected := tfStateList[index]
	modal := GetRenameView(tfStateSelected)
	pages.AddPage("modal", modal, true, true)
	pages.ShowPage("modal")
	//list.RemoveItem(index)
	//list.InsertItem(index, "[red]" + tfStateSelected, "", 0, nil)
	//list.SetCurrentItem(index)
}
