/*
 * GML - Go QML
 *
 * The MIT License (MIT)
 *
 * Copyright (c) 2019 Roland Singer <roland.singer[at]desertbit.com>
 * Copyright (c) 2019 Sebastian Borchers <sebastian[at]desertbit.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package build

import "text/template"

var tmplFuncMap = template.FuncMap{
	"goParams":     tmplFuncGoParams,
	"cParams":      tmplFuncCParams,
	"cppParams":    tmplFuncCPPParams,
	"goToCParams":  tmplFuncGoToCParams,
	"cToCPPParams": tmplFuncCToCPPParams,
}

func tmplFuncGoParams(params []*genParam, withType, skipFirstComma bool, optPrefix ...string) (s string) {
	var prefix string
	if len(optPrefix) > 0 {
		prefix = optPrefix[0]
	}

	for i, p := range params {
		if !skipFirstComma || i != 0 {
			s += ", "
		}
		s += prefix + p.Name
		if withType {
			s += " " + p.Type
		}
	}
	return
}

func tmplFuncCParams(params []*genParam, withType, skipFirstComma bool) (s string) {
	for i, p := range params {
		if !skipFirstComma || i != 0 {
			s += ", "
		}
		if withType {
			s += p.CType + " "
		}
		s += p.Name
	}
	return
}

func tmplFuncCPPParams(params []*genParam, withType, skipFirstComma bool) (s string) {
	for i, p := range params {
		if !skipFirstComma || i != 0 {
			s += ", "
		}
		if withType {
			s += p.CPPType + " "
		}
		s += p.Name
	}
	return
}

func tmplFuncGoToCParams(params []*genParam, prefix string, optsIndent ...int) (s string) {
	var ident string
	if len(optsIndent) > 0 {
		for i := 0; i < optsIndent[0]; i++ {
			ident += " "
		}
	}

	addLine := func(l string) {
		s += "\n" + ident + l
	}

	for _, p := range params {
		cName := prefix + p.Name

		switch p.Type {
		case "bool":
			addLine("var " + cName + " C.u_int8_t")
			addLine("if " + p.Name + " { " + cName + " = 1 }")
		case "byte":
			addLine(cName + " := C.char(" + p.Name + ")")
		case "string":
			addLine(cName + " := C.CString(" + p.Name + ")")
			addLine("defer C.free(unsafe.Pointer(" + cName + "))")
		case "rune":
			addLine(cName + " := C.int32_t(" + p.Name + ")")

		case "float32":
			addLine(cName + " := C.float(" + p.Name + ")")
		case "float64":
			addLine(cName + " := C.double(" + p.Name + ")")

		case "int":
			addLine(cName + " := C.int(" + p.Name + ")")
		case "int8":
			addLine(cName + " := C.int8_t(" + p.Name + ")")
		case "uint8":
			addLine(cName + " := C.u_int8_t(" + p.Name + ")")
		case "int16":
			addLine(cName + " := C.int16_t(" + p.Name + ")")
		case "uint16":
			addLine(cName + " := C.u_int16_t(" + p.Name + ")")
		case "int32":
			addLine(cName + " := C.int32_t(" + p.Name + ")")
		case "uint32":
			addLine(cName + " := C.u_int32_t(" + p.Name + ")")
		case "int64":
			addLine(cName + " := C.int64_t(" + p.Name + ")")
		case "uint64":
			addLine(cName + " := C.u_int64_t(" + p.Name + ")")

		default:
			addLine(cName + " := (C.gml_variant)(gml.ToVariant(" + p.Name + ").Pointer())")
		}
	}
	return
}

func tmplFuncCToCPPParams(params []*genParam, skipFirstComma bool) (s string) {
	for i, p := range params {
		if !skipFirstComma || i != 0 {
			s += ", "
		}

		switch p.Type {
		case "bool":
			s += "bool(" + p.Name + ")"
		case "byte":
			s += p.Name
		case "string":
			s += "QString(" + p.Name + ")"
		case "rune":
			s += "QChar(" + p.Name + ")"

		case "float32":
			s += p.Name
		case "float64":
			s += p.Name

		case "int":
			s += p.Name
		case "int8":
			s += p.Name
		case "uint8":
			s += p.Name
		case "int16":
			s += p.Name
		case "uint16":
			s += p.Name
		case "int32":
			s += p.Name
		case "uint32":
			s += p.Name
		case "int64":
			s += p.Name
		case "uint64":
			s += p.Name

		default:
			s += "QVariant(*((QVariant*)" + p.Name + "))" // Create a copy of the passed QVariant. The old QVariant will be deleted by Go.
		}
	}
	return
}
