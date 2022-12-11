package diyTools

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// 数组，和连接他们的符号
func arrToString(arr []interface{}, s string) string {
	var rs string
	num := len(arr) - 1
	for i, item := range arr {
		if i != num {
			rs += item.(string) + s
		} else {
			rs += item.(string)
		}
	}
	return rs

}

func CheckArr(user, exam string) bool {
	userArr := strings.Split(user, "@.@")
	sort.Slice(userArr, func(i, j int) bool {
		return userArr[i] > userArr[j]
	})
	user = strings.Join(userArr, "@.@")

	examArr := strings.Split(exam, "@.@")
	sort.Slice(examArr, func(i, j int) bool {
		return examArr[i] > examArr[j]
	})
	exam = strings.Join(examArr, "@.@")

	return exam == user

}

// 存放计算时间的一些工具
// 时间差
func SubTime(startTime, endTime string) string {

	//待转化为时间戳的字符串
	timeLayout := "2006-01-02 15:04:05" //转化所需模板，go默认时间
	//go默认时间 很好记  6 1 2 3 4 5
	sTime, _ := time.ParseInLocation(timeLayout, startTime, time.Local) //使用模板在对应时区转化为time.time类型
	fmt.Printf("开始时间:%s\n", sTime)
	eTime, _ := time.ParseInLocation(timeLayout, endTime, time.Local)
	fmt.Printf("结束时间:%s\n", eTime)
	//计算时间差
	left := eTime.Sub(sTime) //计算时间差
	fmt.Println(left)        //1h25m41.938256s 现在看打印的值是符号我们的要求
	return left.String()
}

// 计算现在的时间过某个时间后
func GetDuring(during int64) (s, e string) {

	nowTime := time.Now()
	startTime := nowTime.Format("2006-01-02 15:04:05")
	x := nowTime.Unix() + int64(60*during)
	endTime := time.Unix(x, 0).Format("2006-01-02 15:04:05")
	return startTime, endTime
}

func GetWeek() (week string) {

	nowTime := time.Now()
	// nowDay := nowTime.Format("2006-01-02 15:04:05")
	str := []string{}
	for i := 1; i < 8; i++ {
		t := nowTime.Unix() - int64(60*24*60*i)

		timeFm := time.Unix(t, 0).Format("2006-1-2")
		str = append(str, "'"+timeFm+"'")
	}
	week = strings.Join(str, ",")
	return week
}

//生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

func GetPage(count, setPage, len int64) (start, end, getPage, lenPage int64) {
	if (setPage-1)*count > len {
		setPage = 1
	}

	start = (setPage - 1) * count

	if len > count {

		if len < setPage*count {
			fmt.Printf("len:%d,need:%d\n", len, (setPage * count))
			end = len
		} else {
			end = setPage * count
		}
	} else {

		end = len
	}
	getPage = setPage
	num := float64(len) / float64(count)
	lenPage = int64(math.Ceil(num))
	// 开始节点与结束，当前第几页，总页数
	return start, end, getPage, lenPage

}
