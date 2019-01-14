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

var cppHeaderTmpl = template.Must(template.New("t").Funcs(tmplFuncMap).Parse(cppHeaderTmplText))

const cppHeaderTmplText = `// This file is auto-generated by gml.
#ifndef GML_GEN_CPP_{{.PackageName}}_H
#define GML_GEN_CPP_{{.PackageName}}_H

#include "../gen_c/{{.PackageName}}.h"

#include <stdlib.h>
#include <stdint.h>
#include <iostream>
#include <gml.h>

#include <QObject>
#include <QByteArray>
#include <QString>
#include <QChar>
#include <QVariant>

{{/* Struct loop */ -}}
{{range $struct := .Structs -}}
//###
//### {{$struct.Name}}
//###

class {{$struct.CPPBaseName}} : public QObject
{
    Q_OBJECT

private:
    void* goPtr;

public:
    {{$struct.CPPBaseName}}(void* goPtr);

{{/* Signals */ -}}
signals:
{{- range $signal := $struct.Signals }}
    void {{$signal.CPPName}}({{cppParams $signal.Params true true}});
{{- end}}

{{/* Slots */ -}}
public slots:
{{- range $slot := $struct.Slots }}
    void {{$slot.CPPName}}({{cppParams $slot.Params true true}});
{{end}}
};

{{- /* End of struct loop */ -}}
{{- end}}

#endif
`

var cppSourceTmpl = template.Must(template.New("t").Funcs(tmplFuncMap).Parse(cppSourceTmplText))

const cppSourceTmplText = `// This file is auto-generated by gml.
#include "{{.PackageName}}.h"

{{/* Struct loop */ -}}
{{range $struct := .Structs -}}
//###
//### {{$struct.Name}}
//###

{{$struct.CBaseName}} {{$struct.CBaseName}}_new(void* go_ptr) {
    auto _vv = new {{$struct.CPPBaseName}}(go_ptr);
    return (void*)_vv;
}

void {{$struct.CBaseName}}_free({{$struct.CBaseName}} _v) {
    if (_v == NULL) return;
    auto _vv = ({{$struct.CPPBaseName}}*)_v;
    delete _vv;
    _v = NULL;
}

{{$struct.CPPBaseName}}::{{$struct.CPPBaseName}}(void* goPtr) :
    goPtr(goPtr)
{}

{{/* Signals */ -}}
{{- range $signal := $struct.Signals }}
void {{$struct.CBaseName}}_{{$signal.Name}}({{$struct.CBaseName}} _v{{cParams $signal.Params true false}}) {
    auto _vv = ({{$struct.CPPBaseName}}*)_v;
    emit _vv->{{$signal.CPPName}}({{cToCPPParams $signal.Params true}});
}
{{end}}

{{- /* Slots */ -}}
{{- range $slot := $struct.Slots }}
{{$struct.CBaseName}}_{{$slot.Name}}_cb_t {{$struct.CBaseName}}_{{$slot.Name}}_cb = NULL;

void {{$struct.CBaseName}}_{{$slot.Name}}_cb_register({{$struct.CBaseName}}_{{$slot.Name}}_cb_t cb) {
    {{$struct.CBaseName}}_{{$slot.Name}}_cb = cb;
}

void {{$struct.CPPBaseName}}::{{$slot.CPPName}}({{cppParams $slot.Params true true}}) {
    try {
        {{$struct.CBaseName}}_{{$slot.Name}}_cb(this->goPtr{{cppToCParams $slot.Params false}});
    }
    catch (std::exception& e) {
        std::cerr << "gml: catched slot exception: " << e.what() << std::endl;
    }
    catch (...) {
        std::cerr << "gml: catched slot exception" << std::endl;
    }
}
{{end}}

{{- /* End of struct loop */ -}}
{{- end}}
`
