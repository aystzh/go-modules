package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	st "github.com/aystzh/go-modules/crawl/strcts"
)

func main() {
	//封装请求参数
	params := st.RequestParams{
		Id:          "YTHSHH001",
		IdType:      "USERID",
		Index:       1,
		ProductType: "",
		Rows:        200,
		Status:      "0",
	}

	//将结构体对象转换成json字符串
	jsonPar, err := json.Marshal(&params)
	checkErr(err)
	resp, errPost := http.Post("https://zhebwx.cic.cn/product/load/list", "application/json", bytes.NewBuffer(jsonPar))
	checkErr(errPost)
	defer resp.Body.Close()

	postResult, _ := ioutil.ReadAll(resp.Body)
	//拿到response返回的json字符串
	//fmt.Println("result :", string(result))

	var postStr st.PostData
	//解析json字符串到结构体当中
	errors := json.Unmarshal(postResult, &postStr)
	checkErr(errors)

	//fmt.Println(postStr)
	//循环遍历list节点数据
	for idx, val := range postStr.Data.ProductList.Data {
		fmt.Printf("idx:%v 保险名称:%v  二级编码:%v \n", idx, val.CName, val.CCode)
		ccode := val.CCode
		getUrl := strings.Join([]string{"https://zhebwx.cic.cn/product/load/instance", ccode}, "/")
		getRep, gErr := http.Get(getUrl)
		checkErr(gErr)
		defer getRep.Body.Close()
		getResult, _ := ioutil.ReadAll(getRep.Body)

		//fmt.Println(string(getResult))

		var getStr st.GetData
		gjErr := json.Unmarshal(getResult, &getStr)
		checkErr(gjErr)

		//0777表示：创建了一个普通文件，所有人拥有所有的读、写、执行权限
		//0666表示：创建了一个普通文件，所有人拥有对该文件的读、写权限，但是都不可执行
		//0644表示：创建了一个普通文件，文件所有者对该文件有读写权限，用户组和其他人只有读权限，都没有执行权限
		output2File := "./data/" + val.CName + val.CCode + ".txt"
		errGWrite := ioutil.WriteFile(output2File, getResult, 0666)
		checkErr(errGWrite)
	}
	errPWrite := ioutil.WriteFile("./data/one.txt", postResult, 0666)
	checkErr(errPWrite)
}

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
