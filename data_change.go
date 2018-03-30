package goutils

import (
	"fmt"
	"strconv"
	"strings"
)

type DataChange struct {
	//线程睡time.Sleep(1 * 1e9)
}

func (dc *DataChange) Interface2Int(i interface{}) int {

	switch v := i.(type) {
	case int:
		fmt.Println("整型", v)
		var s int
		s = v
		fmt.Println(s)
	case string:
		fmt.Println("字符串", v)
	}

	if value, ok := i.(int); ok {
		fmt.Println("类型匹配整型：%d\n", value)
		return value
	} else {
		return -1
	}
}

func (dc *DataChange) Interface2String(i interface{}) string {

	if value, ok := i.(string); ok {
		//fmt.Printf("类型匹配字符串:%s\n", value)
		return value
	} else {
		return ""
	}
}

func (dc *DataChange) Int2String(i int) string {
	s := strconv.Itoa(i)
	//s := strconv.FormatInt(int64(i), 10)
	return s
}

func (dc *DataChange) String2Int(s string) int {

	if i, err := strconv.Atoi(s); err == nil {
		return i
	} else {
		fmt.Println(err)
		return -1
	}
}

//字符串操作
func (dc *DataChange) StringOperate(s string) {

	//字符串截取
	/*s := "abcdefg"
	s = string([]byte(s)[:3])
	fmt.Println(s) //得到 "abc"*/

}

func main() {
	/*字符串基本操作--strings*/
	str := "wangdy"
	//是否包含
	fmt.Println(strings.Contains(str, "wang"), strings.Contains(str, "123")) //true false
	//获取字符串长度
	fmt.Println(len(str)) //6
	//获取字符在字符串的位置 从0开始,如果不存在，返回-1
	fmt.Println(strings.Index(str, "g")) //3
	fmt.Println(strings.Index(str, "x")) //-1
	//判断字符串是否以 xx 开头
	fmt.Println(strings.HasPrefix(str, "wa")) //true
	//判断字符串是否以 xx 结尾
	fmt.Println(strings.HasSuffix(str, "dy")) //true
	//判断2个字符串大小，相等0，左边大于右边-1，其他1
	str2 := "hahaha"
	fmt.Println(strings.Compare(str, str2)) //1
	//分割字符串
	strSplit := strings.Split("1-2-3-4-a", "-")
	fmt.Println(strSplit) //[1 2 3 4 a]
	//组装字符串
	fmt.Println(strings.Join(strSplit, "#")) //1#2#3#4#a
	//去除字符串2端空格
	fmt.Printf("%s,%s\n", strings.Trim("  我的2边有空格   1  ", " "), "/////") //我的2边有空格   1,/////
	//大小写转换
	fmt.Println(strings.ToUpper("abDCaE")) //ABDCAE
	fmt.Println(strings.ToLower("abDCaE")) //abdcae
	//字符串替换:意思是：在sourceStr中，把oldStr的前n个替换成newStr，返回一个新字符串，如果n<0则全部替换
	sourceStr := "123123123"
	oldStr := "12"
	newStr := "ab"
	n := 2
	fmt.Println(strings.Replace(sourceStr, oldStr, newStr, n))

	/*字符串转换--strconv*/
	//整型-字符串
	fmt.Println(strconv.Itoa(12)) //12
	//字符串-整型
	num, err := strconv.Atoi("-315")
	if err != nil {
		fmt.Println("occur err : ", err)
	} else {
		fmt.Println("转换后的字符串是： ", num) //-315
	}
	//bool-字符串
	fmt.Println(strconv.ParseBool("false")) //false <nil>
	//float-字符串 传递一个位数
	fmt.Println(strconv.ParseFloat("3.14", 32)) //3.140000104904175 <nil>
	fmt.Println(strconv.ParseFloat("3.14", 64)) //3.14 <nil>
	fmt.Println(strconv.ParseInt("343", 0, 64)) //343 <nil>
	//格式化操作，传递进制数
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(123, 10)) //123
	fmt.Println(strconv.FormatInt(123, 2))  //1111011
	fmt.Println(strconv.FormatInt(123, 16)) //7b
}

// 通过map主键唯一的特性过滤重复元素
func (dc *DataChange) RemoveRepByMap(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// 通过两重循环过滤重复元素
func (dc *DataChange) RemoveRepByLoop(slc []int) []int {
	result := []int{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 元素去重
func (dc *DataChange) RemoveRep(slc []int) []int {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return dc.RemoveRepByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return dc.RemoveRepByMap(slc)
	}
}
