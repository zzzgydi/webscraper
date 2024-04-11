package utils

import "regexp"

var (
	multiSpaceRe = regexp.MustCompile(`\n+`)
	imageRegex   = regexp.MustCompile(`!\[.*?\]\([^\)]*\)`)
	linkRegex    = regexp.MustCompile(`\[(.*?)\]\(.*?\)`)
)

func Substring(s string, start, end int) string {
	// 将字符串转换为 rune 切片，每个 rune 表示一个 UTF-8 字符
	runes := []rune(s)
	length := len(runes)

	// 校验 start 和 end 的有效性
	if start < 0 || start > length {
		start = 0
	}
	if end < 0 || end > length {
		end = length
	}
	if start > end {
		start, end = end, start
	}

	// 返回截取后的字符串
	return string(runes[start:end])
}

func SubstringMid(s string, maxSize, leftWeight, rightWeight int) string {
	// 根据最大字符，根据两边权重，去掉两边多余的字符
	if len(s) <= maxSize {
		return s
	}

	leftSize := maxSize * leftWeight / (leftWeight + rightWeight)
	return Substring(s, leftSize, leftSize+maxSize)
}

func Ellipsis(s string, max int) string {
	runes := []rune(s)
	length := len(runes)

	if length > max {
		return string(runes[:max]) + "..."
	}

	return s
}

func RemoveMultiSpace(s string) string {
	return multiSpaceRe.ReplaceAllString(s, "\n")
}

func RemoveMarkdownImages(text string) string {
	return imageRegex.ReplaceAllString(text, "")
}

func RemoveMarkdownLink(text string) string {
	return linkRegex.ReplaceAllString(text, "$1")
}
