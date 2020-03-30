package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	/*
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
	*/

	//e.PUT("/nsmsf-sms/v2/ue-contexts/:supi", handleSubsReg)
	e.GET("/nsmsf-sms/v2/ue-contexts/:supi", handleSubsReg)

	e.Logger.Fatal(e.Start(":1323"))
}

func handleSubsReg(c echo.Context) error {

	authHeaderValue := c.Request().Header.Get("Authorization")
	fmt.Printf("authHeaderValue: %#v\n", authHeaderValue)

	tmpTokens := strings.Split(authHeaderValue, "Bearer ")
	fmt.Printf("len(tmpTokens): %#v\n", len(tmpTokens))
	if len(tmpTokens) < 2 {
		c.Logger().Error("Invalid auth header value")
		return c.String(http.StatusBadRequest, "Invalid auth header value\n")
	}

	accessToken := tmpTokens[1]
	fmt.Printf("accessToken: %#v\n", accessToken)

	supi := c.Param("supi")
	return c.String(http.StatusOK, supi+"\n")
}

/*
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email+"\n")
}

func save2(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}
*/
