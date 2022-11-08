package main // 包，表明代码所在模块

import (
	"fmt"
	"os"
) // 引入代码依赖

// func implementation main函数不支持任何返回值
func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello world", os.Args[1])
	}

}
