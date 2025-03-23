package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/IUnpy/aiagent/internal/config"
	"github.com/IUnpy/aiagent/internal/translator"
)

type TranslatorUI struct {
	window        fyne.Window
	input         *widget.Entry
	fromLang      *widget.Select
	toLang        *widget.Select
	config        *config.Config
	chatContainer *fyne.Container   // 添加聊天记录容器
	chatScroll    *container.Scroll // 添加滚动容器
	toolbar       *fyne.Container   // 添加工具栏
}

// NewTranslatorUI 创建新的翻译器UI
func NewTranslatorUI(window fyne.Window, config *config.Config) *TranslatorUI {
	ui := &TranslatorUI{
		window: window,
		config: config,
	}
	ui.createUI()
	return ui
}

// createUI 创建UI组件
func (t *TranslatorUI) createUI() {
	// 创建输入框
	t.input = widget.NewMultiLineEntry()
	t.input.SetPlaceHolder("请输入要翻译的文本...")
	t.input.Resize(fyne.NewSize(400, 100))

	// 创建语言选择下拉框
	languages := []string{"中文", "英语", "日语", "韩语", "法语", "德语"}
	t.fromLang = widget.NewSelect(languages, func(s string) {})
	t.toLang = widget.NewSelect(languages, func(s string) {})

	// 设置默认值
	t.fromLang.SetSelected("中文")
	t.toLang.SetSelected("英语")

	// 创建聊天记录容器
	t.chatContainer = container.NewVBox()
	t.chatScroll = container.NewScroll(t.chatContainer)
	t.chatScroll.Resize(fyne.NewSize(480, 300))

	// 创建工具栏
	t.toolbar = container.NewHBox(
		widget.NewLabel("从:"),
		t.fromLang,
		widget.NewLabel("到:"),
		t.toLang,
		layout.NewSpacer(),
		widget.NewButton("翻译", t.handleTranslate),
	)

	// 创建主布局
	content := container.NewBorder(
		t.toolbar, // 顶部工具栏
		t.input,   // 底部输入框
		nil, nil,
		t.chatScroll, // 中间聊天记录
	)

	t.window.SetContent(content)
	t.window.Resize(fyne.NewSize(500, 600))
}

// handleTranslate 处理翻译按钮点击事件
func (t *TranslatorUI) handleTranslate() {
	text := t.input.Text
	if text == "" {
		return
	}

	// 添加用户输入到聊天记录
	wrappedInput := ChineseWrap(text, 20) // 用户输入每行20字
	inputLabel := widget.NewLabel(wrappedInput)
	inputLabel.Wrapping = fyne.TextWrapOff // 关闭自动换行，使用我们的换行

	userCard := container.NewPadded(inputLabel)
	userBox := container.NewHBox(
		layout.NewSpacer(),
		userCard,
	)
	t.chatContainer.Add(userBox)

	// 调用翻译
	translated := t.translate(text, t.fromLang.Selected, t.toLang.Selected)

	// 添加翻译结果到聊天记录
	wrappedOutput := ChineseWrap(translated, 25) // 翻译结果每行25字
	outputLabel := widget.NewLabel(wrappedOutput)
	outputLabel.Wrapping = fyne.TextWrapOff

	botCard := container.NewPadded(outputLabel)
	botBox := container.NewHBox(
		botCard,
		layout.NewSpacer(),
	)
	t.chatContainer.Add(botBox)

	// 清空输入框
	t.input.SetText("")

	// 滚动到底部
	t.chatScroll.ScrollToBottom()
}

// translate 调用翻译服务
func (t *TranslatorUI) translate(text, from, to string) string {
	if t.config.APIKey == "" {
		return "翻译失败: API密钥未设置"
	}

	translator := translator.NewTranslator(t.config.APIKey)
	translated, err := translator.Translate(text, from, to)
	if err != nil {
		return "翻译失败: " + err.Error()
	}
	return translated
}
