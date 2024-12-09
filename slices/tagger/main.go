package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	tag := []sms{}
	for _, msg := range messages {
		NewMsg := sms{id: msg.id, content: msg.content, tags: tagger(msg)}
		tag = append(tag, NewMsg)
	}
	return tag
}

func tagger(msg sms) []string {
	tags := []string{}
	
		cont := strings.ToLower(msg.content)
		if strings.Contains(cont, "urgent") {
			tags = append(tags, "Urgent")
		}
		if strings.Contains(cont, "sale") {
			tags = append(tags, "Promo")
		}
		
		return tags
	
}
