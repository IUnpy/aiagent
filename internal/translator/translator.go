package translator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Translator 翻译器接口
type Translator interface {
	Translate(text, from, to string) (string, error)
}

// SimpleTranslator 简单翻译器实现
type SimpleTranslator struct {
	apiKey string
	apiURL string
}

// NewTranslator 创建新的翻译器实例
func NewTranslator(apiKey string) *SimpleTranslator {
	return &SimpleTranslator{
		apiKey: apiKey,
		apiURL: "https://api.siliconflow.cn/v1/chat/completions",
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

	// 构建请求体
	reqBody := fmt.Sprintf(`{
		"model": "Pro/deepseek-ai/DeepSeek-V3",
		"messages": [
			{
				"role": "user",
				"content": "Translate the following text from %s to %s:\n%s"
			}
		],
		"stream": false,
		"max_tokens": 512,
		"temperature": 0.7
	}`, from, to, text)

	// 创建请求
	req, err := http.NewRequest("POST", t.apiURL, strings.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Add("Authorization", "Bearer "+t.apiKey)
	req.Header.Add("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	// 解析响应JSON
	var response APIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	if len(response.Choices) == 0 {
		return "", errors.New("翻译结果为空")
	}

	return response.Choices[0].Message.Content, nil
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
