package parser

import (
	"crawler/engine"
	"regexp"
	"strconv"
)

var UserRe = regexp.MustCompile(`<div class="name fl"><a href="([^"]*)" target="_blank" class="trans">([^<]*)</a></div>`)
var pagenumRe = regexp.MustCompile(`<span class="tips">共([0-9]*)页&nbsp;&nbsp;到第`)
var CityOver = make(map[string]bool)

func ParseCity(contents []byte, cityName string) (result engine.ParserResult) {
	if !CityOver[cityName] {
		pagestr := pagenumRe.FindStringSubmatch(string(contents))
		pagenum, _ := strconv.Atoi(pagestr[1])
		if pagenum > 1 {
			for i := 2; i <= pagenum; i++ {
				result.Requests = append(result.Requests, engine.Request{
					Url: `http://` + cityName + afterUrl + "/index_" + strconv.Itoa(i) + ".html",
					ParserFunc: func(c []byte) engine.ParserResult {
						return ParseCity(c, cityName)
					},
				})
			}
			CityOver[cityName] = true
		}
	}

	matches := UserRe.FindAllStringSubmatch(string(contents), -1)

	for _, m := range matches {
		UserUrl := "http:" + m[1]
		UserName := m[2]
		result.Requests = append(result.Requests, engine.Request{
			Url: UserUrl,
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseProfile(c, UserName, UserUrl)
			},
		})
	}
	return
}
