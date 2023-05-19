package strings

import "strings"

type Case string

const (
	LowerCase          Case = "^[a-z0-9]+$"                                     // LowerCase 全小写式 e.g. "xxyyzz"
	UpperCase          Case = "^[A-Z0-9]+$"                                     // UpperCase 全大写式 e.g. "XXYYZZ"
	SnakeCase          Case = "^[a-z0-9]+(?:_[a-z0-9]+)+$"                      // SnakeCase 蛇形（小蛇式）e.g. "xx_yy_zz"
	ScreamingSnakeCase Case = "^[A-Z0-9]+(?:_[A-Z0-9]+)+$"                      // ScreamingSnakeCase 大蛇式 e.g. "XX_YY_ZZ"
	CamelCase          Case = "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)+$"             // CamelCase 驼峰式（小驼峰式）e.g. "xxYyZz"
	CamelSnakeCase     Case = "^[a-z][a-z0-9]*(?:_[A-Z][a-z0-9]*)+$"            // CamelSnakeCase 驼峰式蛇形（小驼峰式蛇形） e.g. "xx_Yy_Zz"
	PascalCase         Case = "(?=.*[a-z])^(?:[A-Z][a-z0-9]*)+$"                // PascalCase 帕斯卡式（大驼峰式）e.g. "XxYyZz"
	PascalSnakeCase    Case = "(?=.*[a-z])^[A-Z][a-z0-9]*(?:_[A-Z][a-z0-9]*)+$" // PascalSnakeCase 帕斯卡蛇形（大驼峰式蛇形） e.g. "Xx_Yy_Zz"
	KebabCase          Case = "^[a-z0-9]+(?:-[a-z0-9]+)+$"                      // KebabCase 烤串式（小烤串式） e.g. "xx-yy-zz"
	CobolCase          Case = "^[A-Z0-9]+(?:-[A-Z0-9]+)+$"                      // CobolCase 科博尔式（大烤串式） e.g. "XX-YY-ZZ"
	TrainCase          Case = "(?=.*[a-z])^[A-Z][a-z0-9]*(?:-[A-Z][a-z0-9]*)+$" // TrainCase 列车式 e.g. "Xx-Yy-Zz"

	FlatCase       = LowerCase
	UpperFlatCase  = UpperCase
	LowerCamelCase = CamelCase
	UpperCamelCase = PascalCase
	HttpHeaderCase = TrainCase
)

// CaseOf 获取字符串的命名风格
func CaseOf(s string) Case {
	// TODO
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
