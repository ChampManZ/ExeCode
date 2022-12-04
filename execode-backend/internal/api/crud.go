package api

import (
	"net/http"
	"strconv"

	"github.com/ChampManZ/ExeCode/v2/entities"
	"github.com/labstack/echo/v4"
)

/* --------------USERS-------------- */

// CreateUserHandler godoc
// @Summary     Creates the user defined by the request
// @Description Creates the user defined by the request
// @Tags        Users
// @Accept      application/json
// @Produce     json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param       UserDescription body     api.CreateUserHandler.request  true "Description of the user to created"
// @Success     200             {object} api.CreateUserHandler.response "Describes the created user"
// @Router      /users [post]
func CreateUserHandler(c echo.Context) error {
	type request struct {
		UserName string `json:"user_name"`
		Name     string `json:"name"`
		Email    string `json:"email"`
	} // @name CreateUser
	var body request
	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
	}

	user, err := entities.CreateUser(body.UserName, body.Name, body.Email)
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
// @Tags        Users
// @Param       offset query    int                          false "Page size to return" default(1)
// @Param       limit  query    int                          false "Page to return" default(10)
// @Success     200    {object} api.GetUsersHandler.response "Describes the result of the execution"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /users [get]
func GetUsersHandler(c echo.Context) error {
	pagesizeParam := c.QueryParam("limit")
	pageParam := c.QueryParam("offset")

	page, pagesize, err := parseOffsetLimit(pageParam, pagesizeParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid page or pagesize parameter"})
	}
	users, count, err := entities.GetUsers(pagesize, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}
	type response struct {
		User  []entities.APIUserBasic `json:"users"`
		Count int64                   `json:"count"`
		// Page       int                     `json:"page"`
		// TotalPages int64                   `json:"total_pages"`
	} // @name UsersList
	// resp := response{users, page, totalPages}
	resp := response{users, count}

	return c.JSON(http.StatusOK, resp)
}

// GetUserHandler godoc
// @Summary     Gets single user defined by username parameter
// @Description Queries only one user resulting from the specified username
// @Tags        Users
// @Param       userID path     uint                      false "Username to query"
// @Success     200      {object} swaggercompat.Response{result=swaggercompat.UserAdvanceWithRelation} "Describes the user entity"
// @Router      /users/{userID} [get]
func GetUserHandler(c echo.Context) error {
	userID := c.Param("userID")
	uid, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"invalid class ID"})
	}
	user, err := entities.GetUserByUserID(uint(uid))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
	}

	type response struct {
		Result entities.APIUserAdvanced `json:"result"`
	} // @name UserResult
	return c.JSON(http.StatusOK, response{user.Advanced()})
}

/* --------------CLASSES-------------- */

// CreateClassHandler godoc
// @Summary     Creates the class defined by the request
// @Description Creates the class defined by the request, the classes will be lectured by the specified lecturers
// @Tags        Classes
// @Accept      application/json
// @Produce     json
// @Param       ClassDescription body     api.CreateClassHandler.request  true "Description of the class to created"
// @Success     200             {object} api.CreateClassHandler.response "Describes the created class"
// @Router      /classes [post]
func CreateClassHandler(c echo.Context) error {
	type request struct {
		ClassName        string   `json:"class_name"`
		ClassDescription string   `json:"class_description"`
		Lecturers        []string `json:"lecturers"`
	}
	body := request{}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
	}

	if len(body.Lecturers) < 1 || body.ClassName == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"missing required fields"})
	}

	lecturers := make([]entities.User, len(body.Lecturers))
	if err := entities.QueryUsersByUsername(body.Lecturers, &lecturers); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	created := entities.Class{
		ClassName:        body.ClassName,
		ClassDescription: body.ClassDescription,
		User:             lecturers,
	}
	if err := entities.Create(&created); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	type response struct {
		Result entities.APIClassBasic `json:"result"`
	}
	return c.JSON(http.StatusOK, response{
		entities.APIClassBasic{
			ID:               created.ID,
			ClassName:        created.ClassName,
			ClassDescription: created.ClassDescription,
		},
	})
}

// GetClassHandler godoc
// @Summary     Gets all classes (paged)
// @Description Gets all classes, set page size and page number through query parameters.
// @Tags        Classes
// @Param       classID  path    uint                          false "ID of class to return"
// @Success     200    {object} api.GetClassesHandler.response "Describes the result of the execution"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /classes/{classID} [get]
func GetClassHandler(c echo.Context) error {
	classID := c.Param("classID")
	CID, err := strconv.Atoi(classID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"invalid class ID"})
	}
	class, err := entities.GetClassByID(uint(CID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}
	ret := class.Advanced()

	type response struct {
		Result entities.APIClassAdvanced
	}
	return c.JSON(http.StatusOK, response{ret})
}

