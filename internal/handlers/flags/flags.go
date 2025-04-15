package flags

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	db "github.com/j4ck4l-24/Ex0r/pkg/database"
	"github.com/j4ck4l-24/Ex0r/pkg/models"
	"github.com/j4ck4l-24/Ex0r/pkg/utils"
)

func CreateFlag(c *fiber.Ctx) error {
	flag := new(models.Flag)
	dbConn := db.DB
	if err := c.BodyParser(flag); err != nil {
		return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid request body")
	}

	query := "INSERT INTO Flags (content, type, chall_id) VALUES ($1, $2, $3) RETURNING id"
	err := dbConn.QueryRow(query, flag.Content, flag.Type, flag.ChallId).Scan(&flag.FlagId)
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusConflict, "This flag already exists for this challenge")
	}
	return utils.SendSingleFlagResp(c, fiber.StatusOK, fmt.Sprintf("Flag with flag_id = %d created successfully for chall_id = %d", flag.FlagId, flag.ChallId), *flag)
}

func GetAllFlag(c *fiber.Ctx) error {
	m := c.Queries()
	args := []interface{}{}
	idx := 1
	dbConn := db.DB
	query := "SELECT id, content, type, chall_id, created_at, updated_at FROM flags "

	if strChall_id := m["chall_id"]; strChall_id != "" && strChall_id != "null" {
		if chall_id, err := strconv.Atoi(strChall_id); err == nil && chall_id > 0 {
			query += fmt.Sprintf(" WHERE chall_id = $%d", idx)
			args = append(args, chall_id)
			idx++
		}
	}

	limit, offset := utils.Pagination(c)

	query += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", idx, idx+1)
	args = append(args, limit, offset)

	rows, err := dbConn.Query(query, args...)

	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusNotFound, "Flags not found")
	}
	defer rows.Close()
	var flags []models.Flag
	for rows.Next() {
		var flag models.Flag
		err = rows.Scan(&flag.FlagId, &flag.Content, &flag.Type, &flag.ChallId, &flag.CreatedAt, &flag.UpdatedAt)
		if err != nil {
			return utils.SendGeneralResp(c, fiber.StatusNotFound, "Flags not found")
		}
		flags = append(flags, flag)
	}

	if len(flags) == 0 {
		return utils.SendGeneralResp(c, fiber.StatusNotFound, "Flags not found")
	}
	return utils.SendMultipleFlagResp(c, fiber.StatusOK, "Flags fetched Successfully", flags)

}

func GetFlag(c *fiber.Ctx) error {
	dbConn := db.DB
	flag_id, err := c.ParamsInt("id")
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusNotFound, "Invalid flag_id")
	}
	query := "SELECT id, content, type, chall_id, created_at, updated_at FROM flags WHERE id = $1"

	rows, err := dbConn.Query(query, flag_id)
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusNotFound, "Flag not found")
	}
	defer rows.Close()
	var flag models.Flag
	rows.Next()
	err = rows.Scan(&flag.FlagId, &flag.Content, &flag.Type, &flag.ChallId, &flag.CreatedAt, &flag.UpdatedAt)
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusNotFound, "Flag not found")
	}
	return utils.SendSingleFlagResp(c, fiber.StatusOK, "Flag fetched Successfully", flag)
}

func UpdateFlag(c *fiber.Ctx) error {
	dbConn := db.DB
	flag := new(models.Flag)
	err := c.BodyParser(flag)
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid request body")
	}

	flag.FlagId, err = c.ParamsInt("id")
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusNotFound, "Invalid chall_id")
	}

	query := "UPDATE Flags SET content = $1, type = $2, chall_id = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4 RETURNING id;"
	err = dbConn.QueryRow(query, flag.Content, flag.Type, flag.ChallId, flag.FlagId).Scan(&flag.FlagId)
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusInternalServerError, "Failed to update flag either there is an error or flag with flag_id doens't exist")
	}
	return utils.SendGeneralResp(c, fiber.StatusOK, fmt.Sprintf("Flag updated successfully with id = %d", flag.FlagId))}


func DeleteFlag(c *fiber.Ctx) error {
	dbConn := db.DB
	flag_id, _ := c.ParamsInt("id")

	query := "DELETE FROM Flags WHERE id = $1 RETURNING id"

	err := dbConn.QueryRow(query, flag_id).Scan(&flag_id)

	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusOK, "Failed to delete flag")
	}
	return utils.SendGeneralResp(c, fiber.StatusOK, fmt.Sprintf("Flag deleted successfully with chall_id = %d", flag_id))
}
