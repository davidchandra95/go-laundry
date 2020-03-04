package http

import (
	"github.com/davidchandra95/go-laundry/modules"
	"github.com/davidchandra95/go-laundry/modules/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	UserService user.Service
}

func NewUserHandlers(userService user.Service) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}


func (u *UserHandler) FetchUsers(c *gin.Context) {
	var (
		err      error
		status   int
		start    = time.Now()
		data     interface{}
		response modules.APIResponse
	)

	defer func() {
		processTime := time.Since(start)
		response.Status = http.StatusOK
		response.ServerProcessTime = processTime.String()

		if err != nil {
			status = http.StatusInternalServerError
			response.MessageError = []string{err.Error()}
			c.JSON(status, gin.H{
				"message": err,
				"status":  status,
			})

			c.Abort()
			return
		}

		response.Data = data

		c.JSON(200, gin.H{"res": response})
		return
	}()

	data, err = u.UserService.GetUsers()
}

func (u *UserHandler) GetUser(c *gin.Context) {
	var (
		err      error
		status   int
		start    = time.Now()
		data     interface{}
		response modules.APIResponse

		id int64
	)
	defer func() {
		processTime := time.Since(start)
		response.Status = http.StatusOK
		response.ServerProcessTime = processTime.String()

		if err != nil {
			status = http.StatusInternalServerError
			response.MessageError = []string{err.Error()}
			c.JSON(status, gin.H{
				"message": err,
				"status":  status,
			})

			c.Abort()
			return
		}

		response.Data = data

		c.JSON(200, gin.H{"res": response})
		return
	}()

	idStr := c.Param("id")
	id, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}

	data, err = u.UserService.GetUser(id)
}
