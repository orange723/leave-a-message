package message

import (
	"strconv"
	"time"

	"leave-a-message/database"
	"leave-a-message/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Message struct {
	database.DefaultModel
	Message string `json:"message"`
}

func GetMessages(c *fiber.Ctx) error {
	db := database.DB
	var messages []Message

	// 从查询参数中获取 limit，默认为 10
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}
	// 设置一个最大 limit，防止滥用
	if limit > 100 {
		limit = 100
	}

	// 从查询参数中获取 page，默认为 1
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	// 计算正确的 offset
	offset := (page - 1) * limit

	result := db.Raw("SELECT * FROM messages ORDER BY created_at DESC LIMIT ? OFFSET ?", limit, offset).Scan(&messages)
	if result.Error != nil {
		log.Error(result.Error)

		return pkg.Unexpected("Failed to get messages")
	}

	return c.JSON(messages)
}

func NewMessage(c *fiber.Ctx) error {
	db := database.DB
	message := new(Message)

	if err := c.BodyParser(message); err != nil {
		return pkg.BadRequest("Invalid params")
	}

	result := db.Exec("INSERT INTO messages (created_at, updated_at, message) VALUES (?, ?, ?)", time.Now(), time.Now(), message.Message)
	if result.Error != nil {
		log.Error(result.Error)

		return pkg.Unexpected("Failed to create message")
	}

	return c.SendStatus(200)
}
