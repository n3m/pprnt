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

		switch value := value.(type) {
		case map[string]interface{}:
			newMap[key] = dh.CleanMap(value)
		case []interface{}:
			newMap[key] = dh.CleanArray(value)
		default:
			newMap[key] = value
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
		switch each := each.(type) {
		case map[string]interface{}:
			newArr = append(newArr, dh.CleanMap(each))
		case []interface{}:
			newArr = append(newArr, dh.CleanArray(each))
		default:
			newArr = append(newArr, each)
		}
	}

	return newArr
}
