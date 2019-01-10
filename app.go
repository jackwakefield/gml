/*
 *  GML - Go QML
 *  Copyright (c) 2019 Roland Singer [roland.singer@deserbit.com]
 *  Copyright (c) 2019 Sebastian Borchers [sebastian@deserbit.com]
 */

package gml

// #include <gml.h>
import "C"
import (
	"os"
	"runtime"
	"unsafe"
)

func init() {
	// Ensure the thread is locked within the main thread.
	runtime.LockOSThread()
}

// TODO: free
type App struct {
	app  C.gml_app
	argc int
	argv **C.char // TODO: free
}

func NewApp() (a *App, err error) {
	a = &App{
		argc: len(os.Args), // TODO: really pass all arguments?
		argv: toCharArray(os.Args),
	}
	a.app = C.gml_app_new(C.int(a.argc), a.argv)
	return
}

// Exec executes the application and returns the exit code.
// This method is blocking.
func (a *App) Exec() int {
	return int(C.gml_app_exec(a.app))
}

// Quit the application.
func (a *App) Quit() int {
	return int(C.gml_app_quit(a.app))
}

// Load the root QML file located at url.
func (a *App) Load(url string) error {
	urlC := C.CString(url)
	defer C.free(unsafe.Pointer(urlC))

	// TODO:
	_ = int(C.gml_app_load(a.app, urlC))
	return nil
}

// Load the QML given in data.
func (a *App) LoadData(data string) error {
	dataC := C.CString(data)
	defer C.free(unsafe.Pointer(dataC))

	// TODO:
	_ = int(C.gml_app_load_data(a.app, dataC))
	return nil
}