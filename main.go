package main

import (
	myjs "syscall/js"

	"github.com/robertkrimen/otto"
)

var document = myjs.Global().Get("document")

func getElementByID(id string) myjs.Value {
	return document.Call("getElementById", id)
}

func renderEditor(parent myjs.Value) myjs.Value {
	editorMarkup := `
		<div id="editor" style="display: flex; flex-flow: row wrap;">
			<textarea id="input" style="width: 50%; height: 400px"></textarea>
			<div id="preview" style="width: 50%;"></div>
		</div>
	`
	parent.Call("insertAdjacentHTML", "beforeend", editorMarkup)
	return getElementByID("editor")
}

func main() {
	quit := make(chan struct{}, 0)

	// See example 2: Enable the stop button
	stopButton := getElementByID("stop")
	stopButton.Set("disabled", false)
	stopButton.Set("onclick", myjs.NewCallback(func([]myjs.Value) {
		println("stopping")
		stopButton.Set("disabled", true)
		quit <- struct{}{}
	}))

	editor := renderEditor(document.Get("body"))
	preview := getElementByID("preview")

	input := getElementByID("input")

	vm := otto.New()

	// renderButton := getElementByID("render")
	input.Set("oninput", myjs.NewCallback(func([]myjs.Value) {
		v, _ := vm.Run(input.Get("value").String())
		s, _ := v.ToString()

		preview.Set("innerHTML", s)
	}))

	<-quit
	editor.Call("remove")
}
