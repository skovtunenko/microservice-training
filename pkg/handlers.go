package pkg

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//----------
// Handlers
//----------

func GreetingHandler(c echo.Context) error {
	type Greeting struct {
		Msg string `json:"msg"`
	}
	g := &Greeting{
		"Hello, Exadel developers!",
	}
	if err := c.Bind(g); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, g)
}

func CreateUserHandler(c echo.Context) error {
	u := &User{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func GetUserHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func UpdateUserHandler(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func DeleteUserHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
