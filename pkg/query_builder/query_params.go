package query_builder

import "strconv"

type QueryParams map[string]string

func (qp QueryParams) GetInt(key string) interface{} {
	mapVal := qp[key]

	if mapVal == "" {
		return ""
	} else {
		val, err := strconv.Atoi(mapVal)
		if err != nil {
			return ""
		}
		return val
	}
}

func (qp QueryParams) GetString(key string) interface{} {
	return qp[key]
}
