package cleaner

import "log"

//CleanMap ...
func CleanMap(clmap map[string]interface{}) map[string]interface{} {
	newMap := map[string]interface{}{}
	for key, value := range clmap {
		if value == nil {
			log.Printf("> Skipped field '%+v' cuz nil", key)
			continue
		}

		switch value.(type) {
		case map[string]interface{}:
			childMap := value.(map[string]interface{})
			newMap[key] = CleanMap(childMap)
			break
		case []interface{}:
			arr := value.([]interface{})
			newMap[key] = CleanArray(arr)
			break
		default:
			newMap[key] = value
			break
		}
	}
	return newMap
}

//CleanArray ...
func CleanArray(clarr []interface{}) []interface{} {
	newArr := []interface{}{}
	for _, each := range clarr {
		if each == nil {
			continue
		}
		switch each.(type) {
		case map[string]interface{}:
			childMap := each.(map[string]interface{})
			newArr = append(newArr, CleanMap(childMap))
			break
		case []interface{}:
			arr := each.([]interface{})
			newArr = append(newArr, CleanArray(arr))
			break
		default:
			newArr = append(newArr, each)
			break
		}
	}

	return newArr
}
