// utils/validation.go
package utils

import (
	"errors"
	"strings"
)

var sensitiveWords = []string{
	"badword1", "badword2", // 英文亵渎词汇
	"curseword1", "curseword2", // 英文亵渎词汇
	"racistword1", "racistword2", // 英文种族歧视
	"sexistword1", "sexistword2", // 英文性别歧视
	"violentword1", "violentword2", // 英文暴力词汇
	"politicalword1", "politicalword2", // 英文政治敏感
	"pornword1", "pornword2", // 英文色情内容
	"inappropriateword1", "inappropriateword2", // 英文其他不当用语
	"脏话1", "脏话2", // 中文亵渎词汇
	"种族歧视词1", "种族歧视词2", // 中文种族歧视
	"性别歧视词1", "性别歧视词2", // 中文性别歧视
	"暴力词1", "暴力词2", // 中文暴力词汇
	"政治敏感词1", "政治敏感词2", // 中文政治敏感
	"色情词1", "色情词2", // 中文色情内容
	"不当用语1", "不当用语2", // 中文其他不当用语
}

// CheckSensitiveWords 检查文本消息中是否包含敏感词
func CheckSensitiveWords(message string) bool {
	for _, word := range sensitiveWords {
		if strings.Contains(message, word) {
			return true
		}
	}
	return false
}

// ValidateFileFormatAndSize 验证文件格式和大小
func ValidateFileFormatAndSize(fileName string, fileSize int64, allowedFormats []string, maxSize int64) error {
	// 检查文件格式
	validFormat := false
	for _, format := range allowedFormats {
		if strings.HasSuffix(fileName, format) {
			validFormat = true
			break
		}
	}
	if !validFormat {
		return errors.New("不支持的文件格式")
	}

	// 检查文件大小
	if fileSize > maxSize {
		return errors.New("文件大小超过限制")
	}

	return nil
}
