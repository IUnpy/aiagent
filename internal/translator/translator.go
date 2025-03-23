package translator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/IUnpy/aiagent/internal/api"
)

// Translator 翻译器接口
type Translator interface {
	Translate(text, from, to string) (string, error)
}

// SimpleTranslator 简单翻译器实现
type SimpleTranslator struct {
	client *api.Client
}

// NewTranslator 创建新的翻译器实例
func NewTranslator(apiKey string) *SimpleTranslator {
	return &SimpleTranslator{
		client: api.NewClient(apiKey),
	}
}

// APIResponse 定义API响应结构
type APIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// Translate 执行翻译
func (t *SimpleTranslator) Translate(text, from, to string) (string, error) {
	if strings.TrimSpace(text) == "" {
		return "", errors.New("翻译文本不能为空")
	}

	prompt := fmt.Sprintf("Translate the following text from %s to %s:\n%s", from, to, text)
	return t.client.Chat(prompt)
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
