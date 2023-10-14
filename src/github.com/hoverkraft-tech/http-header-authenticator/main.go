package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
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

			l := zerolog.New(os.Stdout).With().Timestamp().Logger()
			l.Info().Msg(fmt.Sprintf("Expected header is '%s' and expected value is '%s'", headerName, expectedValue))

			r := gin.New()
			r.Use(gin.Recovery())
			r.Use(requestid.New())
			r.Use(CustomLoggerMiddleware(l))

			// Health check endpoint
			r.GET("/health", func(c *gin.Context) {
				c.String(http.StatusOK, "OK")
			})

			// Check the presence of the header in all requests
			r.NoRoute(func(c *gin.Context) {
				value := c.GetHeader(headerName)
				if value == expectedValue {
					c.JSON(http.StatusOK, "HTTP header is present with the expected value.")
				} else {
					c.JSON(http.StatusForbidden, "HTTP header is not present with the expected value.")
				}
			})

			fmt.Printf("Server is starting on port 8080...\n")
			if err := r.Run(":8080"); err != nil {
				fmt.Printf("Server failed to start: %v\n", err)
			}
		},
	}

	rootCmd.AddCommand(checkCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func CustomLoggerMiddleware(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Continue processing other middleware and the request
		c.Next()

		// Check the request path and exclude the health route from logging
		if c.Request.URL.Path != "/health" {
			logger.Info().
				Str("request_id", c.Request.Header.Get("X-Request-ID")).
				Str("client_ip", strings.Split(c.Request.RemoteAddr, ":")[0]).
				Int("status", c.Writer.Status()).
				Str("method", c.Request.Method).
				Str("path", c.Request.URL.Path).
				Str("protocol", c.Request.Proto).
				Str("user_agent", c.Request.UserAgent()).
				Msg("")
		}
	}
}
