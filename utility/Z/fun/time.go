package fun

import "time"

// DaysBetweenDates
// @Description 计算两个日期相差多少天
// @Author aDuo 2024-08-23 03:46:23
// @Param date1
// @Param date2
// @Return int
func DaysBetweenDates(date1, date2 time.Time) int {
	// 将两个日期都转换为Unix时间戳（秒）
	date1Unix := date1.Unix()
	date2Unix := date2.Unix()

	// 计算两个时间戳之间的差值，并将其转换为天数
	return int(date2Unix-date1Unix) / (60 * 60 * 24)
}

//
// GetToday
// @Description 获取当天的时间
// @Author aDuo 2024-09-03 18:32:30
// @Return time.Time

func GetToday() time.Time {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return today
}

func GetTime() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}
