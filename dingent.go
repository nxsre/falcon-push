package main

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/toolkits/net/httplib"
	"go.uber.org/zap"
)

type DingEntMsg struct {
	tos     string `json:"tos"`
	content string `json:"content"`
}

func (this *DingEntMsg) send() error {
	for _, tel := range strings.Split(this.tos, ",") {
		r := httplib.Post(cfg.DingEnt.URL).SetTimeout(5*time.Second, 30*time.Second)
		data := struct {
			Message  string `json:"message"`
			Tel      string `json:"tel"`
			ClientId string `json:"clientId"`
		}{
			Message:  this.content,
			Tel:      tel,
			ClientId: cfg.APP.ClientId,
		}

		datajb, err := json.Marshal(data)
		if err != nil {
			return err
		}

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
		body := "client_id=" + cfg.APP.ClientId + "&currentUserId=0&data=" + string(datajb)
		r.Body(body)
		resp, err := r.String()
		if err != nil {
			logger.Error("send dingent fail", zap.String("type", "dingent"), zap.String("tel", tel), zap.String("content", this.content), zap.Error(err))
			return err
		}
		logger.Debug("send dingent success", zap.String("type", "dingent"), zap.String("tel", tel), zap.String("content", this.content), zap.String("resp", resp))
	}
	return nil
}
