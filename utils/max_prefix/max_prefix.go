package utils

import "strings"

/* 
	最大前缀匹配算法
*/
func FindMaxPrefix(target string, list []string) (maxPrefix string, index int) {

	// 匹配失败，index == -1
	index = -1
	if len(list) == 0 {
		return
	}
	
	// 从 list 中，找到 target 最匹配的元素
	for i, s := range list {
		if strings.HasPrefix(target, s) && len(s) > len(maxPrefix) {
			maxPrefix = s
			index = i
		}
	}

	return
}
