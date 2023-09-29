package cmd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func ServeCMD() *cobra.Command {
	var serveCMD = &cobra.Command{
		Use:   "serve",
		Short: "`serve` este comanda care porneste serverul pentru teza de licenta. Acesta include pornirea serverului pentru afisarea si procesarea datelor.",
		Run: func(cmd *cobra.Command, args []string) {
			e := echo.New()
			e.GET("/", func(c echo.Context) error {
				return c.String(http.StatusOK, "Hello, World!")
			})
		},
	}

	return serveCMD
}
