package models

import "time"

type Flag struct {
	FlagId    int       `json:"flag_id"`
	Content   string    `json:"content"`
	Type      string    `json:"type"`
	ChallId   int       `json:"chall_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SingleFlagResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Flag   `json:"data"`
}

type MultipleFlagResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Flag `json:"data"`
}
