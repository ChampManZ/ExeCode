package api

import (
	"encoding/json"
	"net/http"

	"github.com/ChampManZ/ExeCode/v2/internal/piston"
	"github.com/labstack/echo/v4"
)

type RuntimeResponse struct {
	Runtimes []piston.Runtime `json:"runtimes"`
}

func RuntimeHandler(c echo.Context) error {
	pistonClient := piston.NewClient(http.DefaultClient, "localhost:2000", "")

	runtimes, err := pistonClient.GetRuntimes()
	if err != nil {
		return err
	}

	b, err := json.Marshal(RuntimeResponse{runtimes})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, string(b))
}
