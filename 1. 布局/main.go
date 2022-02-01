package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/**
 * 水平布局:HBox
 */
func demo1() {
	a := app.New()
	w := a.NewWindow("Hello")

	text1 := widget.NewLabel("text1")
	text2 := widget.NewLabel("text2")
	text3 := widget.NewLabel("text3")
	// 水平布局:HBox
	w.SetContent(container.NewHBox(
		text1,
		text2,
		text3,
	))
	w.ShowAndRun()
}

/**
 * 垂直布局:VBox
 */
func demo2() {
	a := app.New()
	w := a.NewWindow("Hello")

	text1 := widget.NewLabel("text1")
	text2 := widget.NewLabel("text2")
	text3 := widget.NewLabel("text3")
	// 垂直布局:VBox
	w.SetContent(container.NewVBox(
		text1,
		text2,
		text3,
	))
	w.ShowAndRun()
}

/**
 * 中心布局:Center
 */
func demo3() {
	a := app.New()
	w := a.NewWindow("Hello")

	text1 := widget.NewLabel("text1")
	text2 := widget.NewLabel("text2")
	text3 := widget.NewLabel("text3")
	// 中心布局:Center
	w.SetContent(container.NewCenter(
		text1,
		text2,
		text3,
	))
	w.ShowAndRun()
}

/**
 * 表单布局:Form
 */
func demo4() {
	a := app.New()
	w := a.NewWindow("Hello")

	text1 := widget.NewLabel("text1")
	text2 := widget.NewLabel("text2")
	text3 := widget.NewLabel("text3")
	// 表单布局:Form
	w.SetContent(widget.NewForm(
		widget.NewFormItem("属性1", text1),
		widget.NewFormItem("属性2", text2),
		widget.NewFormItem("属性3", text3),
	))
	w.ShowAndRun()
}

/**
 * 网格布局:Grid
 */
func demo5() {
	a := app.New()
	w := a.NewWindow("Hello")

	text1 := widget.NewLabel("text1")
	text2 := widget.NewLabel("text2")
	text3 := widget.NewLabel("text3")
	text4 := widget.NewLabel("text4")
	text5 := widget.NewLabel("text5")
	// 网格布局:Grid 2列
	w.SetContent(container.NewGridWithColumns(2, text1, text2, text3, text4, text5))
	w.ShowAndRun()
}

/**
 * 边界布局:border
 */
func demo6() {
	a := app.New()
	w := a.NewWindow("Hello")

	text1 := widget.NewLabel("text1")
	text2 := widget.NewLabel("text2")
	text3 := widget.NewLabel("text3")
	text4 := widget.NewLabel("text4")
	// 网格布局:border 上、左、下、右
	w.SetContent(container.NewBorder(text1, text2, text3, text4))
	w.ShowAndRun()
}

/**
 * Tab布局
 */
func demo7() {
	a := app.New()
	w := a.NewWindow("Hello")

	text1 := widget.NewLabel("text1")
	text2 := widget.NewLabel("text2")
	// Tab布局
	w.SetContent(container.NewAppTabs(
		container.NewTabItem("table1", text1),
		container.NewTabItem("table2", text2),
	))
	w.ShowAndRun()
}

/**
 * 滚动布局
 */
func demo8() {
	a := app.New()
	w := a.NewWindow("Hello")

	text1 := widget.NewLabel("text1")
	text2 := widget.NewLabel("text2")
	text3 := widget.NewLabel("text3")
	text4 := widget.NewLabel("text4")
	text5 := widget.NewLabel("text5")
	text6 := widget.NewLabel("text6")
	text7 := widget.NewLabel("text7")
	text8 := widget.NewLabel("text8")
	text9 := widget.NewLabel("text9")
	// 垂直滚动
	scroll := container.NewVScroll(container.NewVBox(text1, text2, text3, text4, text5, text6, text7, text8, text9))
	w.SetContent(scroll)

	w.ShowAndRun()
}

/**
 * 练习：组合布局
 */
func demo9() {
	a := app.New()
	w := a.NewWindow("Hello")

	text1 := widget.NewLabel("text1")
	text2 := widget.NewLabel("text2")
	text3 := widget.NewLabel("text2")

	box := container.NewVBox(
		text1,
		container.NewHBox(text2, text3))

	// Tab布局
	w.SetContent(box)
	w.ShowAndRun()
}

func main() {
	demo9()
}
