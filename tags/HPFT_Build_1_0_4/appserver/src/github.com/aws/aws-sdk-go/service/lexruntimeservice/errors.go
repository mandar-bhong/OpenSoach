// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexruntimeservice

const (

	// ErrCodeBadGatewayException for service response error code
	// "BadGatewayException".
	//
	// Either the Amazon Lex bot is still building, or one of the dependent services
	// (Amazon Polly, AWS Lambda) failed with an internal service error.
	ErrCodeBadGatewayException = "BadGatewayException"

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// Request validation failed, there is no usable message in the context, or
	// the bot build failed, is still in progress, or contains unbuilt changes.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// Two clients are using the same AWS account, Amazon Lex bot, and user ID.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeDependencyFailedException for service response error code
	// "DependencyFailedException".
	//
	// One of the dependencies, such as AWS Lambda or Amazon Polly, threw an exception.
	// For example,
	//
	//    * If Amazon Lex does not have sufficient permissions to call a Lambda
	//    function.
	//
	//    * If a Lambda function takes longer than 30 seconds to execute.
	//
	//    * If a fulfillment Lambda function returns a Delegate dialog action without
	//    removing any slot values.
	ErrCodeDependencyFailedException = "DependencyFailedException"

	// ErrCodeInternalFailureException for service response error code
	// "InternalFailureException".
	//
	// Internal service error. Retry the call.
	ErrCodeInternalFailureException = "InternalFailureException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Exceeded a limit.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeLoopDetectedException for service response error code
	// "LoopDetectedException".
	//
	// This exception is not used.
	ErrCodeLoopDetectedException = "LoopDetectedException"

	// ErrCodeNotAcceptableException for service response error code
	// "NotAcceptableException".
	//
	// The accept header in the request does not have a valid value.
	ErrCodeNotAcceptableException = "NotAcceptableException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// The resource (such as the Amazon Lex bot or an alias) that is referred to
	// is not found.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeRequestTimeoutException for service response error code
	// "RequestTimeoutException".
	//
	// The input speech is too long.
	ErrCodeRequestTimeoutException = "RequestTimeoutException"

	// ErrCodeUnsupportedMediaTypeException for service response error code
	// "UnsupportedMediaTypeException".
	//
	// The Content-Type header (PostContent API) has an invalid value.
	ErrCodeUnsupportedMediaTypeException = "UnsupportedMediaTypeException"
)
