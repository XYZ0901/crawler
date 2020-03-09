package parser

import (
	"crawler/engine"
	"regexp"
)

var UserRe = regexp.MustCompile(`<div class="name fl"><a href="([^"]*)" target="_blank" class="trans">([^<]*)</a></div>`)

func ParseCity(contents []byte) (result engine.ParserResult) {
	matches := UserRe.FindAllStringSubmatch(string(contents), -1)

	for _, m := range matches {
		UserUrl := "http:" + m[1]
		UserName := m[2]
		result.Items = append(result.Items, "User "+UserName)
		result.Requests = append(result.Requests, engine.Request{
			Url: UserUrl,
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseProfile(c, UserName, UserUrl)
			},
		})
	}
	return
}
