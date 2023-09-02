package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"testing-nextalent/pkg/response"
)

type Svc interface {
	GetCountry(ctx context.Context, person string) (interface{}, error)
	GetTimeZone(ctx context.Context, timeZone string) (interface{}, error)
	ScriptInsertData(ctx context.Context) (interface{}, error)
}

type (
	// Handler ...
	Handler struct {
		svc Svc
	}
)

// New for bridging product handler initialization
func New(is Svc) *Handler {
	return &Handler{
		svc: is,
	}
}

// ApiHandler will receive request and return response
func (h *Handler) ApiHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      context.Context
		resp     *response.Response
		result   interface{}
		metadata interface{}
		err      error
		errRes   response.Error
	)

	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	ctx = context.Background()

	if err == nil {

		switch r.Method {
		// Check if request method is GET
		case http.MethodGet:
			_type := r.FormValue("type")
			value := r.FormValue("value")
			switch _type {
			case "getcountry":
				result, err = h.svc.GetCountry(ctx, value)
			case "gettimezone":
				result, err = h.svc.GetTimeZone(ctx, value)
			}
		// Check if request method is POST
		case http.MethodPost:
			result, err = h.svc.ScriptInsertData(ctx)
		// Check if request method is PUT
		case http.MethodPut:

		// Check if request method is DELETE
		case http.MethodDelete:

		default:
			err = errors.New("400")
		}
	}

	// If anything from service or data return an error
	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   101,
			Msg:    "101 - Data Not Found",
			Status: true,
		}
		// If service returns an error
		if strings.Contains(err.Error(), "service") {
			// Replace error with server error
			errRes = response.Error{
				Code:   500,
				Msg:    "500 - Internal Server Error",
				Status: true,
			}
		}
		// If error 401
		if strings.Contains(err.Error(), "401") {
			// Replace error with server error
			errRes = response.Error{
				Code:   401,
				Msg:    "401 - Unauthorized",
				Status: true,
			}
		}
		// If error 403
		if strings.Contains(err.Error(), "403") {
			// Replace error with server error
			errRes = response.Error{
				Code:   403,
				Msg:    "403 - Forbidden",
				Status: true,
			}
		}
		// If error 400
		if strings.Contains(err.Error(), "400") {
			// Replace error with server error
			errRes = response.Error{
				Code:   400,
				Msg:    "400 - Bad Request",
				Status: true,
			}
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.StatusCode = errRes.Code
		resp.Error = errRes
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
