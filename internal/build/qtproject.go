/*
 *  GML - Go QML
 *  Copyright (c) 2019 Roland Singer [roland.singer@deserbit.com]
 *  Copyright (c) 2019 Sebastian Borchers [sebastian@deserbit.com]
 */

package build

import (
	"html/template"
	"os"
)

func prepareQtProject(ctx *Context) (err error) {
	tmpl, err := template.New("t").Parse(qtProData)
	if err != nil {
		return
	}

	// Create or open the config file.
	f, err := os.Create(ctx.QtProFile)
	if err != nil {
		return
	}
	defer func() {
		err = f.Close()
	}()

	return tmpl.Execute(f, &ctx)
}

const qtProData = `
QT += core qml quick

TEMPLATE = lib
CONFIG += staticlib

HEADERS += {{.GMLBindingDir}}/headers/*.h
SOURCES += {{.GMLBindingDir}}/sources/*.cpp

OBJECTS_DIR = {{.BuildDir}}
MOC_DIR = {{.BuildDir}}
UI_DIR = {{.BuildDir}}
TARGET = {{.BuildDir}}/gml
`