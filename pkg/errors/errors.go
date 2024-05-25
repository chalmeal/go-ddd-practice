package errors

import "errors"

var (
	// StatusOK(200)
	LOGIN_SUCCESS            = errors.New("Login Successfully")
	REGISTER_ACCOUNT_SUCCESS = errors.New("Register Account Successfully")
	EDIT_ACCOUNT_SUCCESS     = errors.New("Edit Account Successfully")
	DELETE_ACCOUNT_SUCCESS   = errors.New("Delete Account Successfully")
	// BadRequest(400)
	LOGIN_BAD_REQUEST             = errors.New("Login -> Bad Request")
	LOGIN_FAILURE_UPDATE_TOKEN    = errors.New("Login -> Update access token failure")
	GET_ACCOUNT_ACCOUNT_NOT_FOUND = errors.New("GetAccount -> Account not found")
	REGISTER_ACCOUNT_BAD_REQUEST  = errors.New("RegisterAccount -> Bad request")
	EDIT_ACCOUNT_BAD_REQUEST      = errors.New("EditAccount -> Bad request")
	DELETE_ACCOUNT_BAD_REQUEST    = errors.New("DeleteAccount -> Bad request")
	DELETE_ACCOUNT_DELETED        = errors.New("DeleteAccount -> Already deleted account")
	// UnAuthorization(401)
	LOGIN_FAILURE = errors.New("Login -> ID or password is different")
	// InternalServerError(500)
	INTERNAL_SERVER_ERROR = errors.New("There was a problem on the server side")
)

var (
	BadRequest = []error{
		LOGIN_BAD_REQUEST,
		LOGIN_FAILURE_UPDATE_TOKEN,
		GET_ACCOUNT_ACCOUNT_NOT_FOUND,
		REGISTER_ACCOUNT_BAD_REQUEST,
		EDIT_ACCOUNT_BAD_REQUEST,
		DELETE_ACCOUNT_BAD_REQUEST,
		DELETE_ACCOUNT_DELETED,
	}
	UnAuthorized = []error{
		LOGIN_FAILURE,
	}
)
