package redis_cache

import "encoding/json"

func mergeInterface(interface1, interface2 string) (string, error) {
	var data1 map[string]interface{}
	var data2 map[string]interface{}

	if err := json.Unmarshal([]byte(interface1), &data1); err != nil {
		return "", err
	}

	if err := json.Unmarshal([]byte(interface2), &data2); err != nil {
		return "", err
	}

	for key, value := range data2 {
		data1[key] = value
	}

	result, err := json.Marshal(data1)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
