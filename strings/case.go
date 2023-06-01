package strings

import (
	"errors"
	"regexp"
	"strings"
)

type Case int

// 在 CamelCase 和 PascalCase 中，如果遇到缩写则视为一个单词，首字母视情况大小写，其余字母全小写。见：https://en.wikipedia.org/wiki/Camel_case#Programming_and_coding
const (
	Unknown            Case = iota // Unknown 未知
	LowerCase                      // LowerCase 全小写式 e.g. "twowords"
	UpperCase                      // UpperCase 全大写式 e.g. "TWOWORDS"
	SnakeCase                      // SnakeCase 蛇形（小蛇式） e.g. "two_words"
	ScreamingSnakeCase             // ScreamingSnakeCase 大蛇式 e.g. "TWO_WORDS"
	CamelCase                      // CamelCase 驼峰式（小驼峰式） e.g. "twoWords"
	CamelSnakeCase                 // CamelSnakeCase 驼峰式蛇形（小驼峰式蛇形） e.g. "two_Words"
	PascalCase                     // PascalCase 帕斯卡式（大驼峰式） e.g. "TwoWords"
	PascalSnakeCase                // PascalSnakeCase 帕斯卡蛇形（大驼峰式蛇形） e.g. "Two_Words"
	KebabCase                      // KebabCase 烤串式（小烤串式） e.g. "two-words"
	CobolCase                      // CobolCase 科博尔式（大烤串式） e.g. "TWO-WORDS"
	TrainCase                      // TrainCase 列车式 e.g. "Two-Words"

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
)

// pattern 返回对应的正则匹配表达式
func (c Case) pattern() string {
	switch c {
	case LowerCase:
		return "^[a-z0-9]+$"
	case UpperCase:
		return "^[A-Z0-9]+$"
	case SnakeCase:
		return "^[a-z0-9]+(?:_[a-z0-9]+)+$"
	case ScreamingSnakeCase:
		return "^[A-Z0-9]+(?:_[A-Z0-9]+)+$"
	case CamelCase:
		return "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)+$"
	case CamelSnakeCase:
		return "^[a-z][A-Za-z0-9]*(?:_[A-Z][A-Za-z0-9]*)+$"
	case PascalCase:
		return "^(?:[A-Z][a-z0-9]*)+$"
	case PascalSnakeCase:
		return "^[A-Z][A-Za-z0-9]*(?:_[A-Z][A-Za-z0-9]*)+$"
	case KebabCase:
		return "^[a-z0-9]+(?:-[a-z0-9]+)+$"
	case CobolCase:
		return "^[A-Z0-9]+(?:-[A-Z0-9]+)+$"
	case TrainCase:
		return "^[A-Z][A-Za-z0-9]*(?:-[A-Z][A-Za-z0-9]*)+$"
	}
	return "^.*$"
}

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

		if matched, _ := regexp.Match(c.pattern(), []byte(s)); matched {
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
	return
}

// ConvertCase 转换命名风格。
func ConvertCase(s string, c Case) (result string, err error) {
	words := Break(s)
	switch c {
	case LowerCase:
		result = strings.ToLower(strings.Join(words, ""))
	case UpperCase:
		result = strings.ToUpper(strings.Join(words, ""))
	case SnakeCase:
		result = strings.ToLower(strings.Join(words, "_"))
	case ScreamingSnakeCase:
		result = strings.ToUpper(strings.Join(words, "_"))
	case CamelCase:
		camelWords(words, true, true)
		result = strings.Join(words, "")
	case CamelSnakeCase:
		camelWords(words, true, false)
		result = strings.Join(words, "_")
	case PascalCase:
		camelWords(words, false, true)
		result = strings.Join(words, "")
	case PascalSnakeCase:
		camelWords(words, false, false)
		result = strings.Join(words, "_")
	case KebabCase:
		result = strings.ToLower(strings.Join(words, "-"))
	case CobolCase:
		result = strings.ToUpper(strings.Join(words, "-"))
	case TrainCase:
		camelWords(words, false, false)
		result = strings.Join(words, "-")
	default:
		err = errors.New(Unknown.String())
	}
	return
}

// camelWords 驼峰化数组。
func camelWords(words []string, lowerFirstWord bool, treatAbbreviationsAsWords bool) {
	for i := 0; i < len(words); i++ {
		if len(words[i]) == 0 {
			continue
		}
		if i == 0 && lowerFirstWord {
			words[i] = strings.ToLower(words[i])
			continue
		}
		otherLetters := words[i][1:]
		if treatAbbreviationsAsWords {
			otherLetters = strings.ToLower(otherLetters)
		}
		words[i] = strings.ToUpper(words[i][0:1]) + otherLetters
	}
}
