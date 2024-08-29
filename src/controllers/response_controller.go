package controllers

import "github.com/nanda03dev/gcart/src/common"

func ToSuccessResponse(msg any, data any) common.SuccessResponse {
	return common.SuccessResponse{
		Msg:  msg,
		Data: data,
	}
}

func ToErrorResponse(msg any, data any) common.ErrorResponse {
	return common.ErrorResponse{
		Msg:   msg,
		Error: data,
	}
}
