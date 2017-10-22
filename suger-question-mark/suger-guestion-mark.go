package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
)

type InputFile struct {
	PackageName string
	BuildTarget string
}

type Pos struct {
	Line int
	Column int
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


	replaceLine := []Pos{}


	for _, commentGroup := range f.Comments {
		print("hello ")
		//fmt.Println(spew.Sprint(commentGroup))
		for _, 	comment := range  commentGroup.List{
			println(comment.Text)
			if strings.HasPrefix(comment.Text, "? "){
				replaceLine = append(replaceLine, Pos{Line:fset.Position(comment.Slash).Line,
					Column:fset.Position(comment.Slash).Column,
				})
				println(spew.Sdump(fset.Position(comment.Slash)))
			}

		}
	}
	println(spew.Sdump(replaceLine))
	replace(path, replaceLine)
	return nil
}

func replace(path string, replaceLine []Pos) error {
	newFilename := fmt.Sprintf("%s-sugered.go", strings.Split(filepath.Base(path), ".")[0])
	newPath := filepath.Join(filepath.Dir(path), newFilename)
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for _,pos := range replaceLine {
		l := lines[pos.Line]
		index := ""
		for i := 0; i < pos.Column-1; i++ {
			index = fmt.Sprintf("%s%c", index,l[i])
		}
		lines[pos.Line] = index + "if err != nil {"
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(newPath, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
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
