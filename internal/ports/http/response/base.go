package response

type Response struct {
	Data  any   `json:"data"`
	Error error `json:"error"`
}

func Error(err error) Response {
	return Response{Data: nil, Error: err}
}

func Success(data any) Response {
	return Response{Data: data, Error: nil}
}
