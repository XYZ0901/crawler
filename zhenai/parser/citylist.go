package parser

import (
	"crawler/engine"
	"encoding/json"
	"log"
)

type Citys map[string]string
type Province map[string]Citys

var (
	res []interface{}
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
	url := `.zhenai.com/jiaoyou`
	for kp, _ := range province {
		citys := province[kp]
		for kc, _ := range citys {
			result.Items = append(result.Items, "City "+kc)
			result.Requests = append(result.Requests, engine.Request{
				Url:        `http://` + province[kp][kc] + url,
				ParserFunc: ParseCity,
			})
			//break
		}
		break
	}
	return
}
