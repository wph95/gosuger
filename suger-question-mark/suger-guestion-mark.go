package main

import (
	"fmt"
	//"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type InputFile struct {
	PackageName string
	BuildTarget string
}

// Init initializes an InputFile from a path.
func (i *InputFile) Init(path string) error {

	// set the build target
	if strings.HasSuffix(path, ".go") {
		root := strings.TrimSuffix(path, ".go")
		dir, file := filepath.Split(root)
		(*i).BuildTarget = filepath.Join(dir, fmt.Sprintf("%s_tabler.go", file))
	} else {
		return fmt.Errorf("File '%s' is not a Go file.", path)
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(
		fset,
		path,
		nil,
		parser.ParseComments,
	)


	if err != nil {
		fmt.Errorf("Unable to parse '%s': %s", path, err)
	}

	//ast.Print(fset, f)

	// get package name
	//if f.Name != nil {
	//	(*i).PackageName = f.Name.Name
	//} else {
	//	fmt.Errorf("Missing package name in '%s'", path)
	//}

	// build list of tables
	for _, commentGroup := range f.Comments {
		print("hello ")
		//fmt.Println(spew.Sprint(commentGroup))
		for _, 	comment := range  commentGroup.List{
			println(comment.Text)
			if strings.HasPrefix(comment.Text, ""){
				println("mingzhong")
			}
		}
	}

	return nil
}


func main() {
	for _, path := range os.Args[1:] {
		infile := InputFile{}
		if err := infile.Init(path); err != nil {
			log.Printf("%v", err)
			continue
		}
	}
}
