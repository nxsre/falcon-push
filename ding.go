package main

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/toolkits/net/httplib"
	"go.uber.org/zap"
)

type DingTalkMsg struct {
	alarmCode string `json:"alarmCode"`
	tos       string `json:"tos"`
	content   string `json:"content"`
}

func (this *DingTalkMsg) send() error {
	r := httplib.Post(cfg.DingTalk.URL).SetTimeout(5*time.Second, 30*time.Second)
	data := struct {
		AlarmCode string   `json:"alarmCode"`
		Content   string   `json:"content"`
		AT        []string `json:"_at"`
	}{
		AlarmCode: this.alarmCode,
		Content:   this.content,
		AT:        strings.Split(this.tos, ","),
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
		logger.Error("send dingtalk fail", zap.String("alarmCode", this.alarmCode), zap.String("tos", this.tos), zap.String("content", this.content), zap.Error(err))
		return err
	}
	logger.Debug("send dingtalk success", zap.String("alarmCode", this.alarmCode), zap.String("tos", this.tos), zap.String("content", this.content), zap.String("resp", resp))
	return nil
}
