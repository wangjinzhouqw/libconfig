package libconfig

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseConfigFile(t *testing.T) {
	jsonStr,err := ParseConfigFile("/Users/wangjz/Desktop/work/project/opensource/janus-gateway/conf/janus.jcfg.sample.in")
	if err!=nil{
		fmt.Errorf(err.Error())
	}
	fmt.Println(jsonStr)
	fmt.Println("=========================")
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonStr),&jsonMap)
	fmt.Println(jsonMap)
	//ParseJsonConfigFile("/Users/wangjz/Desktop/janus.jcfg.sample.in")
}