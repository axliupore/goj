package main

import (
	"log"
	"os"

	"github.com/axliupore/goj/compile"
	"github.com/axliupore/goj/file"
)

func main() {
	files, err := file.GetFilesInCurrentDir()
	if err != nil {
		log.Fatalln(err)
	}
	fs, ins, _ := fileTypes(files)

	var outs []*file.File

	g := compile.GoRunner{}
	c := compile.CppRunner{}

	goFiles := make([]*file.File, 0)
	cppFiles := make([]*file.File, 0)

	for _, f := range fs {
		if f.Suffix == g.Name() {
			goFiles = append(goFiles, f)
		} else if f.Suffix == c.Name() {
			cppFiles = append(cppFiles, f)
		}

	}

	goOuts, err := goRunner(g, goFiles, ins)
	if err != nil {
		log.Fatalln(err)
	}
	outs = append(outs, goOuts...)

	cppOuts, err := cppRunner(c, cppFiles, ins)
	if err != nil {
		log.Fatalln(err)
	}
	outs = append(outs, cppOuts...)

	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range outs {
		err := saveOutputFile(f)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func saveOutputFile(file *file.File) error {
	f, err := os.Create(file.Name)
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(file.Content))
	if err != nil {
		return err
	}
	return nil
}

func fileTypes(files []*file.File) ([]*file.File, []*file.File, []*file.File) {
	fs := make([]*file.File, 0)
	ins := make([]*file.File, 0)
	outs := make([]*file.File, 0)
	for _, file := range files {
		if file.Suffix == ".in" {
			ins = append(ins, file)
		} else if file.Suffix == ".out" {
			outs = append(outs, file)
		} else {
			fs = append(fs, file)
		}
	}
	return fs, ins, outs
}

func goRunner(g compile.GoRunner, fs, ins []*file.File) ([]*file.File, error) {
	outs, err := g.Exec(fs, ins)
	if err != nil {
		return nil, err
	}
	return outs, nil
}

func cppRunner(c compile.CppRunner, fs, ins []*file.File) ([]*file.File, error) {
	outs, err := c.Exec(fs, ins)
	if err != nil {
		return nil, err
	}
	return outs, nil
}
