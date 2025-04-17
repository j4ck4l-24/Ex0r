package models

import "time"

type Submission struct {
	SubmissionId int       `json:"submission_id"`
	Submitted    string    `json:"submitted"`
	ChallId      int       `json:"chall_id"`
	UserId       int       `json:"user_id"`
	TeamId       int       `json:"team_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type SubmissionBody struct {
	ChallId  int    `json:"chall_id"`
	Submitted string `json:"submitted"`
}

type SingleSubmissionResp struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    Submission `json:"data"`
}

type MultipleSubmissionResp struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    []Submission `json:"data"`
}
