package errors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckError(err error) (int, map[string]any) {
	switch err.(error) {
	case &NotFoundError{}:
		return http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)}
	default:
		return http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("%v", err)}
	}
}

type NotFoundError struct{}

func (nf *NotFoundError) Error() string {
	return "Not found record in database "
}
