package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

func main(){
	//密度 长度 外径 内径 千克 吨
	//Density length outer_diameter inner_diameter kg ton


	//新建app项目
	calculator:=app.New()

	//设置窗口名称 大小
	window :=calculator.NewWindow("Material weight calculator")
	window.Resize(fyne.NewSize(400, 180))

	caculator(window)

	//运行
	window.ShowAndRun()
}

//初始化字体
func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		//fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simhei.ttf") {
			_ = os.Setenv("FYNE_FONT", path)
			break
		}
	}
	fmt.Println("=============")
}