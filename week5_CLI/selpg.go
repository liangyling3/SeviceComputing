package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

// 用于保存参数的结构体
type selpgArgs struct {
	start_page  int    // 起始页
	end_page    int    // 结束页
	page_length int    // 页行数固定的文本
	page_type   bool   // 页由换行符定界的文本
	dest        string // 将选定的页直接发送至打印机
	filename    string // 输入文件名
}

// 初始化参数
func getArgs(args *selpgArgs) {
	// 参数值的绑定
	pflag.IntVarP(&args.start_page, "start_page", "s", -1, "Start page of file")
	pflag.IntVarP(&args.end_page, "end_page", "e", -1, "End page of file")
	pflag.IntVarP(&args.page_length, "page_length", "l", 72, "Number of rows in one page")
	pflag.BoolVarP(&args.page_type, "page_type", "f", false, "Flag splits page")
	pflag.StringVarP(&args.dest, "dest", "d", "", "name of printer")

	// 打印到屏幕上
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: selpg -s startPage -e endPage [-l linePerPage | -f ][-d dest] filename\n\n")
		pflag.PrintDefaults()
	}
	// 标识和参数的解析
	pflag.Parse()

}

// 检查参数
func checkArgs(args *selpgArgs) {
	if args.start_page < 1 || args.end_page < 1 || args.start_page > args.end_page {
		// 输出命令行的提示信息
		pflag.Usage()
	}

	if args.start_page < 1 || args.end_page < 1 {
		fmt.Fprintf(os.Stderr, "\n[Error] The start page and end page are required and must be bigger than 0!\n")
		os.Exit(1)
	}

	if args.start_page > args.end_page {
		fmt.Fprintf(os.Stderr, "[Error] The start page number can't be bigger than the end page number!\n")
		os.Exit(2)
	}

	if (args.page_type == true) && (args.page_length != 72) {
		fmt.Fprintf(os.Stderr, "\n[Error]The command -l and -f are exclusive!\n")
		os.Exit(3)
	}

	if args.page_length < 1 {
		fmt.Fprintf(os.Stderr, "[Error] The page length must be bigger than 1!\n")
		os.Exit(4)
	}

	// Narg()返回解析后参数
	if pflag.NArg() > 0 {
		args.filename = pflag.Arg(0)
		_, err := os.Stat(args.filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[Error]Wrong filepath!\n")
			os.Exit(5)
		}
	}
}

// 处理分页
func handle(args *selpgArgs) {
	filein := os.Stdin
	fileout := os.Stdout
	lineCount := 0
	pageCount := 1

	// 判断输入方式
	if args.filename != "" { // 若给出文件名，则需检查文件是否存在
		err := errors.New("")
		filein, err = os.Open(args.filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Open file failed!\n")
			os.Exit(6)
		}
		// defer后面的函数在defer语句所在的函数执行结束的时候会被调用
		defer filein.Close() // 关闭文件
	}

	// 根据分页符进行分页处理
	readLine := bufio.NewReader(filein)
	if args.page_type == false {
		for {
			line, err := readLine.ReadString('\n')
			if err == io.EOF 
				break
			
			if err != nil {
				fmt.Fprintf(os.Stderr, "Read file error!\n")
				os.Exit(7)
			}

			lineCount++
			if lineCount > args.page_length {
				pageCount++
				lineCount = 1
			}

			// 判断输出方式
			if pageCount >= args.start_page && pageCount <= args.end_page {
				fmt.Fprintf(fileout, "%s", line)
			}
		}
	} else {
		for {
			page, err := readLine.ReadString('\f')
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Fprintf(os.Stderr, "Read file error!\n")
				os.Exit(7)
			}

			pageCount++
			if pageCount >= args.start_page && pageCount <= args.end_page {
				fmt.Fprintf(fileout, "%s", page)
			}
		}
	}

	cmd := exec.Command("cat", "-n")
	_, err := cmd.StdinPipe()
	if err != nil { // 创建管道失败
		fmt.Fprintf(os.Stderr, "Create pipe error\n")
		os.Exit(8)
	}

	// 若有dest，则直接输出到目标打印机
	if args.dest != "" {
		cmd.Stdout = fileout
		cmd.Run()
	}

	filein.Close()
	fileout.Close()
}

func main() {
	args := selpgArgs{0, 0, 72, false, "", ""}
	getArgs(&args)
	checkArgs(&args)
	handle(&args)
}
