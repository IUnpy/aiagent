package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CustomTheme struct{}

func NewCustomTheme() fyne.Theme {
	return &CustomTheme{}
}

func (t *CustomTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	if n == theme.ColorNameForeground {
		return color.RGBA{R: 0, G: 0, B: 0, A: 255} // 纯黑色文本
	}
	if n == theme.ColorNameBackground {
		return color.RGBA{R: 240, G: 240, B: 240, A: 255} // 浅灰色背景
	}
	if n == theme.ColorNameDisabled {
		return color.RGBA{R: 0, G: 0, B: 0, A: 255} // 禁用状态也使用黑色
	}
	return theme.DefaultTheme().Color(n, v)
}

func (t *CustomTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

func (t *CustomTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (t *CustomTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

func (t *CustomTheme) TextColor() color.Color {
	return color.RGBA{R: 0, G: 0, B: 0, A: 255} // 纯黑色，A值控制透明度(0-255)
}
