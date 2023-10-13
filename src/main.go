package main

import (
	"fmt"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"net/http"
)

func main() {
	var headerName, expectedValue string

	var rootCmd = &cobra.Command{Use: "header-checker"}
	rootCmd.PersistentFlags().StringVarP(&headerName, "header", "H", "", "Name of the HTTP header to check")
	rootCmd.PersistentFlags().StringVarP(&expectedValue, "value", "V", "", "Expected value of the HTTP header")

	var checkCmd = &cobra.Command{
		Use:   "check",
		Short: "Check the presence of an HTTP header in a request",
		Run: func(cmd *cobra.Command, args []string) {
			if headerName == "" || expectedValue == "" {
				fmt.Println("Header name and expected value are required.")
				return
			}

			r := gin.Default()

			l := logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
				return l.Output(gin.DefaultWriter).With().Logger()
			})
			r.Use(requestid.New())
			r.Use(logger.SetLogger(l))

			r.GET("/", func(c *gin.Context) {
				value := c.GetHeader(headerName)
				if value == expectedValue {
					c.String(http.StatusOK, "HTTP header is present with the expected value.")
				} else {
					c.String(http.StatusNotFound, "HTTP header is not present with the expected value.")
				}
			})

			fmt.Printf("Server is starting on port 8080...\n")
			r.Run(":8080")
		},
	}

	rootCmd.AddCommand(checkCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
