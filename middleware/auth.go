package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Authentication(c *fiber.Ctx, session *session.Store) error {

	sess, err := session.Get(c)
	if err != nil {
		return err
	}
	name := sess.Get("name")
	if name != nil {
		c.Locals("name", name)
	}

	return c.Next()

}
