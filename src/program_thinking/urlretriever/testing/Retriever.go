package testing

type Retriever struct{} //不需要任何参数,仅提供方法

func (Retriever) Get(url string) string { //因为只需要Retriver提供方法,所以直接不给它命名
	return "this is test"
}
