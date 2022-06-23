package common

import "time"

// NowTime 时间格式化
func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// WeekTime 获取本周当天之前的日期
// 比如1，days = 0 指获取当前的日期
// 比如2，days = 1 指获取昨天的日期
// 比如3，days = 2 指获取前天的日期
func WeekTime(days int) (dayTime string) {
	switch days {
	case 0:
		dayTime = BeforeTime(0 * time.Hour)
	case 1:
		dayTime = BeforeTime(-24 * time.Hour)
	case 2:
		dayTime = BeforeTime(-48 * time.Hour)
	case 3:
		dayTime = BeforeTime(-72 * time.Hour)
	case 4:
		dayTime = BeforeTime(-96 * time.Hour)
	case 5:
		dayTime = BeforeTime(-120 * time.Hour)
	case 6:
		dayTime = BeforeTime(-144 * time.Hour)
	case 7:
		dayTime = BeforeTime(-168 * time.Hour)
	default:
	}
	return dayTime
}

// BeforeTime 获取之前的日期
func BeforeTime(d time.Duration) string {
	return time.Now().Add(d).Format("2006-01-02")
}
