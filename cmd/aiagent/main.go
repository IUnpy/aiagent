package main

import (
	"fmt"

	"github.com/IUnpy/aiagent/internal/config"
	"github.com/IUnpy/aiagent/internal/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		// 使用默认配置继续运行
	}

	myApp := app.New()

	if config.Theme.Dark {
		myApp.Settings().SetTheme(ui.NewCustomTheme())
	} else {
		myApp.Settings().SetTheme(theme.LightTheme())
	}
	// 使用自定义主题

	window := myApp.NewWindow("AIgent Translator")

	// 使用配置中的窗口大小
	window.Resize(fyne.NewSize(
		float32(config.Window.Width),
		float32(config.Window.Height),
	))

	// 创建UI时传入配置
	ui.NewTranslatorUI(window, config)

	window.ShowAndRun()
}
