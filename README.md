# goj
在写算法的时候，编译、运行、比对输出结果这些操作是很繁琐的，本项目就是为了简化操作而进行的，只需在需要进行编译运行的目录下执行 `goj` 命令，即可进行一系列操作（类 Uinx 系统） 。

## 实现语言
- [x] Go
- [x] Cpp

## 使用教程
1. 克隆本项目
2. 在项目的根目录下执行 `go build` 命令
3. 将要输入的测试文件创建一个以 `.in` 结尾的文件
4. 在当前目录下执行 `goj` 命令，即可获取编译运行后的文件

## 注意
> 需要将编译好的可执行文件加入到系统的环境变量当中去 