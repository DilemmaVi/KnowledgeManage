package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"sync/atomic"
	"time"
)

var loggerQueue = &logQueue{channel: make(chan *Logger, 100), isRuning: 0}

type logQueue struct {
	channel  chan *Logger
	isRuning int32
}

// Logger struct .
type Logger struct {
	LoggerId int64 `json:"log_id"`
	MemberId int   `json:"member_id"`
	// 日志类别：operate 操作日志/ system 系统日志/ exception 异常日志 / document 文档操作日志
	Category     string    `json:"category"`
	Content      string    `json:"content"`
	OriginalData string    `json:"original_data"`
	PresentData  string    `json:"present_data"`
	CreateTime   time.Time `json:"create_time"`
	UserAgent    string    `json:"user_agent"`
	IPAddress    string    `json:"ip_address"`
}

// TableName 获取对应数据库表名.
func (m *Logger) TableName() string {
	return "logs"
}

func NewLogger() *Logger {
	return &Logger{}
}

func (m *Logger) Add() error {
	if m.MemberId <= 0 {
		return errors.New("用户ID不能为空")
	}
	if m.Category == "" {
		m.Category = "system"
	}
	if m.Content == "" {
		return errors.New("日志内容不能为空")
	}
	loggerQueue.channel <- m
	if atomic.LoadInt32(&(loggerQueue.isRuning)) <= 0 {
		atomic.AddInt32(&(loggerQueue.isRuning), 1)
		go addLoggerAsync()
	}
	return nil
}

func addLoggerAsync() {
	defer atomic.AddInt32(&(loggerQueue.isRuning), -1)
	o := orm.NewOrm()

	for {
		logger := <-loggerQueue.channel

		o.Insert(logger)
	}
}
