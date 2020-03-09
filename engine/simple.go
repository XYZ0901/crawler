package engine

import (
	"crawler/fetcher"
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

func worker(r Request) (ParserResult, error) {
	log.Println("Fetching ", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Println("Fetcher: error fetching url ", r.Url, ":", err)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}
