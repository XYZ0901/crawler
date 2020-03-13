package view

import (
	"crawler/frontend/model"
	model2 "crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("./template.html")
	out, err := os.Create("template.test.html")
	page := model.SearchResult{}
	page.Hits = 123
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items,
			model2.Profile{
				UserName:      "青城",
				Age:           38,
				Height:        160,
				Income:        "12001~20000元",
				Marriage:      "未婚",
				Education:     "大专",
				Occupation:    "其他职业",
				WorkPlace:     "陕西西安",
				Constellation: "双鱼座",
				NativePlace:   "陕西西安",
				Memo:          "婚姻状态是离异！资料无法改，所以在这里说明下！对于我来说，用最好的年华祭奠了对婚姻的向往，成就了今天的沉稳与百变不惊！希望因爱遇到，无关其他！",
				PhotoUrls: []string{
					"https://photo.zastatic.com/images/photo/435618/1742470603/14664428913298946.jpg?scrop=1&crop=1&w=300&h=300&cpos=north",
					"https://photo.zastatic.com/images/photo/435618/1742470603/6447400824770651.jpg?scrop=1&crop=1&w=300&h=300&cpos=north",
					"https://photo.zastatic.com/images/photo/435618/1742470603/418102268957789.jpg?scrop=1&crop=1&w=300&h=300&cpos=north",
					"https://photo.zastatic.com/images/photo/435618/1742470603/417491093102829.jpg?scrop=1&crop=1&w=300&h=300&cpos=north",
				},
				UserUrl: "http://xa.zhenai.com/2489243713.html",
			})
	}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
