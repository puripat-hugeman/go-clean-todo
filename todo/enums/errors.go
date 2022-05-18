package enums

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	ErrFile error = errors.New("error working with multipart file")
)

var (
	MapErrHandler = struct {
		MultipartError gin.H
		Unmarshal      gin.H
	}{
		MultipartError: mapErrMultipart,
		Unmarshal:      mapErrUnmarshal,
	}

	mapErrMultipart gin.H = gin.H{
		"status": "bad multipart request",
	}
	mapErrUnmarshal gin.H = gin.H{
		"status": "failed to unmarshal JSON",
	}
)
