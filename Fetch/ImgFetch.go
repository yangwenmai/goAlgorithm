package main

import (
	"log"
	"time"
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"regexp"
	"container/list"
)

var websites []string

func main() {

	/*websites = []string{
		"http://www.tmall.com/"}
	// 并发5个运行
	pnum := 10 // 默认设置10个并发测试
	parallelRequests(pnum, websites)*/
	init_url := "http://www.baidu.com"
	getUrls(init_url)




}


func getUrls(init_url string) {

	f,err1 := os.OpenFile("123.txt",os.O_RDWR|os.O_APPEND|os.O_CREATE,os.ModePerm)
	queue := list.New()

	if err1!=nil{
		panic(err1)
		return
	}
	i:=1
	for {
		urls := fetchHtml(i,init_url)
		for _,url:=range urls{

		}
		f.WriteString(string(buf[:n]))
	}

	defer f.Close()

}

func fetchHtml(i int,url string) ([]string) {
	startTime := time.Now().UnixNano()
	resp, _ := http.Get(url)
	defer (func() {
		resp.Body.Close()
		processed := float32(time.Now().UnixNano()-startTime) / 1e9
		log.Printf("NO: %02d, url: %s, end, time: %.3fs", i, url,  processed)
	})()
	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Println(err.Error())
	}
	html:=string(body)
	r, _ := regexp.Compile("href=\"(.+?)\"")

	return r.FindAllString(html,-1)
}

func getImgs() {

}


func parallelRequests(pnum int, websites []string) { // 并行抓取
	total := len(websites)
	if pnum <= 0 { // 在设定为0时，全部并发
		pnum = total
	}
	if pnum > total {
		pnum = total
	}
	startTime := time.Now().UnixNano()
	fetchData := make(map[string]string, total) // 反馈抓取后的数据结果，可以写入文件查看
	execChans := make(chan bool, pnum)          // 控制并发数量的通道，第二个参数指定通道可以容纳的数量，会阻塞执行
	doneChans := make(chan bool, 1)             // 用来传递完成信号，完成信号只需要设定容纳一位即可，完成后再次读取新的任务
	for i := 0; i < total; i++ {
		go requests(i, websites[i], execChans, doneChans, fetchData) // 以协程方式运行
	}

	for i := 0; i < total; i++ {
		r := <-doneChans // 完成一个，同时获取下一个任务
		<-execChans      // 紧接着读取下一个任务，类是于beanstalkd的任务分发机制
		if !r {          // 获取失败时，打印该网址失败。
			log.Printf("第 %s 项获取失败", i)
		}
	}
	close(doneChans)                                            // 关闭完成信号
	close(execChans)                                            // 关闭执行信号
	processed := float32(time.Now().UnixNano()-startTime) / 1e9 // 统计总耗时
	log.Printf("全部完成，并发数量：%d, 共耗时：%.3fs", pnum, processed)
	log.Printf("data: %q", fetchData)
}

func requests(i int, url string, execChans chan bool, doneChans chan bool, fetchData map[string]string) {
	execChans <- true // 放在函数的开始处，用来阻塞执行，如果通道里的数量超过设定数量，在没有读取完成前，不会运行
	log.Printf("NO: %02d, url: %s, start...", i, url)
	isOk := false
	startTime := time.Now().UnixNano()
	resp, _ := http.Get(url)
	defer (func() {
		resp.Body.Close()
		doneChans <- isOk
		processed := float32(time.Now().UnixNano()-startTime) / 1e9
		log.Printf("NO: %02d, url: %s, end, status: %t, time: %.3fs", i, url, isOk, processed)
	})()
	body, err := ioutil.ReadAll(resp.Body)
	html:=string(body)
	r, _ := regexp.Compile("href=\"(.+?)\"")

	urls := r.FindAllString(html,-1)


	lens := len(body)
	fetchData[url] = fmt.Sprintf("len: %d", lens)
	if err == nil {
		isOk = true
	}
}

func getImg(resp *http.Response)  {

	if resp.StatusCode == http.StatusOK{
		fmt.Println(resp.StatusCode)
	}

	defer resp.Body.Close()

	buf := make([]byte,1024)


	f,err1 := os.OpenFile("123.txt",os.O_RDWR|os.O_APPEND|os.O_CREATE,os.ModePerm)

	if err1!=nil{
		panic(err1)
		return
	}

	for {
		n,_ :=resp.Body.Read(buf)
		if 0==n{
			break
		}
		f.WriteString(string(buf[:n]))
	}

	defer f.Close()

}