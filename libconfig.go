package libconfig

import (
	bufio2 "bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func ParseConfigFile(f string) (string,error){
	fmt.Println(f)
	file,err := os.Open(f)
	if err != nil{
		fmt.Errorf("%s",err.Error())
	}
	defer file.Close()

	var jsonStr string
	jsonStr += "{"
	br := bufio2.NewReader(file)
	for {
		line,err := br.ReadString('\n')
		if err!=nil || err==io.EOF {
			break;
		}
		strList := strings.Split(line,"#")
		strVaule := strings.TrimSpace(strList[0])
		strFinalVaule := strings.Replace(strVaule,"=",":",-1)
		if len(strFinalVaule)>0{
			configEle := strings.Split(strFinalVaule,":")
			fmt.Println(strFinalVaule)
			fmt.Println(len(configEle))
			if len(configEle)>1 {
				jsonEle := "\"" + strings.TrimSpace(configEle[0]) + "\"" +  ":" + configEle[1]
				jsonStr += jsonEle
			} else {
				jsonStr += strFinalVaule
			}

			if !strings.HasSuffix(strFinalVaule,"{"){
				jsonStr += ","
			}
		}
	}
	jsonStr += "}"
	fmt.Println("------------------------")
	fmt.Println(jsonStr)
	jsonMap := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(jsonStr),jsonMap)
	if err1 != nil {
		fmt.Println("err1")
		fmt.Println(err1.Error())
	}
	fmt.Println(jsonMap["general"])
	fmt.Println("------------------------")
	fmt.Println(jsonMap)
	return jsonStr,nil
}

func ParseJsonConfigFile(f string) (string,error){
	fmt.Println(f)
	file,err := os.Open(f)
	if err != nil{
		fmt.Errorf("%s",err.Error())
	}
	defer file.Close()

	var jsonStr string
	br := bufio2.NewReader(file)
	for {
		line,err := br.ReadString('\n')
		if err!=nil || err==io.EOF {
			break;
		}
		jsonStr += line
	}

	fmt.Println(jsonStr)
	jsonMap := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(jsonStr),&jsonMap)
	if err1 != nil {
		fmt.Println("err1")
		fmt.Println(err1.Error())
	}
	fmt.Println(jsonMap)
	return jsonStr,nil
}
