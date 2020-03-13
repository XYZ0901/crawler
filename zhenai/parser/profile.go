package parser

import (
	"crawler/engine"
	"crawler/model"
	"crawler/persist"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var count = 1
var ItemSave = persist.ItemSaver()

var (
	//UserNameRe      = regexp.MustCompile(`<div class="username">([^<]*)<span`)
	IdRe      = regexp.MustCompile(`http://[a-z0-9^.]*.zhenai.com/([0-9]*).html`)
	profileRe = map[string]*regexp.Regexp{
		"Age":           regexp.MustCompile("年龄：([^<]*)岁[^<]*</div>"),
		"Height":        regexp.MustCompile("身高：([^<]*)CM[^<]*</div>"),
		"Income":        regexp.MustCompile("月收入：([^<]*)</div>"),
		"Marriage":      regexp.MustCompile("婚况：([^<]*)</div>"),
		"Education":     regexp.MustCompile("学历：([^<]*)</div>"),
		"Occupation":    regexp.MustCompile("职业：([^<]*)</div>"),
		"WorkPlace":     regexp.MustCompile("工作地：([^<]*)</div>"),
		"Constellation": regexp.MustCompile("星座：([^<]*)</div>"),
		"NativePlace":   regexp.MustCompile("籍　贯：([^<]*)</div>"),
		"Memo":          regexp.MustCompile("内心独白：</span>([^<]*)</div>"),
		"PhotoUrls":     regexp.MustCompile(`<img width="100%" data-src="([^"]*)" alt="">`),
	}
)

// 这里可以先做一个map然后转json然后映射到profile
func ParseProfile(contents []byte, UserName, UserUrl string) (result engine.ParserResult) {
	item := model.Item{}
	profile := model.Profile{}
	profile.UserName, profile.UserUrl, item.Type = UserName, UserUrl, "zhenai"
	item.Id = IdRe.FindStringSubmatch(UserUrl)[1]
	for k, v := range profileRe {
		switch k {
		case "Age":
			res := v.FindStringSubmatch(string(contents))
			if len(res) > 1 {
				res[1] = strings.Replace(res[1], " ", "", -1)
				profile.Age, _ = strconv.Atoi(res[1])
			}
		case "Height":
			res := v.FindStringSubmatch(string(contents))
			if len(res) > 1 {
				res[1] = strings.Replace(res[1], " ", "", -1)
				profile.Height, _ = strconv.Atoi(res[1])
			}
		case "Income":
			getData(v, contents, &(profile.Income))
		case "Marriage":
			getData(v, contents, &profile.Marriage)
		case "Education":
			getData(v, contents, &profile.Education)
		case "Occupation":
			getData(v, contents, &profile.Occupation)
		case "WorkPlace":
			getData(v, contents, &profile.WorkPlace)
		case "Constellation":
			getData(v, contents, &profile.Constellation)
		case "NativePlace":
			getData(v, contents, &profile.NativePlace)
		case "Memo":
			getData(v, contents, &profile.Memo)
		case "PhotoUrls":
			res := v.FindAllStringSubmatch(string(contents), -1)
			if len(res) > 0 {
				for _, u := range res {
					profile.PhotoUrls = append(profile.PhotoUrls, u[1])
				}
			}
		}
	}
	if count%100 == 0 {
		time.Sleep(3 * time.Second)
	}
	item.Profile = profile
	fmt.Println("Item Saver: got item", count, "th ", profile)
	count++
	ItemSave <- item
	return
}

func getData(v *regexp.Regexp, contents []byte, profile *string) {
	res := v.FindStringSubmatch(string(contents))
	if len(res) > 1 {
		res[1] = strings.Replace(res[1], " ", "", -1)
		*profile = res[1]
	}
}
