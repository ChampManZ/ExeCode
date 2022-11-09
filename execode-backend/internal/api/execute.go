package api

import (
	"encoding/json"
	"fmt"
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
func (env Env) RuntimeHandler(c echo.Context) error {
	pistonClient := piston.NewClient(
		http.DefaultClient,
		fmt.Sprintf("%v:%d", env.PistonHost, env.PistonPort),
		env.PistonAPIKey,
	)

	runtimes, statusCode, err := pistonClient.GetRuntimes()
	if err != nil {
		return HandleErrorFromPiston(c, statusCode, err)
	}

	return c.JSON(http.StatusOK, RuntimeResponse{runtimes})
}

type RuntimeResponse struct {
	Runtimes []piston.Runtime `json:"runtimes"`
} // @name PistonRuntimeList

// ExecuteHandler godoc
// @Summary     Execute code and handle user submissions.
// @Description Execute the code in the language specified in the POST request body.
// @Description If it is a submission task then save user submissions as well.
// @Tags        Execute
// @Accept      application/json,text/xml
// @Produce     json
// @Param       JobDescription body     api.ExecuteHandler.request  true "Description of the job to be run"
// @Success     200            {object} api.ExecuteHandler.response "Describes the result of the execution"
// @Router      /execute [post]
func (env Env) ExecuteHandler(c echo.Context) error {
	pistonClient := piston.NewClient(
		http.DefaultClient,
		fmt.Sprintf("%v:%d", env.PistonHost, env.PistonPort),
		env.PistonAPIKey,
	)

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	type request struct {
		Language string  `json:"language"`
		Version  string  `json:"version"`
		Name     string  `json:"name"`
		Content  string  `json:"content"`
		Inputs   []input `json:"inputs"`
	} // @name ExecutionRequest
	executionRequest := new(request)
	if err = json.Unmarshal(b, &executionRequest); err != nil {
		return err
	}

	type response struct {
		Run    []piston.JobOutput             `json:"run"`
		Errors map[string]pistonErrorResponse `json:"errors,omitempty"`
		// Language string             `json:"language"`
		// Version  string             `json:"version"`
	} // @name ExecutionResults
	respBody := new(response)
	respBody.Run = make([]piston.JobOutput, len(executionRequest.Inputs))
	respBody.Errors = make(map[string]pistonErrorResponse)
	// result := new(piston.ExecutionResult)

	var wg sync.WaitGroup
	doRequest := func(task *piston.ExecutionTask, i int) {
		defer wg.Done()
		result, statusCode, err := pistonClient.Execute(task)
		if err != nil {
			if statusCode == -1 {
				statusCode = http.StatusInternalServerError
			}
			respBody.Errors[fmt.Sprint(i)] = pistonErrorResponse{statusCode, err.Error()}
			return
		}
		respBody.Run[i] = result.Run
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

type input struct {
	Stdin string   `json:"stdin"`
	Args  []string `json:"args"`
} // @name ExecutionInput
type pistonErrorResponse struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"message"`
} // @name PistonError
