package apperror

const (
	FailUnmarshalResponseBodyError ErrorType = "ER400 Fail to unmarshal response body"  // used by controller
	ObjectNotFound                 ErrorType = "ER404 Object %s is not found"           // used by injected repo in interactor
	UnrecognizedEnum               ErrorType = "ER500 %s is not recognized %s enum"     // used by enum
	DatabaseNotFoundInContextError ErrorType = "ER500 Database is not found in context" // used by repoimpl
)
