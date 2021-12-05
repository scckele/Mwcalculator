package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"math"
	"strconv"
)
type data struct{
	d, l, o, i, w, pi float64
}

func Weight(s *data) (float64, error) {
	s.o=math.Pow(s.o, 2)
	s.i= math.Pow(s.i, 2)
	s.pi=3.14
	s.w=s.pi*(s.o-s.i)*s.l*s.d*10e-6
	fmt.Println("weight :",s.w)
	return s.w, nil
}

func caculator(w fyne.Window) {
	//组件区
	//输入框组件
	dt:=widget.NewLabel("密度")
	d:=widget.NewEntry()
	d.SetPlaceHolder("g/cm³")


	lt:=widget.NewLabel("长度")
	l := widget.NewEntry()
	l.SetPlaceHolder("mm")


	ot:=widget.NewLabel("外径")
	o := widget.NewEntry()
	o.SetPlaceHolder("mm")


	it:=widget.NewLabel("内径")
	i := widget.NewEntry()
	i.SetPlaceHolder("mm")

	outt:=widget.NewLabel("结果:")
	//green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	out:=widget.NewEntry()

	//单位
	//vv := widget.NewLabel("g/cm³")
	//mm := widget.NewLabel("mm")
	//kg := widget.NewLabel("kg")


	//按钮组件
	button:=container.NewCenter(
		widget.NewButton("求值", func() {
			D, _ :=strconv.ParseFloat(d.Text,64)
			L, _ :=strconv.ParseFloat(l.Text,64)
			I, _ :=strconv.ParseFloat(i.Text,64)
			O, _ :=strconv.ParseFloat(o.Text,64)
			J := &data{D, L, O/2, I/2, 0, 3.14}
			f, _ := Weight(J)
			str := binding.NewString()
			_ = str.Set(strconv.FormatFloat(f, 'f', 5, 32))
			out.Bind(str)
		}))

	//容器布局区
	form1:=container.New(layout.NewFormLayout(),dt,d,ot,o)
	form2:=container.New(layout.NewFormLayout(),lt,l,it,i)
	inputdata:=container.New(layout.NewGridLayout(2),form1,form2)

	form3:=container.New(layout.NewFormLayout(),outt,out)
	content := container.NewVBox(inputdata,button,form3)

	//写入窗口
	w.SetContent(content)
}
