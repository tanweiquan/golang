package hello_world

import "fmt"

func main() {
	fmt.Println("hello world")
}

// go get 包路径  就是下载第三方的依赖
// go mod tidy    就是检测并下载依赖
// 并将第三方依赖默认保存在$GOPATH/pkg/mod下
