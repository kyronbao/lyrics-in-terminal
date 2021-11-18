package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {

	//获取时间
	filePath := os.Args[1]
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	linesStr := strings.Split(string(data), "\n")

	var numsBefore []string
	for _, lineStr := range linesStr {
		re := regexp.MustCompile(`\d{1,2}\:\d{2}\.\d{2,3}`) //15:04.000
		numsBefore = append(numsBefore, re.FindString(lineStr))
	}

	//转换时间
	var nums []int64
	for _,numBefore := range numsBefore {
		numBefore1 := strings.Replace(numBefore, ":", "m", -1)
		numBefore2 := strings.Replace(numBefore1, ".", "s", -1)
		numBefore2 += "ms"
		//numBefore2 := 15m04s000ms
		newTime,_ := time.ParseDuration(numBefore2)
		nums = append(nums, newTime.Milliseconds())
	}


	//播放
	begin := time.Now()
	//nums := []int{1000,2000,3000}
	for key, num := range nums {

		now := begin.Add( time.Duration(num) * time.Millisecond)
		for {
			if time.Now().After(now) {
				re := regexp.MustCompile(`\[.*\]`)
				content := re.ReplaceAllString(linesStr[key], "")
				fmt.Println(content)
				break
			}
		}
	}




}
