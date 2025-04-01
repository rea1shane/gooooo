package time

import (
	"fmt"
	"strings"
	"time"
)

// FormatDuration 格式化 time.Duration，保留正负号，精度到秒
// 样例：-2d 15h 4m 5s
func FormatDuration(duration time.Duration) string {
	// 处理负数情况
	sign := ""
	if duration < 0 {
		sign = "-"
		duration = -duration
	}

	// 提取天数
	days := duration / (24 * time.Hour)
	duration %= 24 * time.Hour

	// 提取小时
	hours := duration / time.Hour
	duration %= time.Hour

	// 提取分钟
	minutes := duration / time.Minute
	duration %= time.Minute

	// 提取秒数 (忽略小数部分)
	seconds := duration / time.Second

	// 构建结果字符串
	var parts []string

	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}

	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}

	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}

	// 始终添加秒数，即使为0
	parts = append(parts, fmt.Sprintf("%ds", seconds))

	// 使用空格连接并添加符号
	return sign + strings.Join(parts, " ")
}
