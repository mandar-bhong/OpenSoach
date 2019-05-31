// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

const (

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// The resource hierarchy is changing.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeConflictingOperationException for service response error code
	// "ConflictingOperationException".
	//
	// Another operation is in progress on the resource that conflicts with the
	// current operation.
	ErrCodeConflictingOperationException = "ConflictingOperationException"

	// ErrCodeCustomMetadataLimitExceededException for service response error code
	// "CustomMetadataLimitExceededException".
	//
	// The limit has been reached on the number of custom properties for the specified
	// resource.
	ErrCodeCustomMetadataLimitExceededException = "CustomMetadataLimitExceededException"

	// ErrCodeDeactivatingLastSystemUserException for service response error code
	// "DeactivatingLastSystemUserException".
	//
	// The last user in the organization is being deactivated.
	ErrCodeDeactivatingLastSystemUserException = "DeactivatingLastSystemUserException"

	// ErrCodeDocumentLockedForCommentsException for service response error code
	// "DocumentLockedForCommentsException".
	//
	// This exception is thrown when the document is locked for comments and user
	// tries to create or delete a comment on that document.
	ErrCodeDocumentLockedForCommentsException = "DocumentLockedForCommentsException"

	// ErrCodeDraftUploadOutOfSyncException for service response error code
	// "DraftUploadOutOfSyncException".
	//
	// This exception is thrown when a valid checkout ID is not presented on document
	// version upload calls for a document that has been checked out from Web client.
	ErrCodeDraftUploadOutOfSyncException = "DraftUploadOutOfSyncException"

	// ErrCodeEntityAlreadyExistsException for service response error code
	// "EntityAlreadyExistsException".
	//
	// The resource already exists.
	ErrCodeEntityAlreadyExistsException = "EntityAlreadyExistsException"

	// ErrCodeEntityNotExistsException for service response error code
	// "EntityNotExistsException".
	//
	// The resource does not exist.
	ErrCodeEntityNotExistsException = "EntityNotExistsException"

	// ErrCodeFailedDependencyException for service response error code
	// "FailedDependencyException".
	//
	// The AWS Directory Service cannot reach an on-premises instance. Or a dependency
	// under the control of the organization is failing, such as a connected Active
	// Directory.
	ErrCodeFailedDependencyException = "FailedDependencyException"

	// ErrCodeIllegalUserStateException for service response error code
	// "IllegalUserStateException".
	//
	// The user is undergoing transfer of ownership.
	ErrCodeIllegalUserStateException = "IllegalUserStateException"

	// ErrCodeInvalidArgumentException for service response error code
	// "InvalidArgumentException".
	//
	// The pagination marker or limit fields are not valid.
	ErrCodeInvalidArgumentException = "InvalidArgumentException"

	// ErrCodeInvalidCommentOperationException for service response error code
	// "InvalidCommentOperationException".
	//
	// The requested operation is not allowed on the specified comment object.
	ErrCodeInvalidCommentOperationException = "InvalidCommentOperationException"

	// ErrCodeInvalidOperationException for service response error code
	// "InvalidOperationException".
	//
	// The operation is invalid.
	ErrCodeInvalidOperationException = "InvalidOperationException"

	// ErrCodeInvalidPasswordException for service response error code
	// "InvalidPasswordException".
	//
	// The password is invalid.
	ErrCodeInvalidPasswordException = "InvalidPasswordException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The maximum of 100,000 folders under the parent folder has been exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeProhibitedStateException for service response error code
	// "ProhibitedStateException".
	//
	// The specified document version is not in the INITIALIZED state.
	ErrCodeProhibitedStateException = "ProhibitedStateException"

	// ErrCodeRequestedEntityTooLargeException for service response error code
	// "RequestedEntityTooLargeException".
	//
	// The response is too large to return. The request must include a filter to
	// reduce the size of the response.
	ErrCodeRequestedEntityTooLargeException = "RequestedEntityTooLargeException"

	// ErrCodeResourceAlreadyCheckedOutException for service response error code
	// "ResourceAlreadyCheckedOutException".
	//
	// The resource is already checked out.
	ErrCodeResourceAlreadyCheckedOutException = "ResourceAlreadyCheckedOutException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// One or more of the dependencies is unavailable.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeStorageLimitExceededException for service response error code
	// "StorageLimitExceededException".
	//
	// The storage limit has been exceeded.
	ErrCodeStorageLimitExceededException = "StorageLimitExceededException"

	// ErrCodeStorageLimitWillExceedException for service response error code
	// "StorageLimitWillExceedException".
	//
	// The storage limit will be exceeded.
	ErrCodeStorageLimitWillExceedException = "StorageLimitWillExceedException"

	// ErrCodeTooManyLabelsException for service response error code
	// "TooManyLabelsException".
	//
	// The limit has been reached on the number of labels for the specified resource.
	ErrCodeTooManyLabelsException = "TooManyLabelsException"

	// ErrCodeTooManySubscriptionsException for service response error code
	// "TooManySubscriptionsException".
	//
	// You've reached the limit on the number of subscriptions for the WorkDocs
	// instance.
	ErrCodeTooManySubscriptionsException = "TooManySubscriptionsException"

	// ErrCodeUnauthorizedOperationException for service response error code
	// "UnauthorizedOperationException".
	//
	// The operation is not permitted.
	ErrCodeUnauthorizedOperationException = "UnauthorizedOperationException"

	// ErrCodeUnauthorizedResourceAccessException for service response error code
	// "UnauthorizedResourceAccessException".
	//
	// The caller does not have access to perform the action on the resource.
	ErrCodeUnauthorizedResourceAccessException = "UnauthorizedResourceAccessException"
)
