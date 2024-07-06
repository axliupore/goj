package file

import (
	"os"
	"strings"
)

type File struct {
	Name    string // 文件名
	Suffix  string // 后缀
	Content string // 文件内容
	Stem    string // 	去除文件后缀的文件名
}

var extMap = map[string]string{
	"input":  ".in",
	"output": ".out",
	"go":     ".go",
	"cpp":    ".cpp",
}

func GetFilesInCurrentDir() ([]*File, error) {
	var files []*File

	dir, err := os.Getwd()
	if err != nil {
		return nil, nil
	}

	dirId, err := os.Open(dir)
	if err != nil {
		return nil, nil
	}

	// 获取当前文件下的所有文件和目录
	fileInfos, err := dirId.Readdir(0)
	if err != nil {
		return nil, nil
	}

	for _, file := range fileInfos {
		if !filterFile(file) {
			continue
		}
		name := file.Name()

		content, err := fileContent(name)
		if err != nil {
			return nil, err
		}

		stem := fileStem(name)

		files = append(files, &File{
			Name:    name,
			Suffix:  fileExt(name),
			Content: content,
			Stem:    stem,
		})
	}

	return files, nil
}

// 过滤目录和其他文件，只保留已经实现的编程语言的文件和 in,out 文件
func filterFile(file os.FileInfo) bool {
	if file.IsDir() {
		return false
	}

	// 文件名
	filename := file.Name()
	// 过滤隐藏文件
	if strings.HasPrefix(filename, ".") {
		return false
	}

	// 文件后缀
	ext := fileExt(filename)

	for _, v := range extMap {
		if ext == v {
			return true
		}
	}

	return false
}

func fileExt(filename string) string {
	parts := strings.SplitAfter(filename, ".")
	if len(parts) > 1 {
		return "." + parts[len(parts)-1]
	}
	return ""
}

func fileContent(name string) (string, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		return "", nil
	}
	return string(content), nil
}

func fileStem(name string) string {
	return strings.TrimSuffix(name, fileExt(name))
}
