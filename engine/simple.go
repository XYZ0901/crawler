package engine

import (
	"log"
)

type SimpleEngine struct {
}

func (SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		// 保证程序退出
		//if parseResult.Items == nil {
		//	return
		//}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Println("Got item ", item)
		}
	}
}
