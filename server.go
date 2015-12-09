package main

import (
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/gin-gonic/contrib/static"
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"os"
	"io/ioutil"
)

func ServerCommand() cli.Command {
	return cli.Command{
		Name:        "server",
		Usage:       "run web app",
		Description: "Use enviroment variable PORT to set which port to listen to.",
		ArgsUsage:   "",
		Action:      ServerAction,
		Flags:       []cli.Flag{},
	}
}

func ServerAction(c *cli.Context) {
	r := gin.Default()

	r.POST("/convert", func(c *gin.Context) {

		file, _, err := c.Request.FormFile("messages")
		defer file.Close()
		check(err)

		b, err := ioutil.ReadAll(file)
		check(err)
		fbData := FromHTML(string(b))

		c.Header("Content-disposition", "attachment; filename=messages.json");
		c.Header("Content-type", "application/json");

		c.JSON(200, fbData)
	})

	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	r.Run(GetAddr())
}

func GetAddr() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	host := os.Getenv("HOST")

	return host + ":" + port
}