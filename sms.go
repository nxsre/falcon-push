package main

import (
	"strings"
	"time"

	"github.com/toolkits/net/httplib"
	"go.uber.org/zap"
)

type SMSMsg struct {
	tos     string `json:"tos"`
	content string `json:"content"`
}

func (this *SMSMsg) send() error {
	for _, to := range strings.Split(this.tos, ",") {
		r := httplib.Post(cfg.SMS.URL).SetTimeout(5*time.Second, 30*time.Second)

		// clientId=60004&context=$$SMS_CONTENT$$&tel=$$MOBILE_NUM$$
		r.Header("X-Requested-With", "XMLHttpRequest")
		r.Header("Accept-Encoding", "identity")
		r.Header("Content-Type", "application/x-www-form-urlencoded")
		/* 	r.SetProxy(
			func(req *http.Request) (*url.URL, error) {
				u, _ := url.ParseRequestURI("http://127.0.0.1:44444")
				return u, nil
			},
		) */
		// r.Header("Accept-Encoding", "")
		body := "clientId=" + cfg.APP.ClientId + "&currentUserId=0&context=" + this.content + "&tel=" + to
		r.Body(body)
		resp, err := r.String()
		if err != nil {
			logger.Error("send sms fail", zap.String("type", "sms"), zap.String("to", to), zap.String("content", this.content), zap.Error(err))
		} else {
			logger.Debug("send sms success", zap.String("type", "sms"), zap.String("to", to), zap.String("content", this.content), zap.String("resp", resp))
		}
	}

	return nil
}
