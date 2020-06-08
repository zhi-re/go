package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `{"error_code":0,"reason":"成功","result":{"ip":"115.183.152.152","country":"中国","province":"北京","city":"北京","district":null,"isp":"鹏博士"}}`
	fmt.Println(s)
	var resultMap map[string]interface{} /*创建集合 */
	resultMap = make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &resultMap)
	if err != nil {
		fmt.Println(err)
	}
	infoMap := resultMap["result"].(map[string]interface{})
	fmt.Println(infoMap["ip"])
	fmt.Println(infoMap["isp"])
	fmt.Println("======================")
	jsonStr := `{"accessToken":"507b5e08ee444dck887b66bd08672905",
"clientToken":"64e3a5415bfe405d9485f1jf2ea5c68e",
"selectedProfile":{"id":"selID","name":"Bluek404"},
"availableProfiles":[{"id":"测试ava","name":"Bluek404"}]}`
	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)

		mapTmp := dat["selectedProfile"].(map[string]interface{})
		fmt.Println(mapTmp["id"])

		mapTmp2 := (dat["availableProfiles"].([]interface{}))[0].(map[string]interface{})
		//mapTmp3 := mapTmp2[0].(map[string]interface {})
		fmt.Println(mapTmp2["id"])
	}
}
