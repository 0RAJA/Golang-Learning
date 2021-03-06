package main

import (
	"bufio"
	"bytes"
	"container/list"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func main() {
	//preorder := []int{3, 9, 20, 15, 7}
	//inorder := []int{9, 3, 15, 20, 7}
	//buildTree(preorder, inorder)
	//fmt.Println(isNumber("0.8"))
	//str := "1^1^1^"
	//fmt.Println(len(strings.Split(str, "^")))
	//fmt.Println(fib(3))
	//fmt.Println(translateNum(123454321))
	//fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	//testAppend1()
	//testAppend2()
	//testJson()
	//testInput()
	//testSearch()
	//testW()
	//TestSearch()
	//fmt.Println(TestlargestNumber([]int{0, 0}))
	//TestByteToString()
	//constructArr([]int{1, 2, 3, 4, 5})
	//TestSplit()
	//testList()
	// TestDiff()
	//TestRace()
	//TestEqual()
	//TestTcp()
	//TestRand()
	//TestSelect()
	//TestJson()
	//TestSize()
	//TestBinary()
	//TestBuffer()
	//TestSub()
	testMain()
}

func testMain() {
	fmt.Println("hello world")
}

func TestSub() {
	t1, _ := time.Parse("15:04", "23:59")
	t2, _ := time.Parse("15:04", "00:00")
	fmt.Println(t2.Sub(t1).Minutes())
}

func TestBuffer() {
	buff := bytes.NewBuffer([]byte{})
	buff.Write([]byte("123"))
	fmt.Println(buff.Bytes())
}

func TestBinary() {
	nums := make([]byte, 10)
	fmt.Println(binary.PutVarint(nums, 1))
	fmt.Println(binary.Varint(nums))
	//binary.Read()
}

func TestSize() {
	//num := int64(1)
	//num2 := int32(1)
	nums := []int{21, 3, 4}
	fmt.Println(unsafe.Sizeof(nums))
}

func TestJson() {
	var path = "test"
	flag.StringVar(&path, "path", "", "路径")
	flag.Parse()
	fmt.Println(path)
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}

func TestSelect() {
	x := make(chan int, 1)
	go func() {
		select {
		case <-x:
		case <-time.NewTimer(2 * time.Second).C:
		}
		fmt.Println("ok")
		x <- 1
	}()
	<-x
}

func TestRand() {
	random := rand.NewSource(10)
	fmt.Println(random.Int63())
}

func TestTcp() {
	lister, err := net.Listen("tcp", ":80")
	if err != nil {
		panic(err)
	}
	defer lister.Close()
	conn, err := lister.Accept()
	if err != nil {
		return
	}
	data := make([]byte, 100)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data[:n])
}

func TestEqual() {
	a := struct {
		m, n int
	}{m: 1, n: 2}
	b := struct {
		m, n int
	}{m: 1, n: 2}
	fmt.Println(a == b)
	c := make(chan int, 1)
	c <- 1
	d := make(chan int, 1)
	d <- 1
	fmt.Println(c == d)
	e := map[int]int{}
	f := map[int]int{}
	fmt.Printf("%p %p", e, f)
}

func TestRace() {
	a := 1
	go func() {
		a = 2
	}()
	a = 3
	fmt.Println("a is ", a)

	time.Sleep(time.Second * 2)
}

func TestDiff() {
	a := 1
	b := 2
	fmt.Println(a | b)
}

func testList() {
	var l list.List
	l.PushFront(1)
	fmt.Println(l.Len())
}

func TestSplit() {
	str := "a/b/c/.."
	for _, v := range strings.Split(str, "/") {
		fmt.Println(v)
	}
	fmt.Println("*")
}

func constructArr(a []int) (ret []int) {
	pro := make([]int, len(a))
	back := make([]int, len(a))
	ret = make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if i == 0 {
			pro[i] = a[i]
		} else {
			pro[i] = pro[i-1] * a[i]
		}
	}
	for i := len(a) - 1; i >= 0; i-- {
		if i == len(a)-1 {
			back[i] = a[i]
		} else {
			back[i] = back[i+1] * a[i]
		}
	}
	for i := 0; i < len(a); i++ {
		if i == 0 {
			ret[i] = back[i+1]
		} else if i == len(a)-1 {
			ret[i] = pro[i-1]
		} else {
			ret[i] = pro[i-1] * back[i+1]
		}
	}
	return
}

// TestByteToString 测试[]byte和string的相互转换
func TestByteToString() {
	bts := []byte{'h', 'e', 'l', 'l', 'o'}
	str := "hello"
	//普通转换
	fmt.Println(string(bts), str)
	//reflect和unsafe
	//string -> []byte
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	fmt.Println(*((*[]byte)(unsafe.Pointer(&bh))))
	//[]byte -> string
	fmt.Println(*(*string)(unsafe.Pointer(&bts)))
}
func TestlargestNumber(nums []int) (ret string) {
	sort.Slice(nums, func(i, j int) bool {
		a := strconv.Itoa(nums[i])
		b := strconv.Itoa(nums[j])
		return a+b > b+a
	})
	for i := range nums {
		s := strconv.Itoa(nums[i])
		ret += s
	}
	for i := 0; i < len(ret); i++ {
		if i != len(ret)-1 && ret[i] == '0' {
			continue
		} else {
			return ret[i:]
		}
	}
	return
}
func TestSearch() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(sort.SearchInts(nums, 5))
}

