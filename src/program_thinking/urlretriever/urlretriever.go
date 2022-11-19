package urlretriever

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct{} //不需要任何参数,仅提供方法

func (Retriever) Get(url string) string { //因为只需要Retriver提供方法,所以直接不给它命名
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body) //暂时不用err参数,以后再说错误处理
	return string(bytes)
}
