/*
 * @Description: 关于时间处理的方法
 * @Author: wubo_home
 * @Date: 2022-11-28 09:54:37
 * @LastEditors: wubo_home 13996437157@139.com
 * @LastEditTime: 2022-11-28 18:55:18
 * @FilePath: \gofirstd:\code\go\goutils\date\index.go
 * @sign: 小小的我，大大的世界我也想去走走 可不可以。。。。
 * Copyright (c) 2022 by wubo_home 13996437157@139.com, All Rights Reserved.
 */
package date

import (
	"time"
)

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

func FormatTimeByUnix(timestamp int64, formatStr string) string {
	tm := time.Unix(timestamp, 0)
	return FormatTime(tm, formatStr)
}

func FormatTimeByStr(timestamp string, inLayout string, outLayout string) string {
	tm, _ := time.Parse(inLayout, timestamp)
	return FormatTime(tm, outLayout)
}

func FormatTime(tm time.Time, formatStr string) string {
	return tm.Format(formatStr)
}
