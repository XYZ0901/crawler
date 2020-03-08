package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items []interface{}
}

// 为处理Request中ParserFunc函数为nil时的错误 所以设定一个nil函数
func NilParser([]byte) ParserResult {
	return ParserResult{}
}