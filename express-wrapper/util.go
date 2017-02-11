package express

import (
	"encoding/json"
	"errors"
)

func convertToMapOfStringSlices(i interface{}) (m map[string][]string, err error) {
	m = make(map[string][]string)

	if i == nil {
		err = errors.New("Got nil trying to convert interface to map of string slices")
		return
	}

	tempMap, isMap := i.(map[string]interface{})

	if isMap {

		for k, v := range tempMap {

			s, isString := v.(string)
			if isString {
				m[k] = []string{s}
				break
			}

			si, isSlice := v.([]interface{})

			if isSlice {

				sss, serr := convertToStringSlice(si)
				if serr == nil {
					m[k] = sss
				}
			}
		}
	} else {
		err = errors.New("Not a valid map")
	}
	return
}

func convertToMapOfStrings(i interface{}) (m map[string]string, err error) {
	m = make(map[string]string)

	if i == nil {
		err = errors.New("Got nil trying to convert interface to map of string slices")
		return
	}

	tempMap, isMap := i.(map[string]interface{})

	if isMap {
		for k, v := range tempMap {
			s, isString := v.(string)
			if isString {
				m[k] = s
			}
		}
	} else {
		err = errors.New("Not a valid map")
	}
	return
}

func convertToStringSlice(i interface{}) (ss []string, err error) {

	ss = []string{}

	if i == nil {
		err = errors.New("Got nil trying to convert interface to string slice")
		return
	}

	si, isSlice := i.([]interface{})

	if isSlice {
		for _, v := range si {
			s, isString := v.(string)
			if isString {
				ss = append(ss, s)
			}
		}
	} else {
		err = errors.New("Not a valid slice")
	}
	return
}

func convertToBytes(i interface{}) (b []byte, err error) {
	if i == nil {
		err = errors.New("Got nil trying to convert interface to bytes")
		return
	}

	b, err = json.Marshal(i)
	return
}
