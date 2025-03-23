package main

import (
	"github.com/IUnpy/aiagent/internal/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(ui.NewCustomTheme()) // 重新启用主题
	window := myApp.NewWindow("AIgent Translator")

	// 创建UI
	ui.NewTranslatorUI(window)

	// 运行应用
	window.ShowAndRun()
}
