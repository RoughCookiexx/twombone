package main

import (
	"fmt"
	"net/http"
	"strings"
	
	"github.com/RoughCookiexx/gg_elevenlabs"
	"github.com/RoughCookiexx/gg_sse"
	"github.com/RoughCookiexx/twitch_chat_subscriber"
)

func slide(message string)(string) {
	message = afterLastColon(message)
	soundEffect, err := gg_eleven.GenerateSoundEffect(message)
	if err != nil {
		fmt.Println("OH DANG!")
	}
	sse.SendBytes(soundEffect)
	
	return ""
}

func afterLastColon(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 || idx+1 >= len(s) {
		return ""
	}
	return s[idx+1:]
}     

func main() {
	fmt.Println("Subscribing to chat messages")
	targetURL := "http://0.0.0.0:6969/subscribe"
	filterPattern := "custom-reward-id=c235c8fc-5e69-4b20-b7c7-84a54d5c73f6"
	twitch_chat_subscriber.SendRequestWithCallbackAndRegex(targetURL, slide, filterPattern, 6973)
	sse.Start()
	http.ListenAndServe((":6973"), nil)
}
