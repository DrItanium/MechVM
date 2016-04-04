// main function
package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func versionStr() string {
	return "donuts"
}
func main() {
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_JOYSTICK); err != nil {
		fmt.Println(err)
		return
	}
	_, err := sdl.CreateWindow(versionStr(), 0, 0, 0, 0, sdl.WINDOW_OPENGL|sdl.WINDOW_RESIZABLE)
	if err != nil {
		fmt.Println(err)
		return
	}
}
