package engine

import "crawler/fetcher"

func worker(r Request) (ParserResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}
