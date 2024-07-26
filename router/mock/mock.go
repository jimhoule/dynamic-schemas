package mock

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetRequestWithUrlParams(request *http.Request, urlParams map[string]string) *http.Request {
	routeContext := chi.NewRouteContext()
	for key, value := range urlParams {
		routeContext.URLParams.Add(key, value)
	}

	return request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, routeContext))
}