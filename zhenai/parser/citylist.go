package parser

import (
	"crawler/engine"
	"encoding/json"
	"log"
)

type Citys map[string]string
type Province map[string]Citys

var (
	res      []interface{}
	afterUrl = `.zhenai.com/jiaoyou`
)

func GetProvinceCitys(all []byte) (result engine.ParserResult) {
	err := json.Unmarshal(all, &res)
	if err != nil {
		log.Fatal(err)
	}
	jsonbyte, err := json.Marshal(res[2])
	var province Province
	err = json.Unmarshal(jsonbyte, &province)
	if err != nil {
		log.Fatal(err)
	}
	for kp, _ := range province {
		citys := province[kp]
		for kc, _ := range citys {
			cityName := province[kp][kc]
			result.Requests = append(result.Requests, engine.Request{
				Url: `http://` + cityName + afterUrl,
				ParserFunc: func(c []byte) engine.ParserResult {
					return ParseCity(c, cityName)
				},
			})
		}
	}
	return
}
