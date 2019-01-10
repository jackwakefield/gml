/*
 *  GML - Go QML
 *  Copyright (c) 2019 Roland Singer [roland.singer@deserbit.com]
 *  Copyright (c) 2019 Sebastian Borchers [sebastian@deserbit.com]
 */

package main

import (
	"log"
	"os"

	"github.com/desertbit/gml"
)

type Bridge struct {
	_ struct {
		State     int               `gml:"property"`
		Connect   func(addr string) `gml:"slot"`
		Connected func()            `gml:"signal"`
	}
}

func (b *Bridge) Connect(addr string) {

}

func main() {
	app, err := gml.NewApp()
	if err != nil {
		log.Fatalln(err)
	}

	err = app.Load("qml/main.qml")
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(app.Exec())
}