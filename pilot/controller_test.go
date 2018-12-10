package pilot

import (
	"fmt"
	"os"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

func TestRun(t *testing.T) {
	var window *sdl.Window
	var winTitle string = "test"
	var winWidth int32 = 640
	var winHeight int32 = 400
	var err error
	//var ctrl Controller
	//var ch chan data.Event

	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return
	}
	defer window.Destroy()
	/*
		ch = make(chan data.Event)
		defer close(ch)

		//go func(ch chan<- data.Event) { ctrl.Run(ch) }(ch)

		go func(ch <-chan data.Event) {
			var evt data.Event
			for {
				fmt.Println("TEST")
				//evt = <-ch
				//fmt.Println(evt.Msg)
			}
		}(ch)
	*/
}
