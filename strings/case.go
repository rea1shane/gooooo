package strings

import (
	"regexp"
	"strings"
)

type Case string

func (c Case) String() string {
	switch c {
	case LowerCase:
		return "lowercase | flatcase"
	case UpperCase:
		return "UPPERCASE | UPPERFLATCASE"
	case SnakeCase:
		return "snake_case | snail_case | pothole_case"
	case ScreamingSnakeCase:
		return "SCREAMING_SNAKE_CASE | MACRO_CASE | CONSTANT_CASE"
	case CamelCase:
		return "camelCase | lowerCamelCase | dromedaryCase"
	case CamelSnakeCase:
		return "camel_Snake_Case"
	case PascalCase:
		return "PascalCase | UpperCamelCase | StudlyCase"
	case PascalSnakeCase:
		return "Pascal_Snake_Case | Title_Case"
	case KebabCase:
		return "kebab-case | dash-case | lisp-case | spinal-case"
	case CobolCase:
		return "COBOL-CASE | SCREAMING-KEBAB-CASE"
	case TrainCase:
		return "Train-Case | HTTP-Header-Case"
	}
	return "unknown case"
}

const (
	LowerCase          Case = "^[a-z0-9]+$"                                // LowerCase 全小写式 e.g. "twowords"
	UpperCase          Case = "^[A-Z0-9]+$"                                // UpperCase 全大写式 e.g. "TWOWORDS"
	SnakeCase          Case = "^[a-z0-9]+(?:_[a-z0-9]+)+$"                 // SnakeCase 蛇形（小蛇式） e.g. "two_words"
	ScreamingSnakeCase Case = "^[A-Z0-9]+(?:_[A-Z0-9]+)+$"                 // ScreamingSnakeCase 大蛇式 e.g. "TWO_WORDS"
	CamelCase          Case = "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)+$"        // CamelCase 驼峰式（小驼峰式） e.g. "twoWords"
	CamelSnakeCase     Case = "^[a-z][A-Za-z0-9]*(?:_[A-Z][A-Za-z0-9]*)+$" // CamelSnakeCase 驼峰式蛇形（小驼峰式蛇形） e.g. "two_Words"
	PascalCase         Case = "^(?:[A-Z][a-z0-9]*)+$"                      // PascalCase 帕斯卡式（大驼峰式） e.g. "TwoWords"
	PascalSnakeCase    Case = "^[A-Z][A-Za-z0-9]*(?:_[A-Z][A-Za-z0-9]*)+$" // PascalSnakeCase 帕斯卡蛇形（大驼峰式蛇形） e.g. "Two_Words"
	KebabCase          Case = "^[a-z0-9]+(?:-[a-z0-9]+)+$"                 // KebabCase 烤串式（小烤串式） e.g. "two-words"
	CobolCase          Case = "^[A-Z0-9]+(?:-[A-Z0-9]+)+$"                 // CobolCase 科博尔式（大烤串式） e.g. "TWO-WORDS"
	TrainCase          Case = "^[A-Z][A-Za-z0-9]*(?:-[A-Z][A-Za-z0-9]*)+$" // TrainCase 列车式 e.g. "Two-Words"

	FlatCase = LowerCase

	UpperFlatCase = UpperCase

	SnailCase   = SnakeCase
	PotholeCase = SnakeCase

	MacroCase    = ScreamingSnakeCase
	ConstantCase = ScreamingSnakeCase

	LowerCamelCase = CamelCase
	DromedaryCase  = CamelCase

	UpperCamelCase = PascalCase
	StudlyCase     = PascalCase

	TitleCase = PascalSnakeCase

	DashCase   = KebabCase
	LispCase   = KebabCase
	SpinalCase = KebabCase

	ScreamingKebabCase = CobolCase

	HttpHeaderCase = TrainCase

	Unknown Case = "^.*$" // Unknown 未知
)

var allCases = []Case{
	LowerCase,
	UpperCase,
	SnakeCase,
	ScreamingSnakeCase,
	CamelCase,
	CamelSnakeCase,
	PascalCase,
	PascalSnakeCase,
	KebabCase,
	CobolCase,
	TrainCase,
}

// CaseOf 获取字符串的命名风格。
func CaseOf(s string) Case {
	for _, c := range allCases {
		// 检查是否包含小写字母
		// 因为 Go 的正则 RE2 不支持 PERL 的 look-ahead assertion "?=" 语法，所以才出此下策检测是否存在小写字母，否则可以直接在正则表达式前面加 "(?=.*[a-z])" 来判断
		// PERL assertion 删除于 commit 8a4620430f018e1b626c8ab8c755c5cac2b23b01
		if c == PascalCase || c == PascalSnakeCase || c == TrainCase {
			if contained, _ := regexp.Match("[a-z]", []byte(s)); !contained {
				continue
			}
		}

		if matched, _ := regexp.Match(string(c), []byte(s)); matched {
			return c
		}
	}
	return Unknown
}

// Break 将字符串分解成单词列表。
// 数字不会单独分解出来。
// 如果命名风格为 LowerCase UpperCase Unknown 其中一个，不会进行分解。
func Break(s string) (words []string) {
	switch CaseOf(s) {
	case LowerCase, UpperCase, Unknown:
		words = append(words, s)
	case CamelCase, PascalCase:
		cursor := 0
		for i := 1; i < len(s); i++ {
			letter := s[i]
			if letter >= 'A' && letter <= 'Z' {
				words = append(words, s[cursor:i])
				cursor = i
			}
		}
		words = append(words, s[cursor:])
	case SnakeCase, ScreamingSnakeCase, CamelSnakeCase, PascalSnakeCase:
		words = strings.Split(s, "_")
	case KebabCase, CobolCase, TrainCase:
		words = strings.Split(s, "-")
	}
	return words
}

// ConvertCase 转换命名风格。
func ConvertCase(s string, c Case) (string, error) {
	return "", nil
}
