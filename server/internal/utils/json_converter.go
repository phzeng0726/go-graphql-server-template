package utils

import (
	"encoding/json"
	"fmt"
)

func DomainToDoc(model interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}

	// 將結構體轉換為 JSON 格式的字串
	jsonData, err := json.Marshal(model)
	if err != nil {
		fmt.Println("無法轉換為 JSON 格式:", err)
		return result, err
	}

	// 將 JSON 字串轉換為 map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Println("無法轉換為 map[string]interface{}:", err)
		return result, err
	}

	return result, nil
}
