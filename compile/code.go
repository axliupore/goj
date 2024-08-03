package compile

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/axliupore/goj/file"
)

type CodeRunner struct {
}

func (c *CodeRunner) Compile(file *file.File, compileName string, args ...string) (string, error) {
	exeFile := file.Stem + "_exec"
	args = append(args, "-o", exeFile, file.Name)
	compileCmd := exec.Command(compileName, args...)

	// 创建一个缓冲区来捕获标准错误输出
	var stderr bytes.Buffer
	compileCmd.Stderr = &stderr

	if err := compileCmd.Run(); err != nil {
		return "", errors.New(stderr.String())
	}

	exePath, err := filepath.Abs(exeFile)
	if err != nil {
		return "", err
	}
	return exePath, nil
}

func (g *CodeRunner) Run(exePath string, input string, args ...string) (string, error) {
	runCmd := exec.Command(exePath, args...)

	// 创建标准输入管道
	stdin, err := runCmd.StdinPipe()
	if err != nil {
		return "", err
	}

	// 获取命令的输出
	var stdoutBuf bytes.Buffer
	runCmd.Stdout = &stdoutBuf
	runCmd.Stderr = &stdoutBuf

	// 启动命令
	if err := runCmd.Start(); err != nil {
		return "", err
	}

	// 向标准输入写入数据
	if _, err := stdin.Write([]byte(input)); err != nil {
		return "", err
	}
	stdin.Close()

	// 等待命令执行完毕
	if err := runCmd.Wait(); err != nil {
		return "", err
	}

	output := strings.TrimSpace(stdoutBuf.String())
	return output, nil
}

func (c *CodeRunner) Remove(exePath string) {
	parts := strings.Split(exePath, "/")
	if len(parts) == 0 {
		return
	}
	name := parts[len(parts)-1]

	dir, _ := os.Getwd()

	files, _ := file.GetFilesDir(dir)
	for _, f := range files {
		if strings.HasPrefix(f.Name(), name) {
			// 删除所有以 name 开头的文件
			_ = os.Remove(f.Name())
		}
	}
}
