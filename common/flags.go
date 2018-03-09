package common

import (
	"flag"
	"strings"
	"github.com/GUMI-golang/example/example"
	"fmt"
	"github.com/pkg/errors"
	"github.com/GUMI-golang/gumi"
)

var (
	Example = flag.String("example", "HelloWorld", "[" + strings.Join(example.List(), ", ") + "]")
	Size = flag.String("size", "VGA", "Set ScreenSize")
)
var (
	width, height int
	ex gumi.GUMI
)


func CheckFlags() error {
	flag.Parse()
	width, height = ParseSize(*Size)
	ex = example.Select(*Example)
	// check setting
	if ex == nil{
		return errors.New(fmt.Sprintf("Unknown example : %s", *Example))
	}
	if width == 0 || height == 0{
		return errors.New(fmt.Sprintf("Invalid Size : %s", *Size))
	}
	return nil
}

func GetSize() (w, h int) {
	return width, height
}

func GetExample() (gumi.GUMI) {
	return ex
}