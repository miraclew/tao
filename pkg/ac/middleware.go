package ac

import (
	"github.com/miraclew/tao/pkg/auth"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func NewMiddleware(authority auth.Verifier, skipPaths []string) echo.MiddlewareFunc {
	isLocal := false
	if os.Getenv("ENV") == "local" {
		isLocal = true
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, s := range skipPaths {
				if s == c.Path() {
					return next(c)
				}
			}

			authorization := getAuthorization(c)
			if authorization == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": "authorization invalid"})
			}

			if isLocal && strings.HasPrefix(authorization, "FT_") {
				userStr := strings.TrimPrefix(authorization, "FT_")
				userID, _ := strconv.ParseInt(userStr, 10, 64)
				identity := &auth.Identity{
					UserID:   userID,
					Roles:    []string{"admin"},
					DeviceID: "DEMO_DEVICE",
				}
				c.Set(UserIdContextKey, identity)
				return next(c)
			}

			identity, expireAt, err := authority.Verify(authorization)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": "authorization invalid"})
			}
			if expireAt < time.Now().Unix() {
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": "authorization expired"})
			}

			c.Set(UserIdContextKey, &Session{
				Identity: identity,
				Authorization: authorization,
			})
			return next(c)
		}
	}
}

func getAuthorization(c echo.Context) string {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		token = c.QueryParam("Authorization")
	}
	if token == "" {
		c, err := c.Cookie("Authorization")
		if err == nil {
			token = c.Value
		}
	}
	return token
}

type Session struct {
	Identity *auth.Identity
	Authorization string
}