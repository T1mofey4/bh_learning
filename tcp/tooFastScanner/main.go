package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 1; i <= 65535; i++ {
// 		wg.Add(1)
// 		go func(j int) {
// 			defer wg.Done()
// 			address := fmt.Sprintf("scanme.nmap.org:%d", j)
// 			fmt.Println(address)
// 			conn, err := net.Dial("tcp", address)
// 			if err != nil {
// 				return
// 			}
// 			conn.Close()
// 			fmt.Printf("%d open\n", j)
// 		}(i)

// 	}
// 	wg.Wait()
// }
