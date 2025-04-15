package models

import (
	"time"
)

type PublicChallenge struct {
	ChallId          int       `json:"chall_id"`
	ChallName        string    `json:"chall_name"`
	ChallDesc        string    `json:"chall_desc"`
	Category         string    `json:"category"`
	Points           int       `json:"current_points"`
	SolvedByMe       bool      `json:"solved_by_me"`
	MaxAttempts      int       `json:"max_attempts"`
	Type             string    `json:"type"`
	AuthorName       string    `json:"author_name"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	ConnectionString string    `json:"connection_string"`
}

type AdminChallenge struct {
	ChallId          int       `json:"chall_id"`
	ChallName        string    `json:"chall_name"`
	ChallDesc        string    `json:"chall_desc"`
	Category         string    `json:"category"`
	Points           int       `json:"current_points"`
	SolvedByMe       bool      `json:"solved_by_me"`
	MaxAttempts      int       `json:"max_attempts"`
	Type             string    `json:"type"`
	AuthorName       string    `json:"author_name"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	ConnectionString string    `json:"connection_string"`
	Hidden           bool      `json:"hidden"`
	MaxPoints        int       `json:"max_points"`
	DecayType        string    `json:"decay_type"`
	DecayValue       int       `json:"decay_value"`
	MinPoints        int       `json:"min_points"`
}

type PublicChallengesResp struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []PublicChallenge `json:"data"`
}

type AdminChallengesResp struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    []AdminChallenge `json:"data"`
}

type SinglePublicChallResp struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    PublicChallenge `json:"data"`
}

type SingleAdminChallResp struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    AdminChallenge `json:"data"`
}
