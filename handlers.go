package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func SmsHandle(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	appid := vals.Get("appid")
	tos := vals.Get("tos")
	content := vals.Get("content")
	if appid == "" || tos == "" || content == "" {
		logger.Error("必选参数不能为空", zap.String("appid", appid), zap.String("tos", tos), zap.String("content", content))
		fmt.Fprintln(w, "appid 或 tos 或 content 不能为空")
		return
	}
	sms := &SMSMsg{
		tos:     tos,
		content: content,
	}
	if err := sms.send(); err != nil {
		fmt.Fprintln(w, "send msg failed:"+err.Error())
		return
	}
	fmt.Fprintln(w, "send msg ok")
}

func MailHandle(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	appid := vals.Get("appid")
	tos := vals.Get("tos")
	content := vals.Get("content")
	subject := vals.Get("subject")
	if appid == "" || tos == "" || content == "" || subject == "" {
		logger.Error("必选参数不能为空", zap.String("appid", appid), zap.String("tos", tos), zap.String("content", content), zap.String("subject", subject))
		fmt.Fprintln(w, "appid 或 tos 或 content 不能为空")
		return
	}
	mail := &MailMsg{
		tos:     tos,
		subject: subject,
		content: content,
	}
	if err := mail.send(); err != nil {
		fmt.Fprintln(w, "send mail failed:"+err.Error())
		return
	}
	fmt.Fprintln(w, "send mail ok")
}

func DingTalkHandle(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	appid := vals.Get("appid")
	tos := vals.Get("tos")
	content := vals.Get("content")
	alarmCode := vals.Get("alarmCode")
	if appid == "" || tos == "" || content == "" || alarmCode == "" {
		logger.Error("必选参数不能为空", zap.String("appid", appid), zap.String("tos", tos), zap.String("content", content), zap.String("alarmCode", alarmCode))
		fmt.Fprintln(w, "appid 或 tos 或 content 不能为空")
		return
	}
	ding := &DingTalkMsg{
		alarmCode: alarmCode,
		tos:       tos,
		content:   content,
	}
	if err := ding.send(); err != nil {
		fmt.Fprintln(w, "send ding failed:"+err.Error())
		return
	}
	fmt.Fprintln(w, "send ding ok")
}
