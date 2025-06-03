package main

import (
	"fmt"
	"time"
	"sync"
)

func plus(x, y int) int{
	return x + y
}

func hello(){
	fmt.Println("表示成功")
	return
}

func multireturn(a, b int) (int, int){
	q := a / b
	r := a * b
	return q, r
}

// 並行処理対象関数
func normal(s string,){
	for i := 0; i < 5; i++{
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// 並行処理対象関数
func gorutine(s string, wg *sync.WaitGroup){

	for i := 0; i < 5; i++{
		time.Sleep(200 * time.Millisecond)
		fmt.Println(s)
	}
		wg.Done()
}


func main () {
	result := plus(1, 2)
	fmt.Println(result)
	hello()
	q, r := multireturn(4, 2)
	fmt.Println(q, r)
	var wg sync.WaitGroup
	//wgがdoneされるまで待ってねということ
	wg.Add(1)
	normal("hello")
	go gorutine("world", &wg)
	wg.Wait()
}
