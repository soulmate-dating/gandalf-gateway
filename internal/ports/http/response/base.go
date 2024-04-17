package response

// Response is a generic response structure.
type Response struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

func Error(err error) Response {
	return Response{Data: nil, Error: err.Error()}
}

func Success(data any) Response {
	return Response{Data: data, Error: nil}
}
