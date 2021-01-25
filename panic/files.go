package main

import(
	"fmt"
	"io/ioutil"
	"path/filepath"
)


func reportError() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(err)
	}
}

	
func scanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
		//return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
			//if err != nil {
		//		panic(err)
				//return err
		//	} else {
		//		fmt.Println(filePath)
		//	}
		}
	
	}
	//return nil
	
}

func main () {
	defer reportError()
	scanDirectory("/usr/local/go/src")
	//if err != nil {
//		panic(err)
//	}

}
