package strings

import (
	"regexp"
	"strings"
)

type Case string

func (c Case) String() string {
	switch c {
	case LowerCase:
		return "LowerCase | FlatCase"
	case UpperCase:
		return "UpperCase | UpperFlatCase"
	case SnakeCase:
		return "SnakeCase | SnailCase | PotholeCase"
	case ScreamingSnakeCase:
		return "ScreamingSnakeCase | MacroCase | ConstantCase"
	case CamelCase:
		return "CamelCase | LowerCamelCase | DromedaryCase"
	case CamelSnakeCase:
		return "CamelSnakeCase"
	case PascalCase:
		return "PascalCase | UpperCamelCase | StudlyCase"
	case PascalSnakeCase:
		return "PascalSnakeCase | TitleCase"
	case KebabCase:
		return "KebabCase | DashCase | LispCase | SpinalCase"
	case CobolCase:
		return "CobolCase | ScreamingKebabCase"
	case TrainCase:
		return "TrainCase | HttpHeaderCase"
	}
	return "unknown"
}

const (
	LowerCase          Case = "^[a-z0-9]+$"                                // LowerCase 全小写式 e.g. "twowords"
	UpperCase          Case = "^[A-Z0-9]+$"                                // UpperCase 全大写式 e.g. "TWOWORDS"
	SnakeCase          Case = "^[a-z0-9]+(?:_[a-z0-9]+)+$"                 // SnakeCase 蛇形（小蛇式） e.g. "two_words"
	ScreamingSnakeCase Case = "^[A-Z0-9]+(?:_[A-Z0-9]+)+$"                 // ScreamingSnakeCase 大蛇式 e.g. "TWO_WORDS"
	CamelCase          Case = "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)+$"        // CamelCase 驼峰式（小驼峰式） e.g. "twoWords"
	CamelSnakeCase     Case = "^[a-z][a-z0-9]*(?:_[A-Z][a-z0-9]*)+$"       // CamelSnakeCase 驼峰式蛇形（小驼峰式蛇形） e.g. "two_Words"
	PascalCase         Case = "^(?:[A-Z][a-z0-9]*)+$"                      // PascalCase 帕斯卡式（大驼峰式） e.g. "TwoWords"
	PascalSnakeCase    Case = "^[A-Z][a-z0-9]*(?:_[A-Z][a-z0-9]*)+$"       // PascalSnakeCase 帕斯卡蛇形（大驼峰式蛇形） e.g. "Two_Words"
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

// CaseOf 获取字符串的命名风格
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

// Pascal2Snake 帕斯卡（大驼峰式）转蛇形
// 样例：XxYy => xx_yy | XxYY => xx_y_y
func Pascal2Snake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// Snake2Pascal 蛇形转帕斯卡（大驼峰式）
// 样例：xx_yy => XxYx | xx_y_y => XxYY
func Snake2Pascal(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