// GetClassesHandler godoc
// @Summary     Gets all classes (paged)
// @Description Gets all classes, set page size and page number through query parameters.
// @Tags        Classes
// @Param       offset query    int                          false "Page size to return" default(1)
// @Param       limit  query    int                          false "Page to return" default(10)
// @Success     200    {object} api.GetClassesHandler.response "Describes the result of the execution"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /classes [get]
func GetClassesHandler(c echo.Context) error {
	pagesizeParam := c.QueryParam("limit")
	pageParam := c.QueryParam("offset")

	page, pagesize, err := parseOffsetLimit(pageParam, pagesizeParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid page or pagesize parameter"})
	}

	classes := []entities.APIClassBasic{}
	count, err := entities.QueryMany(&entities.Class{}, &classes, pagesize, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	type response struct {
		User  []entities.APIClassBasic `json:"classes"`
		Count int64                    `json:"count"`
	}

	return c.JSON(http.StatusOK, response{classes, count})
}

// DeleteClassHandler godoc
// @Summary     Get user class specified by class_id
// @Tags        Classes
// @Param       userID path     string  true "username of user to query"
// @Success     200    {object} api.GetUserClassesHandler.response "Describes the result of the execution"
// @Failure		500 {object} ErrorResponse
// @Router      /users/{userID}/classes [get]
func GetUserClassesHandler(c echo.Context) error {
	username := c.Param("username")

	class, err := entities.GetUserClasses(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	type response struct {
		Classes []entities.APIClassBasic `json:"classes"`
	}
	return c.JSON(http.StatusOK, response{class.Basic()})
}

// DeleteClassHandler godoc
// @Summary     Delete class specified by class_id
// @Tags        Classes
// @Param       ClassID path     uint  true "ID of class to delete"
// @Success     200    {object} api.DeleteClassHandler.response "Describes the result of the execution"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /classes/{classID} [delete]
func DeleteClassHandler(c echo.Context) error {
	classID := c.Param("classID")
	CID, err := strconv.Atoi(classID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"invalid class ID"})
	}

	if err := entities.DeleteClass(uint(CID)); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})

	}
	type response struct {
		Status string `json:"status"`
	}
	return c.JSON(http.StatusOK, response{"success"})
}

/* --------------LECTURES-------------- */

// CreateLectureHandler godoc
// @Summary     Create lecture for class specified by class_id
// @Tags        Lectures
// @Param       LectureDescription body     api.CreateLectureHandler.request  true "Describes lecture to be created"
// @Success     200    {object} api.CreateLectureHandler.response "Describes lecture created"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /lectures [post]
func CreateLectureHandler(c echo.Context) error {
	type request struct {
		ClassID            int    `json:"class_id"`
		LectureName        string `json:"lecture_name"`
		LectureDescription string `json:"lecture_description"`
		Content            string `json:"lecture_content"`
	}
	body := request{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
	}

	lecture := entities.Lecture{
		ClassID:            body.ClassID,
		LectureName:        body.LectureName,
		LectureDescription: body.LectureDescription,
		LectureContent:     entities.LectureContent{Content: body.Content},
	}

	if err := entities.Create(&lecture); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	type response struct {
		Result entities.APILectureAdvanced `json:"result"`
	}
	return c.JSON(http.StatusOK, response{lecture.Advanced()})
}

// GetLectureHandler godoc
// @Summary     Get Lecture specified by lecture_id
// @Tags        Lectures
// @Param       LectureID path     uint  true "LectureID to get"
// @Success     200    {object} api.GetLectureHandler.response "Describes lecture created"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /lectures/{LectureID} [get]
func GetLectureHandler(c echo.Context) error {
	str_id := c.Param("lectureID")
	ID, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"invalid lecture ID"})
	}

	lecture, err := entities.GetLectureByID(uint(ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}
	type response struct {
		Result entities.APILectureAdvanced `json:"result"`
	}
	return c.JSON(http.StatusOK, response{lecture.Advanced()})
}

// GetClassLecturesHandler godoc
// @Summary     Get lectures belonging to class
// @Tags        Lectures
// @Param       ClassID path     uint  true "Class ID to query"
// @Success     200    {object} api.GetClassLecturesHandler.response "Describes lecture created"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /classes/{ClassID}/lectures [get]
func GetClassLecturesHandler(c echo.Context) error {
	class := c.Param("class")
	cid, err := strconv.ParseUint(class, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"invalid class ID"})
	}

	lectures, err := entities.GetClassLecturesByID(uint(cid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	type response struct {
		Lectures []entities.APILectureBasic `json:"lectures"`
	}
	return c.JSON(http.StatusOK, response{entities.LectureList(lectures).Basic()})
}

// DeleteLectureHandler godoc
// @Summary     Get Lecture specified by lecture_id
// @Tags        Lectures
// @Param       LectureID path     uint  true "LectureID to get"
// @Success     200    {object} api.GetLectureHandler.response "Describes lecture created"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /lectures/{LectureID} [delete]
func DeleteLectureHandler(c echo.Context) error {
	str_id := c.Param("lectureID")
	ID, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"invalid lecture ID"})
	}

	if err = entities.DeleteLecture(uint(ID)); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	type response struct {
		Status string `json:"status"`
	}
	return c.JSON(http.StatusOK, response{"success"})
}

