package news

var (
	collect chan (string)
	result  chan ([]Topic)

	request chan (string)
)

func init() {
	collect = make(chan (string))
	result = make(chan ([]Topic))
	request = make(chan (string))
}

func Collect(category string) {
	collect <- category
}

func Result(category string) []Topic {
	request <- category
	return <-result
}

func (a Archive) Serve() {
	for {
		select {
		case category := <-collect:
			a.collect(category)
		case category := <-request:
			result <- a.result(category)
		}
	}
}
