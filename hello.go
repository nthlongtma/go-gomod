package hello // import "github.com/nthlongtma/go-gomod"

import "rsc.io/quote/v3"

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
