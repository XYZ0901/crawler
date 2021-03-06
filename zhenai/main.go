package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

// 返回citys类型为map[string]map[string]string
const (
	citysUrl = "https://china.zhenai.com/city.html?do_action=city&MmEwMD=4Js.PRbtdlYRNICHUbXoDCysalppoQy_wYUJs.4VRfbXb_d858vjeXChMiLzkXtitCORJuPYw65JiJ3KNu.tu3OOu1dl0Mw1C2VMI_yenqit.KjiJ7DdBGQGOPyWqhUiJ6owla1SuA5M.4l0ta2AuoWwwysgJ52Uws5I3nZxlqbWrnoQyD3q9zKsF95s_pg1C867jOvAq.AjFyiG295.MC2.LVt8XUVOvwBV..HUpMAPGW.S59ivcI5GuLEmIzaBB9hGCjwWA210_grH5VGghGcoMEoTXGGr4UIexkSrYn947H_gCcmKM7_WFYeIIbo_YodAvoM2U7YD2XNyEgAdx5ZIJPW6123V7dGtRVPcDtvfmrdNaMuKudSaGMaHD18Qj30a"
	// 深圳
	szUrl = "https://sz.zhenai.com/jiaoyou"
)

// 返回html
//const citysUrl = "https://www.zhenai.com/zhenghun"

func main() {
	// the simple vision
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:citysUrl,
	//	ParserFunc: parser.GetProvinceCitys,
	//})

	//the concurrent vision
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	// for china
	e.Run(
		engine.Request{
			Url:        citysUrl,
			ParserFunc: parser.GetProvinceCitys,
		},
	)

	//for one city
	//e.Run(
	//	engine.Request{
	//		Url:        szUrl,
	//		ParserFunc: func(c []byte) engine.ParserResult {
	//		return parser.ParseCity(c,"sz")
	//	},
	//	},
	//)

}
