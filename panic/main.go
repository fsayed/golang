package main

import(
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"

)

func reportPanic() {
	p := recover()
	fmt.Println(reflect.TypeOf(p))
	if p == nil {
		return
	}
	//note that if the panic value can't be converted to error type, it won't report it
	//panic("not good") <--- this is string type, it can't be converetedto error type
	//in that case, renew panic state in the else clause
	err, ok := p.(error)
	if ok  {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func scanDirectory(path string) {
	fmt.Println(path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}


func main() {
	defer reportPanic()
	scanDirectory("temp")
}
