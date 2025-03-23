package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/IUnpy/aiagent/internal/translator"
)

type TranslatorUI struct {
	window       fyne.Window
	input        *widget.Entry
	output       *widget.Entry
	translateBtn *widget.Button
	fromLang     *widget.Select
	toLang       *widget.Select
}

// NewTranslatorUI 创建新的翻译器UI
func NewTranslatorUI(window fyne.Window) *TranslatorUI {
	ui := &TranslatorUI{
		window: window,
	}
	ui.createUI()
	return ui
}

// createUI 创建UI组件
func (t *TranslatorUI) createUI() {
	// 创建输入框
	t.input = widget.NewMultiLineEntry()
	t.input.SetPlaceHolder("请输入要翻译的文本...")

	// 创建输出框
	t.output = widget.NewMultiLineEntry()
	t.output.SetPlaceHolder("翻译结果将显示在这里...")
	t.output.Disable()

	// 创建语言选择下拉框
	languages := []string{"中文", "英语", "日语", "韩语", "法语", "德语"}
	t.fromLang = widget.NewSelect(languages, func(s string) {})
	t.toLang = widget.NewSelect(languages, func(s string) {})

	// 设置默认值
	t.fromLang.SetSelected("中文")
	t.toLang.SetSelected("英语")

	// 创建翻译按钮
	t.translateBtn = widget.NewButton("翻译", t.handleTranslate)

	// 创建语言选择容器
	langContainer := container.NewHBox(
		widget.NewLabel("从:"),
		t.fromLang,
		widget.NewLabel("到:"),
		t.toLang,
	)

	// 创建主布局
	content := container.NewVBox(
		langContainer,
		t.input,
		t.translateBtn,
		t.output,
	)

	// 设置窗口内容和大小
	t.window.SetContent(content)
	t.window.Resize(fyne.NewSize(500, 400))
}

// handleTranslate 处理翻译按钮点击事件
func (t *TranslatorUI) handleTranslate() {
	text := t.input.Text
	from := t.fromLang.Selected
	to := t.toLang.Selected

	// TODO: 调用翻译服务
	translated := t.translate(text, from, to)
	t.output.SetText(translated)
}

// translate 调用翻译服务
func (t *TranslatorUI) translate(text, from, to string) string {
	// TODO: 实现实际的翻译逻辑
	apiKey := "your_api_key_here"
	translator := translator.NewTranslator(apiKey)
	translated, err := translator.Translate(text, from, to)
	if err != nil {
		return "翻译失败: " + err.Error()
	}
	return translated
}
