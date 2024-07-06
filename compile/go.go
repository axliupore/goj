package compile

import (
	"github.com/axliupore/goj/file"
)

type GoRunner struct {
	codeRunner CodeRunner
}

func (g *GoRunner) Name() string {
	return ".go"
}

func (g *GoRunner) Exec(files, inFiles []*file.File) ([]*file.File, error) {
	var outFiles []*file.File
	for _, f := range files {
		for _, inFile := range inFiles {
			exePath, err := g.Compile(f, "go", "build")
			if err != nil {
				return nil, err
			}

			output, err := g.Run(exePath, inFile.Content)
			if err != nil {
				return nil, err
			}

			name := f.Stem + "-" + inFile.Stem + "-go"
			outFiles = append(outFiles, &file.File{
				Name:    name + ".out",
				Suffix:  ".out",
				Content: output,
				Stem:    name,
			})

			g.codeRunner.Remove(exePath)
		}
	}
	return outFiles, nil
}

func (g *GoRunner) Compile(file *file.File, compileName string, args ...string) (string, error) {
	return g.codeRunner.Compile(file, compileName, args...)
}

func (g *GoRunner) Run(exePath string, input string, args ...string) (string, error) {
	return g.codeRunner.Run(exePath, input, args...)
}
