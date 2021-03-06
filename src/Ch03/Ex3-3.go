// 연습문제 3-3
// 두가지의 날짜 및 시간 포멧을 처리하도록 timeDate.go 파일 수정해보기

package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	logs := []string{"127.0.0.1 - - [16/Nov/2020:10:49:46 +0200] 325504",
		"127.0.0.1 - - [16/Nov/2020:10:49:46 +0200] \"GET /CVEN HTTP/1.1\" 20012531 \"-\" \"Mozilla/5.0 AppleWebkit/537.36",
		"127.0.0.1 200 9412 - - [12/Nov/2020:06:26:05 +0200] \"GET \"http://www.mtsoukalos.eu/taxonomy/term/47\" 1507",
		"[12/Nov/2020:16:27:21 +0300]",
		"[12/Nov/2020:20:88:21 +0200]",
		"[12/Nov/2020:20:21 +0200]",
	}
	for _, logEntry := range logs {
		r := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`)

		if r.MatchString(logEntry) {
			match := r.FindStringSubmatch(logEntry)

			dt, err := time.Parse("15/Jan/2019:22:18:05 -0700", match[1])
			if err == nil {
				newFormat := dt.Format(time.RFC850)
				fmt.Println(newFormat)
			} else {
				fmt.Println("time format 위반")
			}
		} else {
			fmt.Println("매칭안됨~!")
		}
	}
}
