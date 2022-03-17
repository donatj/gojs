package main

import (
	"syscall/js"

	"github.com/robertkrimen/otto"
)

var document = js.Global().Get("document")

func getElementByID(id string) js.Value {
	return document.Call("getElementById", id)
}

func renderEditor(parent js.Value) js.Value {
	editorMarkup := `
		<div id="editor" style="display: flex; flex-flow: row wrap;">
			<textarea id="input" style="width: 49vw; height: 400px" placeholder="type JS, it will run in real time">
var x = 1;
var y = 2;

"outputs:\n" + (x + y); // final line's return is value of output
</textarea></td>
			<output id="preview" style="width: 49vw; background: gray">outputs: 3</output>
		</div>
	`
	parent.Call("insertAdjacentHTML", "beforeend", editorMarkup)
	return getElementByID("editor")
}

func main() {
	quit := make(chan struct{}, 0)

	runButton := getElementByID("runButton")
	runButton.Set("style", "display: none")

	editor := renderEditor(document.Get("body"))
	preview := getElementByID("preview")

	input := getElementByID("input")

	vm := otto.New()

	// renderButton := getElementByID("render")
	input.Set("oninput", js.FuncOf(func(js.Value, []js.Value) any {
		v, _ := vm.Run(input.Get("value").String())
		s, _ := v.ToString()

		preview.Set("textContent", s)

		return nil
	}))

	<-quit
	editor.Call("remove")
}
