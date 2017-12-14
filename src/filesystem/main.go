package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// goflags "github.com/jessevdk/go-flags"
)

// var flags struct {
// 	FilePath string `long:"file-path" description:"static file path" required:"true" default:""`
// }

func main() {

	// parser := goflags.NewParser(&flags, goflags.HelpFlag|goflags.PassDoubleDash|goflags.IgnoreUnknown)
	// _, err = parser.ParseArgs(os.Args[1:])
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	exitutils.Failure()
	// }

	router := gin.Default()
	router.Static("/assets", "/Users/chenyajun/Documents/goproject/src/filesystem/test")
	router.StaticFS("/StaticFS", http.Dir("./test"))
	// router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
