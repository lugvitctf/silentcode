package main

import (
	"crypto/rand"
	"encoding/hex"
	"sync"

	"github.com/AnimeKaizoku/cacher"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SESSION *gorm.DB

type Session struct {
	SessionId string `gorm:"primary_key"`
	Flag      string
	Password  string
}

type SessionGroup struct {
	SessionId string `gorm:"primary_key"`
	UserId    int64  `gorm:"primary_key"`
}

var sessionCache = cacher.NewCacher[string, *Session](nil)
var userCache = cacher.NewCacher[int64, string](nil)

var mu = sync.Mutex{}

func generateSessionId() string {
	t := make([]byte, 5)
	rand.Read(t)
	sessionId := hex.EncodeToString(t)
	return sessionId
}

func NewSession(flag string) string {
	sessionId := generateSessionId()
	passwd := generateSecretKey()
	encryptedFlag := encryptFlag(flag, passwd)
	w := Session{
		SessionId: sessionId,
		Password:  passwd,
		Flag:      encryptedFlag,
	}
	sessionCache.Set(sessionId, &w)
	tx := SESSION.Begin()
	tx.Create(&w)
	mu.Lock()
	defer mu.Unlock()
	tx.Commit()
	return sessionId
}

func GetSession(sessionId string) *Session {
	session, ok := sessionCache.Get(sessionId)
	if ok {
		return session
	}
	w := Session{SessionId: sessionId}
	SESSION.First(&w)
	if w.Password == "" {
		return nil
	}
	sessionCache.Set(sessionId, &w)
	return &w
}

func AddUserToSession(sessionId string, userId int64) {
	userCache.Set(userId, sessionId)
	w := SessionGroup{SessionId: sessionId, UserId: userId}
	tx := SESSION.Begin()
	tx.Create(&w)
	mu.Lock()
	defer mu.Unlock()
	tx.Commit()
}

func GetUserSession(userId int64) string {
	sessionId, ok := userCache.Get(userId)
	if ok {
		return sessionId
	}
	var w SessionGroup
	SESSION.Where("user_id = ?", userId).First(&w)
	if w.SessionId == "" {
		return ""
	}
	userCache.Set(userId, w.SessionId)
	return w.SessionId
}

func StartDatabase() error {
	var err error
	SESSION, err = gorm.Open(sqlite.Open("user.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return err
	}
	dB, err := SESSION.DB()
	if err != nil {
		return err
	}
	dB.SetMaxOpenConns(150)
	SESSION.AutoMigrate(&Session{}, &SessionGroup{})
	return nil
}
