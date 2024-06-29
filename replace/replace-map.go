package replace

import "github.com/Jeffail/gabs/v2"

func ReplaceMap(inputMap *map[string]interface{}, input any, jsonPath *gabs.Container) map[string]interface{} {
	newMap := map[string]interface{}{}
	for k, v := range *inputMap {
		switch t := v.(type) {
		case map[string]interface{}:
			newMap[k] = ReplaceMap(&t, input, jsonPath)
		default:
			key, val, err := GetValue(k, v, input, jsonPath)
			if err != nil {
				//s.logger.debug("replacemapError", zap.Error(err))
				newMap[k] = (*inputMap)[k]
			}
			newMap[key] = val
		}
	}
	return newMap
}