func testW() {
	f, _ := os.OpenFile("D:\\桌面\\2.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	defer f.Close()
	for i := 0; i < 25000; i++ {
		f.Write([]byte("A"))
	}
}

func testSearch() {
	nums := []int{0, 1, 2, 3, 4, 4, 5, 6, 7, 8}
	search := func(n int, f func(idx int) bool) int {
		l, r := 0, n
		for l < r {
			mid := l + (r-l)/2
			if !f(mid) {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l
	}
	idx := search(len(nums), func(idx int) bool { return nums[idx] > 9 })
	fmt.Println(idx)
}

func testInput() {
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	str = strings.Replace(str, "\r\n", "", -1)
	str = strings.Replace(str, "\\", "/", -1)
	fmt.Println(str)
}

func testCode() {
	file, _ := os.Open("D:/桌面/1.txt")
	all, _ := ioutil.ReadAll(file)
	bf := bufio.NewReaderSize(bytes.NewReader(all), 50000)
	bf.ReadByte()
	data := make([]byte, 50000)
	cnt, _ := bf.Read(data)
	fmt.Println(string(data[:cnt]))
}

func testAppend1() {
	nums := []int{1, 2, 3}
	ret := [][]int{}
	ret = append(ret, append([]int{}, nums...))
	nums[0] = 3
	fmt.Println(ret)
}

func testAppend2() {
	nums := []int{1, 2, 3}
	ret := [][]int{}
	ret = append(ret, nums)
	nums[0] = 3
	fmt.Println(ret)
}

func getDistances(arr []int) (res []int64) {
	res = make([]int64, len(arr))
	rets := make(map[int]*list.List, len(arr))
	for i, v := range arr {
		if rets[v] == nil {
			rets[v] = list.New()
		}
		rets[v].PushFront(i)
	}
	for i, v := range arr {
		for e := rets[v].Front(); e != nil; e = e.Next() {
			t := int64(math.Abs(float64(i - e.Value.(int))))
			res[i] += t
			res[e.Value.(int)] += t
		}
		rets[v].Remove(rets[v].Back())
	}
	return
}

func combinationSum(candidates []int, target int) (ret [][]int) {
	result := make([]int, 0, 510)
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ret = append(ret, append([]int{}, result...))
			return
		}
		dfs(target, idx+1)
		if target-candidates[idx] >= 0 {
			result = append(result, candidates[idx])
			dfs(target-candidates[idx], idx)
			result = result[:len(result)-1]
		}
	}
	dfs(target, 0)
	return
}
func translateNum(num int) int {
	s := strconv.Itoa(num)
	a, b := 0, 1
	for i := 0; i < len(s); i++ {
		var c int
		if i-1 >= 0 && s[i-1] > '0' && (s[i-1]-'0')*10+(s[i]-'0') <= 25 {
			c = a + b
		} else {
			c = b
		}
		a, b = b, c
	}
	return b
}
func isDecimals(s string) bool {
	if s == "" {
		return false
	}
	if num := strings.Count(s, "+") + strings.Count(s, "-"); num != 0 {
		if num > 1 || s[0] != '+' && s[0] != '-' {
			return false
		}
		s = s[1:]
	}
	if len(s) > 0 && strings.Count(s, ".") == 1 {
		var ok error
		index := strings.Index(s, ".")
		if index == 0 {
			_, ok = strconv.Atoi(s[1:])
		} else if index == len(s)-1 {
			_, ok = strconv.Atoi(s[:len(s)-1])
		} else {
			_, ok1 := strconv.Atoi(s[:index])
			_, ok2 := strconv.Atoi(s[index+1:])
			if ok1 == nil && ok2 == nil {
				ok = nil
			} else {
				ok = errors.New(" ")
			}
		}
		if ok == nil {
			return true
		}
	}
	return false
}
func isInteger(s string) bool {
	if s == "" {
		return false
	}
	if num := strings.Count(s, "+") + strings.Count(s, "-"); num != 0 {
		if num > 1 || s[0] != '+' && s[0] != '-' {
			return false
		}
		s = s[1:]
	}
	if len(s) > 0 && strings.Count(s, ".") == 0 {
		_, ok := strconv.Atoi(s)
		if ok == nil {
			return true
		}
	}
	return false
}
func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if num := strings.Count(s, "e") + strings.Count(s, "E"); num == 0 || num == 1 {
		index := strings.Index(s, "e")
		if index == -1 {
			index = strings.Index(s, "E")
		}
		if index == -1 {
			return isInteger(s) || isDecimals(s)
		} else if index > 0 && index < len(s)-1 {
			return (isInteger(s[:index]) || isDecimals(s[:index])) && isInteger(s[index+1:])
		}
	}
	return false
}

func fib(n int) int {
	nums := make([]int, 101)
	nums[1], nums[2] = 1, 1
	for i := 3; i <= 100; i++ {
		nums[i] = (nums[i-1]+nums[i-2])%1e9 + 7
	}
	return nums[n]
}
