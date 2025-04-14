package services

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/j4ck4l-24/Ex0r/pkg/models"
)

func FetchPublicChallenges(c *fiber.Ctx, db *sql.DB) ([]models.ChallengePublic, error) {
	query := "SELECT id, chall_name, chall_desc, category, current_points, max_attempts, type, author_name, connection_string, created_at, updated_at FROM Challenges WHERE hidden=FALSE"

	m := c.Queries()
	args := []interface{}{}
	idx := 1
	limit := 10
	page := 1
	if category := m["category"]; category != "" {
		query += fmt.Sprintf(" AND category LIKE $%d", idx)
		args = append(args, "%"+category+"%")
		idx++
	}

	if strLimit := m["limit"]; strLimit != "" && strLimit != "null" {
		if parsedLimit, err := strconv.Atoi(strLimit); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if pageStr := m["page"]; pageStr != "" {
		if parsedPage, err := strconv.Atoi(pageStr); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	query += fmt.Sprintf("ORDER BY id LIMIT $%d OFFSET $%d", idx, idx+1)
	args = append(args, limit, (page-1)*limit)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var challenges []models.ChallengePublic
	for rows.Next() {
		var challenge models.ChallengePublic
		err = rows.Scan(&challenge.ChallId, &challenge.ChallName, &challenge.ChallDesc, &challenge.Category, &challenge.Points, &challenge.MaxAttempts, &challenge.Type, &challenge.AuthorName, &challenge.ConnectionString, &challenge.CreatedAt, &challenge.UpdatedAt)
		if err != nil {
			return nil, err
		}
		challenges = append(challenges, challenge)
	}
	return challenges, err
}

func FetchAdminChallenges(c *fiber.Ctx, db *sql.DB) ([]models.ChallengeAdmin, error) {
	query := "SELECT id, chall_name, chall_desc, category, current_points, initial_points, min_points, max_attempts, type, hidden, author_name, COALESCE(decay_type, ''), COALESCE(decay_value, 0), connection_string, created_at, updated_at /*,requirements, next_chall_id*/ FROM Challenges "

	m := c.Queries()
	args := []interface{}{}
	idx := 1
	limit := 10
	page := 1
	if category := m["category"]; category != "" {
		query += fmt.Sprintf(" WHERE category LIKE $%d", idx)
		args = append(args, "%"+category+"%")
		idx++
	}

	if strLimit := m["limit"]; strLimit != "" && strLimit != "null" {
		if parsedLimit, err := strconv.Atoi(strLimit); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if pageStr := m["page"]; pageStr != "" {
		if parsedPage, err := strconv.Atoi(pageStr); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	query += fmt.Sprintf("ORDER BY id LIMIT $%d OFFSET $%d", idx, idx+1)
	args = append(args, limit, (page-1)*limit)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var challenges []models.ChallengeAdmin
	for rows.Next() {
		var challenge models.ChallengeAdmin
		err = rows.Scan(&challenge.ChallId, &challenge.ChallName, &challenge.ChallDesc, &challenge.Category, &challenge.Points, &challenge.MaxPoints, &challenge.MinPoints, &challenge.MaxAttempts, &challenge.Type, &challenge.Hidden, &challenge.AuthorName, &challenge.DecayType, &challenge.DecayValue, &challenge.ConnectionString, &challenge.CreatedAt, &challenge.UpdatedAt)
		if err != nil {
			return nil, err
		}
		challenges = append(challenges, challenge)
	}
	return challenges, err
}

func FetchPublicChallenge(c *fiber.Ctx, db *sql.DB) (models.ChallengePublic, error) {
	chall_id, err := c.ParamsInt("id")
	if err != nil {
		return models.ChallengePublic{}, err
	}
	query := "SELECT id, chall_name, chall_desc, category, current_points, max_attempts, type, author_name, connection_string, created_at, updated_at FROM Challenges WHERE hidden=FALSE AND id = $1"

	rows, err := db.Query(query, chall_id)
	if err != nil {
		return models.ChallengePublic{}, err
	}
	defer rows.Close()
	var challenge models.ChallengePublic
	rows.Next()
	err = rows.Scan(&challenge.ChallId, &challenge.ChallName, &challenge.ChallDesc, &challenge.Category, &challenge.Points, &challenge.MaxAttempts, &challenge.Type, &challenge.AuthorName, &challenge.ConnectionString, &challenge.CreatedAt, &challenge.UpdatedAt)
	if err != nil {
		return models.ChallengePublic{}, err
	}
	return challenge, err
}

func FetchAdminChallenge(c *fiber.Ctx, db *sql.DB) (models.ChallengeAdmin, error) {
	chall_id, err := c.ParamsInt("id")
	if err != nil {
		return models.ChallengeAdmin{}, err
	}
	query := "SELECT id, chall_name, chall_desc, category, current_points, initial_points, min_points, max_attempts, type, hidden, author_name, COALESCE(decay_type, ''), COALESCE(decay_value, 0), connection_string, created_at, updated_at /*,requirements, next_chall_id*/ FROM Challenges WHERE id = $1"

	rows, err := db.Query(query, chall_id)
	if err != nil {
		return models.ChallengeAdmin{}, err
	}
	defer rows.Close()
	var challenge models.ChallengeAdmin
	rows.Next()
	err = rows.Scan(&challenge.ChallId, &challenge.ChallName, &challenge.ChallDesc, &challenge.Category, &challenge.Points, &challenge.MaxPoints, &challenge.MinPoints, &challenge.MaxAttempts, &challenge.Type, &challenge.Hidden, &challenge.AuthorName, &challenge.DecayType, &challenge.DecayValue, &challenge.ConnectionString, &challenge.CreatedAt, &challenge.UpdatedAt)
	if err != nil {
		return models.ChallengeAdmin{}, err
	}
	return challenge, err
}

func UpdateStaticChallenge(dbConn *sql.DB, challenge *models.ChallengeAdmin) (*models.ChallengeAdmin, error) {
	query := "UPDATE Challenges SET chall_name = $1, chall_desc = $2, category = $3, current_points = $4, initial_points = $5, min_points = $6, max_attempts = $7, type = $8, hidden = $9, author_name = $10, connection_string = $11, updated_at = CURRENT_TIMESTAMP WHERE id = $12 RETURNING id;"
	err := dbConn.QueryRow(query, challenge.ChallName, challenge.ChallDesc, challenge.Category, challenge.Points, challenge.MaxPoints, challenge.MinPoints, challenge.MaxAttempts, challenge.Type, challenge.Hidden, challenge.AuthorName, challenge.ConnectionString, challenge.ChallId).Scan(&challenge.ChallId)

	return challenge, err
}

func UpdateDynamicChallenge(dbConn *sql.DB, challenge *models.ChallengeAdmin) (*models.ChallengeAdmin, error) {
	query := "UPDATE Challenges SET chall_name = $1, chall_desc = $2, category = $3, current_points = $4, initial_points = $5, min_points = $6, max_attempts = $7, type = $8, hidden = $9, author_name = $10, decay_type = $11, decay_value = $12, connection_string = $13, updated_at = CURRENT_TIMESTAMP WHERE id = $14 RETURNING id"
	err := dbConn.QueryRow(query, challenge.ChallName, challenge.ChallDesc, challenge.Category, challenge.Points, challenge.MaxPoints, challenge.MinPoints, challenge.MaxAttempts, challenge.Type, challenge.Hidden, challenge.AuthorName, challenge.DecayType, challenge.DecayValue, challenge.ConnectionString, challenge.ChallId).Scan(&challenge.ChallId)

	return challenge, err
}
