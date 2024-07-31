package es_service

import (
	"fmt"
	"strings"
)

/*
	# 标题1
	正文1
	## 标题2
	正文2
*/

func PraseMd(title, content string) {

	var (
		body     string
		headList []string
		bodyList []string
		isCode   bool
	)

	headList = append(headList, title)
	list := strings.Split(content, "\n")

	for _, s := range list {
		if strings.HasPrefix(s, "```") {
			// 代码
			isCode = true
		}
		if strings.HasPrefix(s, "#") && !isCode {
			// 标题
			headList = append(headList, getHead(s))
			bodyList = append(bodyList, body)
			body = ""
			continue
		}
		body += s
	}

	bodyList = append(bodyList, body)
	for i, s := range headList {
		fmt.Println("标题", i, s)
	}
	for i, s := range bodyList {
		fmt.Println("正文", i, s)
	}
}

func getHead(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return head
}
