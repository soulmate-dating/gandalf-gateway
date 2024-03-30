package response

type Response struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
}

func Error(err error) Response {
	return Response{Data: nil, Error: err.Error()}
}

func Success(data any) Response {
	return Response{Data: data, Error: ""}
}
