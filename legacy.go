package pprnt

type deprecatedHolder struct {
}

//CleanMap ...
func (dh *deprecatedHolder) CleanMap(clmap map[string]interface{}) map[string]interface{} {
	newMap := map[string]interface{}{}
	for key, value := range clmap {
		if value == nil {
			continue
		}

		switch value.(type) {
		case map[string]interface{}:
			childMap := value.(map[string]interface{})
			newMap[key] = dh.CleanMap(childMap)
			break
		case []interface{}:
			arr := value.([]interface{})
			newMap[key] = dh.CleanArray(arr)
			break
		default:
			newMap[key] = value
			break
		}
	}
	return newMap
}

//CleanArray ...
func (dh *deprecatedHolder) CleanArray(clarr []interface{}) []interface{} {
	newArr := []interface{}{}
	for _, each := range clarr {
		if each == nil {
			continue
		}
		switch each.(type) {
		case map[string]interface{}:
			childMap := each.(map[string]interface{})
			newArr = append(newArr, dh.CleanMap(childMap))
			break
		case []interface{}:
			arr := each.([]interface{})
			newArr = append(newArr, dh.CleanArray(arr))
			break
		default:
			newArr = append(newArr, each)
			break
		}
	}

	return newArr
}
