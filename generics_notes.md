# How to build the most recent dev implementations of type parameters:
https://www.jetbrains.com/help/go/how-to-use-type-parameters-for-generic-programming.html

```bash
mkdir -p ~/temp/
cd ~/temp
git clone https://go.googlesource.com/go goroot
cd goroot
git checkout dev.typeparams
cd src
./all.bash

## blah blah... compile compile...:
# ALL TESTS PASSED
# ---
# Installed Go for darwin/arm64 in /Users/rich/temp/goroot
# Installed commands in /Users/rich/temp/goroot/bin
# *** You need to add /Users/rich/temp/goroot/bin to your PATH.
# ```

# target the dev build of code from within VS Code:
cd ~/some_go_folder
GOROOT=/Users/rich/temp/goroot code .

# How to use generics to specify the return type of a method:
# (However, type parameters are not allowed in methods!)
```go
package main

import (
	"fmt"
)

// AllInteger is a type constriant that restrics an allowed geric type to only integers
type AllInteger interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
}

func min[T AllInteger](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func getsix[T AllInteger]() (T, error) {
	return 6, nil
}

func main() {
	var i1 uint32 = 3
	var i2 uint32 = 4
	m := min(i1,i2)
	fmt.Printf("%T\n",m)
	fmt.Println(min(i1, i2))

	zzz, _ := getsix[uint64]()
	fmt.Printf("%d,%T\n",zzz,zzz)
}
```