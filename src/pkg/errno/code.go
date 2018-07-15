package errno

var (
	// Common errors
	// ------------------------------------------
	// 成功
	OK 	= &Errno{Code: 0, Message: "OK"}
	// 服务器错误
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}


	// 数据库错误
	ErrDatabase = &Errno{Code: 20002, Message: "Database error."}
	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}

	// 用户错误
	// --------------------------------------------
	ErrEncrypt = &Err{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound = &Err{Code: 20102, Message: "The user was not found."}
)