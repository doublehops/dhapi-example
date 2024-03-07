package app

type ContextVar string

const (
	UserIDKey  ContextVar = "userID"
	TraceIDKey ContextVar = "traceID"
)
