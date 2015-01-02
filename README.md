# vfmt

[![GoDoc](https://godoc.org/github.com/simleb/vfmt?status.svg)](http://godoc.org/github.com/simleb/vfmt)
[![Coverage Status](https://img.shields.io/coveralls/simleb/vfmt.svg)](https://coveralls.io/r/simleb/vfmt)
[![Build Status](https://drone.io/github.com/simleb/vfmt/status.png)](https://drone.io/github.com/simleb/vfmt/latest)

Package `vfmt` formats strings containing references to variables.

The syntax for variable replacement is `{variable name}` or `{variable name|format}`.

Variables are provided in a `map[string]interface{}`.

## Example

Given the string `pic_{id|%04d}_{mode}.jpg` and the variable map:

```go
map[string]interface{}{
	"id": 12,
	"mode": "box",
}
```

the formatted string is `pic_0012_box.jpg`.

## License

The MIT License (MIT)

	Copyright (c) 2015 Simon Leblanc
	
	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:
	
	The above copyright notice and this permission notice shall be included in
	all copies or substantial portions of the Software.
	
	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
	THE SOFTWARE.
