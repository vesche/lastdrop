package main

import (
	"github.com/rivo/tview"
)

func main() {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().SetText(text)
	}
	a := newPrimitive("abc")
	b := newPrimitive("xyz")

	grid := tview.NewGrid().
		SetRows().
		SetColumns().
		SetBorders(true)

	grid.AddItem(a, 0, 0, 1, 1, 0, 100, false).
		AddItem(b, 0, 1, 1, 1, 0, 100, false)

	if err := tview.NewApplication().SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
