/*
 * @Description: 关于时间处理的方法
 * @Author: wubo_home
 * @Date: 2022-11-28 09:54:37
 * @LastEditors: wubo_home 13996437157@139.com
 * @LastEditTime: 2022-11-28 16:46:40
 * @FilePath: \gofirstd:\code\go\goutils\date\index.go
 * @sign: 小小的我，大大的世界我也想去走走 可不可以。。。。
 * Copyright (c) 2022 by wubo_home 13996437157@139.com, All Rights Reserved.
 */
package date

import (
	"time"
)

func getUnix() int64 {
	return time.Now().Unix()
}

func getUnixNano() int64 {
	return time.Now().UnixNano()
}

func formatTimeByUnix(timestamp int64, formatStr string) string {
	tm := time.Unix(timestamp, 0)
	return formatTime(tm, formatStr)
}

func formatTimeByStr(timestamp string, inLayout string, outLayout string) string {
	tm, _ := time.Parse(inLayout, timestamp)
	return formatTime(tm, outLayout)
}

func formatTime(tm time.Time, formatStr string) string {
	return tm.Format(formatStr)
}
