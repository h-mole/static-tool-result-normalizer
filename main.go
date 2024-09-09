package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/h-mole/static-tool-result-normalizer/tools_result_processor"
)

func main() {
	toolType := flag.String("tooltype", "", "工具类型，取值为'cppcheck', 'tscancode', 'flawfinder'")
	inputFile := flag.String("input_file", "", "输入的文件路径")
	outputFile := flag.String("output_file", "", "输出的文件路径")
	projectDir := flag.String("project_dir", "", "代码扫描的根目录")

	// 解析命令行参数
	flag.Parse()

	// 检查必要的参数是否已经设置
	if *toolType == "" {
		fmt.Println("错误：必须指定工具类型 (--tooltype)")
		os.Exit(1)
	}
	if *inputFile == "" {
		fmt.Println("错误：必须指定输入文件 (--input_file)")
		os.Exit(1)
	}
	if *outputFile == "" {
		fmt.Println("错误：必须指定输出文件 (--output_file)")
		os.Exit(1)
	}
	if *projectDir == "" {
		fmt.Println("错误：必须指定项目目录 (--project_dir)")
		os.Exit(1)
	}
	switch *toolType {
	case "cppcheck":
		err := tools_result_processor.ParseCppCheckResult(*projectDir, *inputFile, *outputFile)
		if err != nil {
			fmt.Println("解析cppcheck结果失败:", err)
			os.Exit(1)
		}
	case "tscancode":
		err := tools_result_processor.ParseTscancodeResult(*projectDir, *inputFile, *outputFile)
		if err != nil {
			fmt.Println("解析TScanCode结果失败:", err)
			os.Exit(1)
		}
	case "flawfinder":
		err := tools_result_processor.ParseFlawFinderResult(*projectDir, *inputFile, *outputFile)
		if err != nil {
			fmt.Println("解析Flawfinder结果失败:", err)
			os.Exit(1)
		}
	}
	// 打印参数值，实际应用中可以替换为具体的逻辑处理
	fmt.Printf("工具类型: %s\n", *toolType)
	fmt.Printf("输入文件: %s\n", *inputFile)
	fmt.Printf("输出文件: %s\n", *outputFile)

}
