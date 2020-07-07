/// Author:koangel
/// jackliu100@gmail.com
/// 负责分析并解析字符串格式并转换成一个time
package grapeTimer

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"

	// Day 00:00:00 每天的几点几分执行
	DayRegexp = `Day\s[0-9]{2}:[0-9]{2}:[0-9]{2}`
	// Week 1 00:00:00 每周几的几点执行
	WeekRegexp = `Week\s[0-9]{1,2}\s[0-9]{2}:[0-9]{2}:[0-9]{2}`
	// Month 1 00:00:00 每月几日的几点几分执行
	// 但是如果本月不存在该日期则不执行并报错
	MonthRegexp = `Month\s[0-9]{1,2}\s[0-9]{2}:[0-9]{2}:[0-9]{2}`
)

const (
	Day   = "Day"
	Week  = "Week"
	Month = "Month"
)

/// 错误提示
const (
	error_badFormat   = "Bad Format"
	error_monthDay    = "Date overflow"
	error_badLocation = "Bad Location"
	error_badWeekDay  = "Bad Week Day"
)

var DayCRegexp, _ = regexp.Compile(DayRegexp)
var WeekCRegexp, _ = regexp.Compile(WeekRegexp)
var MonthCRegexp, _ = regexp.Compile(MonthRegexp)

/// 有效防止超出日期如何处理
/// 获取一个月有多少天
func getMonthDay(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return
}

func AtTime(timeFmt string, loc *time.Location) (*time.Time, error) {
	nextTimeStr := fmt.Sprintf("%v %v", time.Now().Format(DateFormat), timeFmt)
	vtime, perror := time.ParseInLocation(TimeFormat, nextTimeStr, loc)
	if perror != nil {
		return nil, perror
	}
	return &vtime, nil
}

/// 基本格式 Day or Week or Month or Year
/// like : Type DayNum Time
/// format: Day 00:00:00
/// 返回一个处理好的下一个时间
func Parser(dateFmt string) (*time.Time, error) {
	// 默认使用上海时区
	loc, _ := time.LoadLocation(LocationFormat)
	return ParserLoc(dateFmt, loc)
}

//// 带有时区的分析体系
func ParserLoc(dateFmt string, loc *time.Location) (*time.Time, error) {
	dayst := strings.Split(dateFmt, " ") // 把字符串切分开
	cnowTime := time.Now()

	if loc == nil {
		return nil, errors.New(error_badLocation)
	}

	// 处理每日的日期格式
	if DayCRegexp.MatchString(dateFmt) {
		vtime, perror := AtTime(dayst[1], loc)
		if perror != nil {
			return nil, perror
		}

		if cnowTime.After(*vtime) {
			*vtime = (*vtime).AddDate(0, 0, 1) // 一天之后的时间
		}

		if CDebugMode {
			log.Printf("[grapeTimer] Day Parser:%v", *vtime)
		}

		return vtime, nil // 返回日期
	} else if WeekCRegexp.MatchString(dateFmt) {
		weekId, _ := strconv.Atoi(dayst[1])
		if weekId >= 7 || weekId < 0 {
			return nil, errors.New(error_badWeekDay) // 周的日期格式不合法
		}

		nowWeekId := int(cnowTime.Weekday())
		vtime, perror := AtTime(dayst[2], loc)
		if perror != nil {
			return nil, perror
		}

		weekDiff := weekId - nowWeekId
		if weekDiff < 0 {
			weekDiff += 7
		}

		if cnowTime.After(*vtime) || weekDiff != 0 {
			*vtime = (*vtime).AddDate(0, 0, weekDiff)
		}

		if CDebugMode {
			log.Printf("[grapeTimer] Week Parser:%v", *vtime)
		}

		return vtime, nil // 返回日期
	} else if MonthCRegexp.MatchString(dateFmt) {
		dayNum, _ := strconv.Atoi(dayst[1])
		if dayNum > getMonthDay(cnowTime.Year(), int(cnowTime.Month())) {
			return nil, errors.New(error_monthDay)
		}

		if dayNum == 0 {
			return nil, errors.New(error_monthDay)
		}

		nextTimer := fmt.Sprintf("%v-%02d-%02d %v", int(cnowTime.Year()), int(cnowTime.Month()), dayNum, dayst[2])
		vtime, perror := time.ParseInLocation(TimeFormat, nextTimer, loc)
		if perror != nil {
			return nil, perror
		}

		if cnowTime.After(vtime) {
			vtime = vtime.AddDate(0, 1, 0)
		}

		if CDebugMode {
			log.Printf("[grapeTimer] Month Parser:%v", vtime)
		}

		return &vtime, nil
	}

	return nil, errors.New(error_badFormat)
}
