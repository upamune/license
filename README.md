# license
license command can generate LICENSE easily.

## Usage

```sh
license -n upamune -y 2015 -f mit -o LICENSE
```

You can get ```LICENSE``` file.

```plain
The MIT License (MIT)

Copyright (c) 2015 upamune

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
```

### Option

|Flag|Mean|Required|Default|
|---:|---:|:--:|:--:|
|f (format)|format|o|-|
|n (name)|copyright holders name|x|Name|
|y (year)|year|x|Current Year|
|o (output)|save file name|x|STDOUT|

### License Format

You can use 7 formats license.

|license|name|
|-:|-:|
|MIT License|mit|
|BSD 2-Clause License|bsd2|
|BSD 3-Clause License|bsd3|
|GPL 3|gpl|
|LGPL|lgpl|
|Apache Software License 2.0|asl|
|CPL|cpl|

## Install

To install, use `go get`:

```bash
$ go get -d github.com/upamune/license
```

Or you can use homebrew.

```sh
$ brew install upamune/license/license
```


## Contribution

1. Fork ([https://github.com/upamune/license/fork](https://github.com/upamune/license/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[upamune](https://github.com/upamune)
