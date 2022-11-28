/*
 * @Description:
 * @Author: wubo_home
 * @Date: 2022-11-28 18:28:44
 * @LastEditors: wubo_home 13996437157@139.com
 * @LastEditTime: 2022-11-28 19:03:54
 * @FilePath: \gofirstd:\code\go\goutils\index.go
 * @sign: 小小的我，大大的世界我也想去走走 可不可以。。。。
 * Copyright (c) 2022 by wubo_home 13996437157@139.com, All Rights Reserved.
 */
package goutils

import (
	"fmt"
	"time"

	"github.com/w89612b/goutils/date"
)

func main() {
	fmt.Printf(date.FormatTime(time.Now(), "2006-01-02 03:04:05"))
}
