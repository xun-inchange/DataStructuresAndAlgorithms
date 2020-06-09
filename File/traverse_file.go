package File

import (
	"errors"
	"fmt"
	"io/ioutil"
)

/*
递归遍历文件
*/

func getAllFiles(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) //读取文件
	if err != nil {
		return files, errors.New("文件读取失败！,errMsg:" + err.Error())
	}
	//遍历文件
	for _, fi := range read {
		if fi.IsDir() { //是否为文件夹
			fullDir := path + "\\" + fi.Name() //构造新的路径
			files = append(files, fullDir)
			files, _ = getAllFiles(fullDir, files)
		}
		//如果file是文件
		fullDir := path + "\\" + fi.Name()
		files = append(files, fullDir)
	}
	return files, nil
}

func traverseFile() {
	path := "D:\\编程进阶\\Golang学习笔记\\Go web"
	var files []string                  //数组字符串
	files, _ = getAllFiles(path, files) //抓取所有文件
	for _, file := range files {
		fmt.Println(file)
	}
}
