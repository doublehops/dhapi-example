package app

type ContextVar string

func (cv ContextVar) String() string {
	return string(cv)
}

const (
	UserIDKey  ContextVar = "userID"
	TraceIDKey ContextVar = "traceID"
)
