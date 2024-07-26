package utils

import (
	"encoding/json"
	"net/http"
)

type HttpError struct{
	Error error
	Status int
	Message string
}

func ReadHttpRequestBody(writer http.ResponseWriter, request *http.Request, requestBody any) error {
	decoder := json.NewDecoder(request.Body);
	err := decoder.Decode(requestBody);
	if err != nil {
		return err;
	}

	return nil
}

func WriteHttpResponse(writer http.ResponseWriter, status int, response any) error {
	// Converts response to JSON
	jsonResponse, err := json.Marshal(response);
	if err != nil {
		return err;
	}


	// Writes headers
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status);

	// Writes response
	_, err = writer.Write(jsonResponse);
	if err != nil {
		return err;
	}

	return nil
}

func WriteHttpError(writer http.ResponseWriter, status int, err error) error {
	httpError := HttpError{
		Error: err,
		Status: status,
		Message: err.Error(),
	}

	return WriteHttpResponse(writer, status, httpError);
}