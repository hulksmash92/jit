package models

import (
	"time"

	"github.com/google/uuid"
)

type ApiClient struct {
	ClientId    uuid.UUID `json:"clientId"`
	SecretKey   string    `json:"secretKey"`
	AppName     string    `json:"appName"`
	Description string    `json:"description"`
	ValidTill   time.Time `json:"validTill"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	UserId      uint      `json:"userId"`
}

type User struct {
	Id            uint           `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Created       time.Time      `json:"created"`
	Updated       time.Time      `json:"updated"`
	GithubUserId  string         `json:"githubUserId"`
	Avatar        string         `json:"avatar"`
	Organisations []Organisation `json:"organisations"`
	ApiClients    []ApiClient    `json:"apiClients"`
}

type Organisation struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Avatar      string    `json:"avatar"`
	Source      string    `json:"source"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type RepoItem struct {
	Id           uint      `json:"id"`
	ItemIdSource string    `json:"itemIdSource"`
	ItemType     string    `json:"itemType"`
	Source       string    `json:"source"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
	RepoName     string    `json:"repoName"`
	Description  string    `json:"description"`
}

type Tag struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type OwnerTrimmed struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type TimeEntry struct {
	Id           uint         `json:"id"`
	User         OwnerTrimmed `json:"user"`
	Organisation OwnerTrimmed `json:"organisation"`
	Comments     string       `json:"comments"`
	Created      time.Time    `json:"created"`
	Updated      time.Time    `json:"updated"`
	Value        float32      `json:"value"`
	ValueType    string       `json:"valueType"`
	Tags         []Tag        `json:"tags,omitempty"`
	RepoItems    []RepoItem   `json:"repoItems,omitempty"`
}
