package server

// NewGinHTTPHandlerDefault ...
func NewGinHTTPHandlerDefault() GinHTTPHandler {
	httpHandler := NewGinHTTPHandler("8080")
	return httpHandler
}
