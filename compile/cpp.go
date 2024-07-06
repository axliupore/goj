package compile

import (
	"github.com/axliupore/goj/file"
)

type CppRunner struct {
	codeRunner CodeRunner
}

func (c *CppRunner) Name() string {
	return ".cpp"
}

func (c *CppRunner) Exec(files, inFiles []*file.File) ([]*file.File, error) {
	var outFiles []*file.File
	for _, f := range files {
		for _, inFile := range inFiles {
			exePath, err := c.Compile(f, "g++")
			if err != nil {
				return nil, err
			}

			output, err := c.Run(exePath, inFile.Content)
			if err != nil {
				return nil, err
			}

			name := f.Stem + "-" + inFile.Stem + "-cpp"
			outFiles = append(outFiles, &file.File{
				Name:    name + ".out",
				Suffix:  ".out",
				Content: output,
				Stem:    name,
			})

			c.codeRunner.Remove(exePath)
		}
	}
	return outFiles, nil
}

func (c *CppRunner) Compile(file *file.File, compileName string, args ...string) (string, error) {
	return c.codeRunner.Compile(file, compileName, args...)
}

func (c *CppRunner) Run(exePath string, input string, args ...string) (string, error) {
	return c.codeRunner.Run(exePath, input, args...)
}
