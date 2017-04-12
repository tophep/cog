package cog

//	TODO
//	1. Add more colors
//	2. Make function to easily log with any available color -- e.g. LogColor("pink", data...)
//	3. Add functionality to customize the way certain data types are printed -- e.g. have a map from data type to stringifying function in toString method

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	normalColor  = ""
	errorColor   = "\x1b[31m"
	warningColor = "\x1b[33m"
	alertColor   = "\x1b[36m"
	logEndColor  = "\x1b[0m"
)

func Log(data ...interface{}) {
	log(normalColor, data...)
}

func Error(data ...interface{}) {
	log(errorColor, data...)
}

func Warning(data ...interface{}) {
	log(warningColor, data...)
}

func Alert(data ...interface{}) {
	log(alertColor, data...)
}

func log(color string, data ...interface{}) {
	prettyData := ToString(data...)
	fmt.Println(color, "\n", time.Now(), "\n\n", prettyData, "\n", logEndColor)
}

func ToString(data ...interface{}) string {
	str := ""
	for _, thing := range data {
		str += toString(thing) + " "
	}
	return str

}

func toString(thing interface{}) string {
	switch thing.(type) {
	case string:
		return thing.(string)
	case []byte:
		return string(thing.([]byte))
	case error:
		return thing.(error).Error()
	default:
		str, err := json.MarshalIndent(thing, "", "  ")
		if err != nil {
			Error("Marshalling inside toString() Failed:", err)
		}
		return string(str)
	}
}
