package main
import (
  "fmt"
  "os"
)
func main(){
  dir, _ := os.Getwd()
  fmt.Println("当前的目录是:", dir)  //当前的目录是: D:\test 我的环境是windows 如果linix 就是/xxx/xxx
}
