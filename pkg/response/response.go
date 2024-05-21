package response

type (
	Response struct {
		Message string `json:"message"`
	}
	WithData[T any] struct {
		Data T `json:"data"`
		Response
	}
	WithErrors struct {
		Errors []string `json:"errors"`
		Response
	}
)
