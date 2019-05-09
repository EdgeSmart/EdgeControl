package types

type (
	// ResponseJSON ResponseJSON
	ResponseJSON struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)
