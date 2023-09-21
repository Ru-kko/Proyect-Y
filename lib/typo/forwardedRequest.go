package typo

type ForwardedRequest[T any] struct {
	Data T `json:"data,omitempty"`
}
