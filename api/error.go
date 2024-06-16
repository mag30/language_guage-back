package api

import (
	"github.com/mag30/project-backend/domain/base"
	"net/http"
)

func ResponseFromServiceError(serviceError base.ServiceError, trackingID string) base.ResponseFailure {
	return base.ResponseFailure{
		Status:     http.StatusText(serviceError.Code),
		Blame:      serviceError.Blame,
		TrackingID: trackingID,
		Message:    serviceError.Message,
	}
}

func GeneralParsingError(trackingID string) base.ResponseFailure {
	return base.ResponseFailure{
		Status:     http.StatusText(http.StatusBadRequest),
		Blame:      base.BlameUser,
		TrackingID: trackingID,
		Message:    "failed to parse request parameters",
	}
}

func ResponseUnauthorizedError(trackingID string) base.ResponseFailure {
	return base.ResponseFailure{
		Status:     http.StatusText(http.StatusUnauthorized),
		Blame:      base.BlameUser,
		TrackingID: trackingID,
		Message:    "unauthorized",
	}
}
