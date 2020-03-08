package parser

import (
	"crawler/engine"
	"regexp"
)

var UserRe = regexp.MustCompile(`<div class="name fl"><a href="([^"])*" target="_blank" class="trans">([^<]*)</a></div>`)

func ParseCity(contents []byte) (result engine.ParserResult) {
	matches := UserRe.FindAllStringSubmatch(string(contents), -1)

	for _, m := range matches {
		result.Items = append(result.Items,"User "+m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        m[1],
			ParserFunc: engine.NilParser,
		})
	}
	return
}
