package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
	"strings"
)

func BindQueryWithSlices[T any](ctx *gin.Context, result *T) error {
	err := ctx.BindQuery(&result)
	if err != nil {
		return errors.Wrap(err, "failed to bind query")
	}

	resultValue := reflect.ValueOf(result).Elem()
	resultType := reflect.TypeOf(*result)
	for i := 0; i < resultValue.NumField(); i++ {
		if resultValue.Field(i).Kind() == reflect.Slice {
			fieldValue := ctx.Query(resultType.Field(i).Tag.Get("form"))
			if fieldValue == "" {
				continue
			}
			stringSlice := strings.Split(fieldValue, ",")
			switch resultValue.Field(i).Interface().(type) {
			case []string:
				resultValue.Field(i).Set(reflect.ValueOf(stringSlice))
			case []int64:
				intSlice, err := convertStringSliceToIntSlice(stringSlice)
				if err != nil {
					return errors.Wrap(err, "failed to convert string slice to int slice")
				}
				resultValue.Field(i).Set(reflect.ValueOf(intSlice))
			}
		}
	}

	return nil
}

func convertStringSliceToIntSlice(stringArray []string) ([]int64, error) {
	intArray := make([]int64, 0)
	for _, value := range stringArray {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert string to int")
		}
		intArray = append(intArray, int64(intValue))
	}

	return intArray, nil
}
