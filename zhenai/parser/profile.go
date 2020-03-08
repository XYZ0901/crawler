package parser

import (
	"crawler/engine"
	"crawler/model"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	//UserNameRe      = regexp.MustCompile(`<div class="username">([^<]*)<span`)
	AgeRe           = regexp.MustCompile("年龄：([^<]*)岁[^<]*</div>")
	HeightRe        = regexp.MustCompile("身高：([^<]*)CM[^<]*</div>")
	IncomeRe        = regexp.MustCompile("月收入：([^<]*)</div>")
	MarriageRe      = regexp.MustCompile("婚况：([^<]*)</div>")
	EducationRe     = regexp.MustCompile("学历：([^<]*)</div>")
	OccupationRe    = regexp.MustCompile("职业：([^<]*)</div>")
	WorkPlaceRe     = regexp.MustCompile("工作地：([^<]*)</div>")
	ConstellationRe = regexp.MustCompile("星座：([^<]*)</div>")
	NativePlaceRe   = regexp.MustCompile("籍　贯：([^<]*)</div>")
	MemoRe          = regexp.MustCompile("内心独白：</span>([^<]*)</div>")
	PhotoUrlsRe     = regexp.MustCompile(`<img width="100%" data-src="([^"]*)" alt="">`)
	profileRe       = map[string]*regexp.Regexp{
		//"UserName":      UserNameRe,
		"Age":           AgeRe,
		"Height":        HeightRe,
		"Income":        IncomeRe,
		"Marriage":      MarriageRe,
		"Education":     EducationRe,
		"Occupation":    OccupationRe,
		"WorkPlace":     WorkPlaceRe,
		"Constellation": ConstellationRe,
		"NativePlace":   NativePlaceRe,
		"Memo":          MemoRe,
		"PhotoUrls":     PhotoUrlsRe,
	}
)

// 这里可以用map转json在直接映射到profile里面 但是这里为了省事 如果需要操作的
// 字段大量的话，可以尝试上面那个办法
//可以将case里的东西都封装成同一个函数 但是现在2：23 要去睡觉了
func ParseProfile(contents []byte, UserName string) (result engine.ParserResult) {
	profile := model.Profile{}
	profile.UserName = UserName
	for k, v := range profileRe {
		switch k {
		//case "UserName":
		//	res := v.FindStringSubmatch(string(contents))
		//	if len(res) > 1 {
		//		res[1] = strings.Replace(res[1]," ","",-1)
		//		profile.UserName = res[1]
		//	}
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
			res := v.FindStringSubmatch(string(contents))
			if len(res) > 1 {
				res[1] = strings.Replace(res[1], " ", "", -1)
				profile.Income = res[1]
			}
		case "Marriage":
			res := v.FindStringSubmatch(string(contents))
			if len(res) > 1 {
				res[1] = strings.Replace(res[1], " ", "", -1)
				profile.Marriage = res[1]
			}
		case "Education":
			res := v.FindStringSubmatch(string(contents))
			if len(res) > 1 {
				res[1] = strings.Replace(res[1], " ", "", -1)
				profile.Education = res[1]
			}
		case "Occupation":
			res := v.FindStringSubmatch(string(contents))
			if len(res) > 1 {
				res[1] = strings.Replace(res[1], " ", "", -1)
				profile.Occupation = res[1]
			}
		case "WorkPlace":
			res := v.FindStringSubmatch(string(contents))
			if len(res) > 1 {
				res[1] = strings.Replace(res[1], " ", "", -1)
				profile.WorkPlace = res[1]
			}
		case "Constellation":
			res := v.FindStringSubmatch(string(contents))
			if len(res) > 1 {
				res[1] = strings.Replace(res[1], " ", "", -1)
				profile.Constellation = res[1]
			}
		case "NativePlace":
			res := v.FindStringSubmatch(string(contents))
			if len(res) > 1 {
				res[1] = strings.Replace(res[1], " ", "", -1)
				profile.NativePlace = res[1]
			}
		case "Memo":
			res := v.FindStringSubmatch(string(contents))
			if len(res) > 1 {
				res[1] = strings.Replace(res[1], " ", "", -1)
				profile.Memo = res[1]
			}
		case "PhotoUrls":
			res := v.FindAllStringSubmatch(string(contents), -1)
			if len(res) > 0 {
				for _, u := range res {
					profile.PhotoUrls = append(profile.PhotoUrls, u[1])
				}
			}
		}
	}
	fmt.Println(profile)
	return
}
