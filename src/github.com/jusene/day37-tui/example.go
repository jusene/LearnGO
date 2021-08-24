package main

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
)

func main() {
	/*box := tui.NewVBox(
		tui.NewLabel("tui-go"),
		tui.NewButton("test"),
	)
	*/
	boxEntry := tui.NewEntry()
	boxEntry.SetText("hello world")
	boxEntry.OnSubmit(func(entry *tui.Entry) {
		fmt.Println(entry.Text())
	})
	//box := tui.NewHBox(tui.NewButton("test"), tui.NewButton("TEST"))
	//ui, err := tui.New(box)

	urlBox := tui.NewHBox(boxEntry)
	urlBox.SetBorder(true)
	urlBox.SetTitle("URL")

	ui, err := tui.New(urlBox)

	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
