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

package gml

// #include <gml.h>
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type Image struct {
	freed bool
	img   C.gml_image
}

func NewImageFromData(data []byte) (img *Image, err error) {
	imgC := C.gml_image_new_from_data((*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)))
	return newImage(imgC, true)
}

func newImage(imgC C.gml_image, free bool) (img *Image, err error) {
	img = &Image{
		freed: !free,
		img:   imgC,
	}

	// Always free the C++ value if defined so.
	if free {
		runtime.SetFinalizer(img, freeImage)
	}

	// TODO:
	// Check if something failed.
	// This should never happen is signalizes a fatal error.
	if img.img == nil {
		err = fmt.Errorf("failed to create gml image: C pointer is nil")
	}
	return
}

func freeImage(img *Image) {
	if img.freed {
		return
	}
	img.freed = true
	C.gml_image_free(img.img)
}

func (img *Image) Free() {
	freeImage(img)
}
