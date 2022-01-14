package validator

import "reflect"

func Empty(compared, empty interface{}) bool {
	return reflect.DeepEqual(compared, empty)
}
