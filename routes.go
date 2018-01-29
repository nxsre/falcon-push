package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{Name: "SmsHandle", Method: "ANY", Pattern: "/api/v1/push/sms", HandlerFunc: SmsHandle},
	Route{Name: "MailHandle", Method: "ANY", Pattern: "/api/v1/push/mail", HandlerFunc: MailHandle},
	Route{Name: "DingTalkHandle", Method: "ANY", Pattern: "/api/v1/push/dingtalk", HandlerFunc: DingTalkHandle},
	Route{Name: "DingEntHandle", Method: "ANY", Pattern: "/api/v1/push/dingent", HandlerFunc: DingEntHandle},
}
