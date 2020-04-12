package src

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Test struct {
	Hello string `json:"hello"`
	World string `json:"world"`
}

type User struct {
	Id        int    `json:"_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func BindRouterWithApp(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API 3H-Shop.")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "API 3H-Shop:pong")
	})

	router.GET("/test", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(func(c *gin.Context) (interface{}, error) {
			return Test{
				Hello: "ok",
				World: "ok",
			}, nil
		}).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/error", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(func(c *gin.Context) (interface{}, error) {
			return nil, errors.New("test error")
		}).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/users", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(func(c *gin.Context) (interface{}, error) {
			connection := connections.Mysql.GetConnection()

			stmt, err := connection.Prepare("SELECT _id, email, name, password, address, status, created_at, updated_at FROM `users`")
			if err != nil {
				return nil, err
			}

			rows, err := stmt.Query()

			if err != nil {
				return nil, err
			}

			defer rows.Close()
			var listUser []*User

			for rows.Next() {
				_user := User{}

				err = rows.Scan(&_user.Id, &_user.Email, &_user.Name, &_user.Password, &_user.Address, &_user.Status, &_user.CreatedAt, &_user.UpdatedAt)

				if err != nil {
					return nil, err
				}

				listUser = append(listUser, &_user)
			}

			if err = rows.Err(); err != nil {
				return nil, err
			}

			return listUser, nil
		}).Then(response.SendSuccess).Catch(response.SendError)
	})
}
