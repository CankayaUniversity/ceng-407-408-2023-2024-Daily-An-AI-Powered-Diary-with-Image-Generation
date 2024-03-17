package health

import (
	"context"
	"net/http"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	var html string
	if dbHealth() {
		html = `
		<!DOCTYPE html>
		<html>
		<head>
		<title>Page Title</title>
		</head>
		<body>
			<h1>Database is up and running...<h1>
		</body>
		</html>
		`
	} else {
		html = `
		<!DOCTYPE html>
		<html>
		<head>
		<title>Page Title</title>
		</head>
		<body>
			<h1>Database connection could not be established...<h1>
		</body>
		</html>
		`
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

func dbHealth() bool {
	// Put your mongodb ping here, for instance
	err := database.GetClient().Ping(context.TODO(), nil)
	return err == nil
}
