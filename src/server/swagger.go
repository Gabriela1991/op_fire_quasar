package server

// Success response
// swagger:response okResponse
type swaggRespOk struct {
	// in:body
	Body struct {
		// HTTP status code 200 - OK
		Message string `json:"message"`
	}
}

// Resource created
// swagger:response resourceCreated
type swaggRespResourceCreated struct {
	// in:body
	Body struct {
		// HTTP status code 201 - Resource created
		Message string `json:"message"`
	}
}

// Not found resource message
// swagger:response notFoundResource
type swaggRespNotFound struct {
	// in:body
	Body struct {
		// HTTP status code 204 - Not found resource message
		Message string `json:"message"`
	}
}

// Error Bad Request
// swagger:response badRequest
type swaggReqBadRequest struct {
	// in:body
	Body struct {
		// HTTP status code 400 -  Bad Request
		Message string `json:"message"`
	}
}

// Error Unauthorized Request
// swagger:response unauthorizedRequest
type swaggReqUnauthorized struct {
	// in:body
	Body struct {
		// HTTP status code 401 - Unauthorized Request
		Message string `json:"message"`
	}
}

// Error Forbidden
// swagger:response forbidden
type swaggReqForbidden struct {
	// in:body
	Body struct {
		// HTTP status code 403 - Forbidden
		Message string `json:"message"`
	}
}

// Error Not Found
// swagger:response notFoundReq
type swaggReqNotFound struct {
	// in:body
	Body struct {
		// HTTP status code 404 -  Not Found
		Message string `json:"message"`
	}
}

// Internal Server Error
// swagger:response internalServer
type swaggInternalServer struct {
	// in:body
	Body struct {
		// HTTP status code 500 - Internal Server Error
		Message string `json:"message"`
	}
}

// Conflict Error
// swagger:response conflict
type swaggConflict struct {
	// in:body
	Body struct {
		// HTTP status code 409 - Conflict Error
		Message string `json:"message"`
	}
}

// Error Service Unavailable
// swagger:response serviceUnavailable
type swaggSrvUnavailable struct {
	// in:body
	Body struct {
		// HTTP status code 503 -  Service Unavailable
		Message string `json:"message"`
	}
}
