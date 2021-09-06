package service

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"reflect"
	"regexp"
)

//@function: WhatAv
//@description: 根据json查询判断
//@param: tasklist string
//@return: list interface{}

func WhatAv(tasklist string) (list interface{}) {
	r, _ := regexp.Compile(`.+\.exe`)
	allProcess := r.FindAllString(tasklist, -1)
	fmt.Println(allProcess)
	resultList := []map[string]interface{}{}
	gq := gojsonq.New().File("./resource/avlist.json")
	for _, v := range allProcess {
		processResult := gq.From("avs").Select("Antivirus", "processes", "url").Where("processes", "contains", v).Get()
		if len(processResult.([]interface{})) == 0 {
			gq.Reset()
		} else {
			gq.Reset()
			avResult := processResult.([]interface{})[0].(map[string]interface{})
			isIn := false
			for _, j := range resultList {
				if reflect.DeepEqual(j, avResult) {
					isIn = true
				}
			}
			if !isIn {
				resultList = append(resultList, avResult)
			}
		}
	}
	return resultList
}
