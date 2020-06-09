package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const N int = 1000

func main() {
	allStrs := make(map[int]string, N)

	path := "D:\\编程进阶"          //读取的文件路径
	file, _ := os.Open(path)    //返回*file,只能读O_RDONLY
	defer file.Close()          //返回的时候关闭文件
	br := bufio.NewReader(file) //读取数据
	for {
		lineByte, _, err := br.ReadLine() //逐行读取
		if err == io.EOF {                //读取完毕，跳出循环
			break
		}
		lineString := string(lineByte)               //将切片转换成字符串
		lines := strings.Split(lineString, "------") //以-----分割字符串为两部分
		if len(lines) == 2 {
			qqUser, _ := strconv.Atoi(lines[0]) //将字符串转换为整数
			qqPass := lines[1]
			allStrs[qqUser] = qqPass //映射到map中
		}
	}
	startTime := time.Now()
	for {
		fmt.Println("输入要查找的数据")
		var QQ int
		fmt.Scanf("%d", &QQ) //查找QQ

		if QQPass, ok := allStrs[QQ]; ok { //查找成功输出查找到的QQ号
			fmt.Println(QQ, QQPass)
		} else {
			fmt.Println("没有找到")
		}
	}
	useTime := time.Since(startTime)
	fmt.Println("花费的时间", useTime)
}
