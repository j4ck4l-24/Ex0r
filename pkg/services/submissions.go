package services

import (
	db "github.com/j4ck4l-24/Ex0r/pkg/database"
	"github.com/j4ck4l-24/Ex0r/pkg/models"
)

func CheckSolved(submission *models.Submission) bool {
	dbConn := db.DB
	query := "SELECT 1 FROM solves WHERE user_id = $1 AND team_id = $2"

	var exists int
	err := dbConn.QueryRow(query, submission.ChallId, submission.TeamId).Scan(&exists)
	return err == nil
}

func CheckForCorrectSubmission(submission *models.Submission) bool {
	var actualFlag string
	dbConn := db.DB
	query := "SELECT content FROM flags WHERE chall_id = $1"

	rows, err := dbConn.Query(query, submission.ChallId)
	if err != nil {
		return false
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&actualFlag)
		if err != nil {
			return false
		}

		if submission.Submitted == actualFlag {
			return true
		}
	}
	return false
}

func UpdateSolves(submission *models.Submission) error {
	query := "INSERT INTO Solves (chall_id, user_id, team_id) VALUES($1, $2, $3) RETURNING id"
	_, err := db.DB.Query(query, submission.ChallId, submission.UserId, submission.TeamId)
	return err
}
