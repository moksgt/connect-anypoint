package main

import (
	"fmt"
	"sync"
	"time"
)

/**
API Flow Design
1. Load config file of API Flow
Data structure
{ api_flow_id:"",
  api_flow_name:"",
  parameters:{}
  api_columns:[{[{seq:"",api_id:"",pre_param_mapping:[{oldParam1:newParam1}]},{[{}]}]},{[{},{}]}]
}
2. Execute a column APIs by a column config, before parameters processing and mapping
3. Waiting result then processing and mapping
4. Return result
*/
func air_flow() {
	// 1. Load config file of API Flow
	// 2. Execute a column APIs by a column config, before parameters processing and mapping
	// 3. Waiting result then processing and mapping
	// 4. Return result
}

func job(index int) int {
	time.Sleep(time.Second)
	return index
}

func main() {

	start := time.Now()
	num := 5
	retCh := make(chan int)
	wg := sync.WaitGroup{}
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(param int) {
			defer wg.Done()
			retCh <- job(param)
		}(i)
	}

	go func() {
		defer close(retCh)
		wg.Wait()
		println("wg end...")
	}()

	for item := range retCh {
		fmt.Println("收到结果：", item)
	}

	end := time.Since(start)
	fmt.Println("耗时：", end.String())

}
