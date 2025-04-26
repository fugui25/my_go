package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 定义缓冲区大小
const bufferSize = 10

// 生产者函数，向通道中发送数据
func producer(dataCh chan<- int, wg *sync.WaitGroup, doneCh <-chan struct{}) {
	defer wg.Done()
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	for {
		select {
		case <-doneCh:
			return
		default:
			num := rand.Intn(100) // 生成0到99之间的随机数
			fmt.Printf("Producer sent data: %d\n", num)
			select {
			case dataCh <- num:
			default:
				fmt.Println("Producer buffer is full, skipping this data")
			}
		}
	}
}

// 消费者函数，从通道中接收数据
func consumer(dataCh <-chan int, wg *sync.WaitGroup, doneCh chan<- struct{}) {
	defer wg.Done()
	for {
		select {
		case num, ok := <-dataCh:
			if !ok {
				// 通道关闭，退出循环
				doneCh <- struct{}{}
				return
			}
			fmt.Printf("Consumer received data: %d\n", num)
			time.Sleep(time.Millisecond * 500) // 模拟消费时间
		}
	}
}

func main() {
	// 创建有缓冲的通道
	dataCh := make(chan int, bufferSize)
	var wg sync.WaitGroup
	doneCh := make(chan struct{})

	// 启动3个生产者
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go producer(dataCh, &wg, doneCh)
	}

	// 启动2个消费者
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go consumer(dataCh, &wg, doneCh)
	}

	time.Sleep(time.Second * 5) // 运行一段时间

	// 关闭通道，通知生产者和消费者结束
	close(doneCh)
	// 等待所有生产者和消费者完成
	wg.Wait()
	close(dataCh)

	fmt.Println("All producers and consumers have finished.")
}
