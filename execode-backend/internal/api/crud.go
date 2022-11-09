package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ChampManZ/ExeCode/v2/entities"
	"github.com/labstack/echo/v4"
)

type createUserRequest struct {
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
} // @name CreateUser

// CreateUserHandler godoc
// @Summary     Creates the user defined by the request
// @Description Creates the user defined by the request
// @Tags        Users
// @Accept      application/json
// @Produce     json
// @Param       UserDescription body     createUserRequest true "Description of the user to created"
// @Success     200             {object} api.CreateUserHandler.response "Describes the created user"
// @Router      /users [post]
func CreateUserHandler(c echo.Context) error {
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
	}

	var body createUserRequest
	err = json.Unmarshal(b, &body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
	}

	user, err := entities.CreateUser(body.UserName, body.FirstName, body.LastName, body.Email, body.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
	}

	type response struct {
		Result entities.APIUserBasic `json:"result"`
	} // @name CreateUserResult
	return c.JSON(http.StatusOK, response{user})
}

// GetUsersHandler godoc
// @Summary     Gets all users (paged)
// @Description Gets all users, set page size and page number through query parameters.
// @Description Defaults to page size = 10 and page = 1
// @Tags        Users
// @Param       offset query    int                            false "Page size to return"
// @Param       limit     query    int                            false "Page to return"
// @Success     200      {object} api.GetUsersHandler.response "Describes the result of the execution"
// @Router      /users [get]
func GetUsersHandler(c echo.Context) error {
	pagesizeParam := c.QueryParam("limit")
	pageParam := c.QueryParam("offset")

	var pagesize int
	var page int
	if pagesizeParam == "" {
		pagesize = 10
	} else {
		var err error
		pagesize, err = strconv.Atoi(pagesizeParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid pagesize parameter"})
		}
	}
	if pageParam == "" {
		page = 1
	} else {
		var err error
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid page parameter"})
		}
	}

	users, totalPages, err := entities.GetUsers(pagesize, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}
	type response struct {
		User       []entities.APIUserBasic `json:"users"`
		Page       int                     `json:"page"`
		TotalPages int64                   `json:"total_pages"`
	} // @name UsersList
	resp := response{users, page, totalPages}

	return c.JSON(http.StatusOK, resp)
}

// GetUserHandler godoc
// @Summary     Gets all users (paged)
// @Description Gets all users, set page size and page number through query parameters.
// @Description Defaults to page size = 10 and page = 1
// @Tags        Users
// @Param       username path    string                            false "Username to query"
// @Success     200      {object} api.GetUserHandler.response "Describes the result of the execution"
// @Router      /users/{username} [get]
func GetUserHandler(c echo.Context) error {
	userName := c.Param("username")
	user, err := entities.GetUserByUsername(userName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
	}

	type response struct {
		Result entities.APIUserAdvanced `json:"result"`
	} // @name UserResult
	return c.JSON(http.StatusOK, response{user})
}
