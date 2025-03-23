package translator

import (
	"errors"
	"strings"
)

// Translator 翻译器接口
type Translator interface {
	Translate(text, from, to string) (string, error)
}

// SimpleTranslator 简单翻译器实现
type SimpleTranslator struct {
	apiKey string
}

// NewTranslator 创建新的翻译器实例
func NewTranslator(apiKey string) *SimpleTranslator {
	return &SimpleTranslator{
		apiKey: apiKey,
	}
}

// Translate 执行翻译
func (t *SimpleTranslator) Translate(text, from, to string) (string, error) {
	if strings.TrimSpace(text) == "" {
		return "", errors.New("翻译文本不能为空")
	}

	// TODO: 这里需要实现实际的翻译API调用
	// 例如：调用Google Translate API 或其他翻译服务
	// 现在返回模拟数据
	return "This is a translated text", nil
}

// 语言代码映射
var LanguageMap = map[string]string{
	"中文": "zh",
	"英语": "en",
	"日语": "ja",
	"韩语": "ko",
	"法语": "fr",
	"德语": "de",
}

// GetLanguageCode 获取语言代码
func GetLanguageCode(language string) string {
	if code, ok := LanguageMap[language]; ok {
		return code
	}
	return "en" // 默认返回英语
}
