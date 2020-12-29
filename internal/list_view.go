package internal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"tfui/internal/tf"
)

var app = tview.NewApplication()
var pages = tview.NewPages()
var list = tview.NewList()
var tfStateList []tf.StateOp = nil

func GetListFrame() *tview.Frame {
	tfStateList = tf.GetStateList()

	list.ShowSecondaryText(false).
		SetSelectedFunc(listSelectHandler)
	for _, el := range tfStateList {
		list.AddItem(el.OriginalName, "", 0, nil)
	}
	list.AddItem("Quit", "Press to exit", 'q', nil)

	frame := tview.NewFrame(list).
		AddText("Rename State", true, tview.AlignCenter, tcell.ColorWhite)
	return frame
}

func GetRenameView(index int, s *tf.StateOp) *tview.Flex {
	modalFunc := func(p tview.Primitive, width, height int) *tview.Flex {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

	form := tview.NewForm()
	form.
		AddInputField("Rename To", s.OriginalName, 30, nil, func(text string) {
			s.NewName = text
			s.Op = tf.RenameOp
		}).
		AddButton("Accept", func() {
			pages.RemovePage("modal")
			refreshByIndex(index)
		}).
		AddButton("Cancel", func() {
			pages.RemovePage("modal")
			refreshByIndex(index)
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

func refreshByIndex(index int) {
	tfStateAtIndex := tfStateList[index]
	list.RemoveItem(index)
	list.InsertItem(index, "[red]"+tfStateAtIndex.OriginalName+" -> [green] "+tfStateAtIndex.NewName, "", 0, nil)
	list.SetCurrentItem(index)
}

func listSelectHandler(index int, _ string, _ string, shortcut rune) {
	if shortcut == 'q' {
		app.Stop()
		return
	}
	if tfStateList == nil {
		panic("tfstatelist is empty despite using handler")
	}
	modal := GetRenameView(index, &tfStateList[index])
	pages.AddPage("modal", modal, true, true)
	pages.ShowPage("modal")
}
