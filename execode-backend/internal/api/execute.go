package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/ChampManZ/ExeCode/v2/internal/piston"
	"github.com/labstack/echo/v4"
)

// RuntimeHandler godoc
// @Summary     List installed Piston runtimes.
// @Description Get all available language runtimes provided by the connected piston instance
// @Tags        Execute
// @Accept      */*
// @Produce     json
// @Success     200 {object} RuntimeResponse
// @Router      /execute/runtimes [get]
func RuntimeHandler(c echo.Context) error {
	pistonClient := piston.NewClient(http.DefaultClient, "localhost:2000", "")

	runtimes, err := pistonClient.GetRuntimes()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, RuntimeResponse{runtimes})
}

type RuntimeResponse struct {
	Runtimes []piston.Runtime `json:"runtimes"`
}

// ExecuteHandler godoc
// @Summary     Execute code and handle user submissions.
// @Description Execute the code in the language specified in the POST request body.
// @Description If it is a submission task then save user submissions as well.
// @Tags        Execute
// @Accept      application/json,text/xml
// @Produce     json
// @Param       JobDescription body     ExecuteRequest  true "Description of the job to be run"
// @Success     200            {object} ExecuteResponse "Describes the result of the execution"
// @Router      /execute [post]
func ExecuteHandler(c echo.Context) error {
	pistonClient := piston.NewClient(http.DefaultClient, "localhost:2000", "")

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	executionRequest := new(ExecuteRequest)
	if err = json.Unmarshal(b, &executionRequest); err != nil {
		return err
	}

	respBody := new(ExecuteResponse)
	respBody.Run = make([]piston.JobOutput, len(executionRequest.Inputs))
	// result := new(piston.ExecutionResult)

	var wg sync.WaitGroup
	doRequest := func(task *piston.ExecutionTask, i int) error {
		defer wg.Done()
		result, err := pistonClient.Execute(task)
		if err != nil {
			return err
		}
		respBody.Run[i] = result.Run
		return nil
	}

	for i, input := range executionRequest.Inputs {
		task := &piston.ExecutionTask{
			Language:           executionRequest.Language,
			Version:            executionRequest.Version,
			Files:              []piston.JobFile{{Name: executionRequest.Name, Content: executionRequest.Content}},
			Stdin:              input.Stdin,
			Args:               input.Args,
			RunTimeout:         0,
			CompileTimeout:     0,
			RunMemoryLimit:     0,
			CompileMemoryLimit: 0,
		}

		wg.Add(1)
		go doRequest(task, i)
	}

	wg.Wait()

	return c.JSON(http.StatusOK, respBody)
}

type ExecuteRequest struct {
	Language string  `json:"language"`
	Version  string  `json:"version"`
	Name     string  `json:"name"`
	Content  string  `json:"content"`
	Inputs   []Input `json:"inputs"`
}

type Input struct {
	Stdin string   `json:"stdin"`
	Args  []string `json:"args"`
}

type ExecuteResponse struct {
	Run []piston.JobOutput `json:"run"`
	// Language string             `json:"language"`
	// Version  string             `json:"version"`
}
