package helper

import (
    "reflect"
    "strings"
    "unicode"
    "unicode/utf8"
)

func GetJSONFormat(o interface{}) interface{} {
    if strings.Contains(reflect.TypeOf(o).String(), "[]model") {
        formmated := []interface{}{}

        vs := reflect.ValueOf(o)
        for i := 0; i < vs.Len(); i++ {
            v := vs.Index(i).Interface()
            formmated = append(formmated, GetJSONFormat(v))
        }

        return formmated
    } else {
        formmated := map[string]interface{}{}

        v := reflect.ValueOf(o)
        for i := 0; i < v.NumField(); i++ {
            key := v.Type().Field(i).Name
            value := v.Field(i).Interface()

            if strings.Contains(reflect.TypeOf(value).String(), "[]model") {
                s := reflect.ValueOf(value)

                nestedValue := []interface{}{}
                for j := 0; j < s.Len(); j++ {
                    nestedValue = append(nestedValue, GetJSONFormat(s.Index(j).Interface()))
                }

                value = nestedValue
            }

            formmated[lowerFirstChar(key)] = value
        }

        return formmated
    }
}

func lowerFirstChar(s string) string {
    if s == "" {
        return ""
    }
    r, n := utf8.DecodeRuneInString(s)
    return string(unicode.ToLower(r)) + s[n:]
} 