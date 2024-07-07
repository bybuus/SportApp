package dto

type Response struct {
	Status Status      `json:"status"`
	Result interface{} `json:"result"`
}

type Status struct {
	Ok          bool   `json:"ok"`
	ErrorReason string `json:"error"`
}

func InternalErrorResponse() Response {
	return Response{
		Status: Status{
			Ok:          false,
			ErrorReason: "Internal server err",
		},
		Result: nil,
	}
}

func ErrorResponse(reason string) Response {
	return Response{
		Status: Status{
			Ok:          false,
			ErrorReason: reason,
		},
		Result: nil,
	}
}

func Ok(result interface{}) Response {
	return Response{
		Status: Status{
			Ok:          true,
			ErrorReason: "",
		},
		Result: result,
	}
}

func OkWithoutBody() Response {
	return Ok(nil)
}
