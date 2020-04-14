package fetcher

import (
	"bufio"
	"crawler/enginetype"

	"crawler/parser"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)
//将数据转换为utf-8
func FetcherContent(Parameter enginetype.Parameter)enginetype.Data{
	resp, err := http.Get(Parameter.Address)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//判断状态是否正确
	if resp.StatusCode != http.StatusOK {
		panic(err)
	}
	//判断数据编码格式
	e:=DecideCharset(resp.Body)
	//在拿次数据因为resp.byte读到1024
	newresp, err := http.Get(Parameter.Address)
	if err != nil {
		panic(err)
	}
	//转化为utf8格式
	utf8:=transform.NewReader(newresp.Body,e.NewDecoder())
	r,err:=ioutil.ReadAll(utf8)
	if err!=nil {
		panic(err)
	}
	url:=enginetype.Data{
		r,
		nil,
		parser.SelectContent,
		Parameter,
	}
	return url
}

func DecideCharset(body io.ReadCloser)encoding.Encoding{
	bufio,err:=bufio.NewReader(body).Peek(1024)
	if err!=nil{
		panic(err)
	}
	e,_,_:=charset.DetermineEncoding(bufio,"")
	return e
}