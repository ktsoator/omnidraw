package resp

// Result is the unified response structure for all API endpoints.
// It ensures that the frontend always receives a consistent JSON format,
// regardless of whether the request succeeded or failed.
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
