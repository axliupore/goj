package template

import (
	"fmt"
	"github.com/axliupore/goj/file"
	"os"
	"path/filepath"
	"strings"
)

// Output 复制源目录中的以 language 结尾的文件内容到目标目录中的对应文件
func Output(language string) {
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 获取源目录中的文件
	sourceDir := "/Users/axliu/code/algorithm/template"
	tfsFiles, err := file.GetFilesDir(sourceDir)
	if err != nil {
		panic(err)
	}

	// 获取目标目录中的文件
	fsFiles, err := file.GetFilesDir(dir)
	if err != nil {
		panic(err)
	}

	// 复制内容
	for _, tfsFile := range tfsFiles {
		if strings.HasSuffix(tfsFile.Name(), language) { // 检查文件扩展名
			content, err := os.ReadFile(filepath.Join(sourceDir, tfsFile.Name()))
			if err != nil {
				fmt.Printf("Error reading source file %s: %v\n", tfsFile.Name(), err)
				continue
			}

			for _, fsFile := range fsFiles {
				if strings.HasSuffix(fsFile.Name(), language) { // 检查文件扩展名
					if err := os.WriteFile(filepath.Join(dir, fsFile.Name()), content, 0644); err != nil {
						fmt.Printf("Error writing to target file %s: %v\n", fsFile.Name(), err)
					} else {
						fmt.Printf("Copied content from %s to %s\n", tfsFile.Name(), fsFile.Name())
					}
				}
			}
		}
	}
}