// CreateProblemHandler godoc
// @Summary     Create lecture for class specified by class_id
// @Tags        Problems
// @Param       ProblemDescription body     api.CreateProblemHandler.request  true "Describes lecture to be created"
// @Success     200    {object} api.CreateProblemHandler.response "Describes lecture created"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /problems [post]
func CreateProblemHandler(c echo.Context) error {
	type testcase struct {
		Input  string `json:"input"`
		Output string `json:"output"`
	}
	type request struct {
		ClassID     int                 `json:"class_id"`
		ProblemName string              `json:"problem_name"`
		Content     string              `json:"content"`
		Testcases   []entities.TestCase `json:"testcases"`
	}
	body := request{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
	}

	problem := entities.Problem{
		ClassID:     body.ClassID,
		ProblemName: body.ProblemName,
		ProblemContent: entities.ProblemContent{
			Content:   body.Content,
			TestCases: body.Testcases,
		},
	}

	if err := entities.Create(&problem); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	type response struct {
		Result entities.Problem `json:"result"`
	}
	return c.JSON(http.StatusOK, response{problem})
}

// GetProblemsHandler godoc
// @Summary     Gets all classes (paged)
// @Description Gets all classes, set page size and page number through query parameters.
// @Tags        Problems
// @Param       offset query    int                          false "Page size to return" default(1)
// @Param       limit  query    int                          false "Page to return" default(10)
// @Success     200    {object} api.GetProblemsHandler.response "Describes the result of the execution"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /problems [get]
func GetProblemsHandler(c echo.Context) error {
	pagesizeParam := c.QueryParam("limit")
	pageParam := c.QueryParam("offset")

	page, pagesize, err := parseOffsetLimit(pageParam, pagesizeParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid page or pagesize parameter"})
	}
	problems := []entities.APIProblemBasic{}
	count, err := entities.QueryMany(&entities.Problem{}, &problems, pagesize, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	type response struct {
		User  []entities.APIProblemBasic `json:"classes"`
		Count int64                      `json:"count"`
	}

	return c.JSON(http.StatusOK, response{problems, count})
}

// GetProblemHandler godoc
// @Summary     Get Lecture specified by lecture_id
// @Tags        Problems
// @Param       ProblemID path     uint  true "LectureID to get"
// @Success     200    {object} api.GetLectureHandler.response "Describes lecture created"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /problems/{ProblemID} [get]
func GetProblemHandler(c echo.Context) error {
	str_id := c.Param("problemID")
	ID, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"invalid problem ID"})
	}

	problem, err := entities.GetProblemByID(uint(ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}
	type response struct {
		Result entities.APIProblemAdvanced `json:"result"`
	}
	return c.JSON(http.StatusOK, response{problem.Advanced()})
}

// GetClassLecturesHandler godoc
// @Summary     Get lectures belonging to class
// @Tags        Lectures
// @Param       ClassID path     uint  true "Class ID to query"
// @Success     200    {object} api.GetClassProblemsHandler.response "Describes lecture created"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /classes/{ClassID}/problems [get]
func GetClassProblemsHandler(c echo.Context) error {
	class := c.Param("class")
	cid, err := strconv.ParseUint(class, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"invalid class ID"})
	}

	problems, err := entities.GetClassProblemsByID(uint(cid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	type response struct {
		Problems []entities.APIProblemBasic `json:"problems"`
	}

	return c.JSON(http.StatusOK, response{entities.ProblemList(problems).Basic()})
}

// DeleteProblemHandler godoc
// @Summary     Get Lecture specified by problem_id
// @Tags        Problems
// @Param       ProblemID path     uint  true "ProblemID to delete"
// @Success     200    {object} api.DeleteProblemHandler.response "Describes lecture created"
// @Failure		400 {object} ErrorResponse
// @Failure		500 {object} ErrorResponse
// @Router      /problems/{ProblemID} [delete]
func DeleteProblemHandler(c echo.Context) error {
	str_id := c.Param("problemID")
	ID, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"invalid problem ID"})
	}

	if err = entities.DeleteProblemByID(uint(ID)); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	type response struct {
		Status string `json:"status"`
	}
	return c.JSON(http.StatusOK, response{"success"})
}
