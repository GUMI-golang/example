package example

import (
	"github.com/GUMI-golang/gumi"
	"reflect"
	"runtime"
	"strings"
)
var (
	funcs = []func() gumi.GUMI{
		HelloWorld,
		HelloButton,
		HelloImage,
		Buttons,
		Modal,
		Progress,
		Radios,
		Toggles,
		Editing,
	}
	names []string
)

func init() {
	for _, v := range funcs{
		temp := strings.FieldsFunc(runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name(), func(r rune) bool {
			if r == '/'{
				return true
			}
			return false
		})
		name := temp[len(temp) - 1]
		names = append(names, strings.Split(name, ".")[1])
	}
}

func List() []string {
	return names
}
func Select(name string) gumi.GUMI{
	var debug gumi.GUMI = nil
	var ex gumi.GUMI
	var cmd = name

	if strings.HasPrefix(name, "debug"){
		temp := strings.FieldsFunc(name, func(r rune) bool {
			return r == '.'
		})
		if len(temp) < 2{
			return nil
		}
		debug = debugAppend(temp[0])
		cmd = temp[1]
	}
	for i, v := range names{
		if v == cmd{
			ex = funcs[i]()
			break
		}
	}
	if debug != nil{
		ex = gumi.LinkingFrom(debug, ex)
	}
	return ex
}
