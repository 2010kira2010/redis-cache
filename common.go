package redis_cache

import "encoding/json"

/*
func mergeInterface(interface1, interface2 string) (string, error) {
	data1 := make(map[string]interface{})
	data2 := make(map[string]interface{})

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
*/

func mergeInterface(interface1, interface2 string) (string, error) {
	data1 := make(map[string]interface{})
	data2 := make(map[string]interface{})

	if err := json.Unmarshal([]byte(interface1), &data1); err != nil {
		return "", err
	}

	if err := json.Unmarshal([]byte(interface2), &data2); err != nil {
		return "", err
	}

	for key, value := range data2 {
		if key == "custom_fields_values" {
			if fields, ok := value.([]interface{}); ok {
				for _, field := range fields {
					if fieldMap, ok := field.(map[string]interface{}); ok {
						if fieldID, ok := fieldMap["field_id"]; ok {
							if fieldIDValue, ok := fieldID.(float64); ok {
								if fields1, ok := data1[key].([]interface{}); ok {
									exists := false
									for i, field1 := range fields1 {
										if field1Map, ok := field1.(map[string]interface{}); ok {
											if fieldID1, ok := field1Map["field_id"]; ok {
												if fieldIDValue1, ok := fieldID1.(float64); ok {
													if int(fieldIDValue) == int(fieldIDValue1) {
														data1[key].([]interface{})[i] = field
														exists = true
														break
													}
												}
											}
										}
									}
									if exists {
										continue
									}
								}

								data1[key] = append(data1[key].([]interface{}), field)
							}
						}
					}
				}
			}
		} else {
			data1[key] = value
		}
	}

	result, err := json.Marshal(data1)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
