// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

const (

	// ErrCodeAccountSendingPausedException for service response error code
	// "AccountSendingPausedException".
	//
	// Indicates that email sending is disabled for your entire Amazon SES account.
	//
	// You can enable or disable email sending for your Amazon SES account using
	// UpdateAccountSendingEnabled.
	ErrCodeAccountSendingPausedException = "AccountSendingPausedException"

	// ErrCodeAlreadyExistsException for service response error code
	// "AlreadyExists".
	//
	// Indicates that a resource could not be created because of a naming conflict.
	ErrCodeAlreadyExistsException = "AlreadyExists"

	// ErrCodeCannotDeleteException for service response error code
	// "CannotDelete".
	//
	// Indicates that the delete operation could not be completed.
	ErrCodeCannotDeleteException = "CannotDelete"

	// ErrCodeConfigurationSetAlreadyExistsException for service response error code
	// "ConfigurationSetAlreadyExists".
	//
	// Indicates that the configuration set could not be created because of a naming
	// conflict.
	ErrCodeConfigurationSetAlreadyExistsException = "ConfigurationSetAlreadyExists"

	// ErrCodeConfigurationSetDoesNotExistException for service response error code
	// "ConfigurationSetDoesNotExist".
	//
	// Indicates that the configuration set does not exist.
	ErrCodeConfigurationSetDoesNotExistException = "ConfigurationSetDoesNotExist"

	// ErrCodeConfigurationSetSendingPausedException for service response error code
	// "ConfigurationSetSendingPausedException".
	//
	// Indicates that email sending is disabled for the configuration set.
	//
	// You can enable or disable email sending for a configuration set using UpdateConfigurationSetSendingEnabled.
	ErrCodeConfigurationSetSendingPausedException = "ConfigurationSetSendingPausedException"

	// ErrCodeCustomVerificationEmailInvalidContentException for service response error code
	// "CustomVerificationEmailInvalidContent".
	//
	// Indicates that custom verification email template provided content is invalid.
	ErrCodeCustomVerificationEmailInvalidContentException = "CustomVerificationEmailInvalidContent"

	// ErrCodeCustomVerificationEmailTemplateAlreadyExistsException for service response error code
	// "CustomVerificationEmailTemplateAlreadyExists".
	//
	// Indicates that a custom verification email template with the name you specified
	// already exists.
	ErrCodeCustomVerificationEmailTemplateAlreadyExistsException = "CustomVerificationEmailTemplateAlreadyExists"

	// ErrCodeCustomVerificationEmailTemplateDoesNotExistException for service response error code
	// "CustomVerificationEmailTemplateDoesNotExist".
	//
	// Indicates that a custom verification email template with the name you specified
	// does not exist.
	ErrCodeCustomVerificationEmailTemplateDoesNotExistException = "CustomVerificationEmailTemplateDoesNotExist"

	// ErrCodeEventDestinationAlreadyExistsException for service response error code
	// "EventDestinationAlreadyExists".
	//
	// Indicates that the event destination could not be created because of a naming
	// conflict.
	ErrCodeEventDestinationAlreadyExistsException = "EventDestinationAlreadyExists"

	// ErrCodeEventDestinationDoesNotExistException for service response error code
	// "EventDestinationDoesNotExist".
	//
	// Indicates that the event destination does not exist.
	ErrCodeEventDestinationDoesNotExistException = "EventDestinationDoesNotExist"

	// ErrCodeFromEmailAddressNotVerifiedException for service response error code
	// "FromEmailAddressNotVerified".
	//
	// Indicates that the sender address specified for a custom verification email
	// is not verified, and is therefore not eligible to send the custom verification
	// email.
	ErrCodeFromEmailAddressNotVerifiedException = "FromEmailAddressNotVerified"

	// ErrCodeInvalidCloudWatchDestinationException for service response error code
	// "InvalidCloudWatchDestination".
	//
	// Indicates that the Amazon CloudWatch destination is invalid. See the error
	// message for details.
	ErrCodeInvalidCloudWatchDestinationException = "InvalidCloudWatchDestination"

	// ErrCodeInvalidConfigurationSetException for service response error code
	// "InvalidConfigurationSet".
	//
	// Indicates that the configuration set is invalid. See the error message for
	// details.
	ErrCodeInvalidConfigurationSetException = "InvalidConfigurationSet"

	// ErrCodeInvalidFirehoseDestinationException for service response error code
	// "InvalidFirehoseDestination".
	//
	// Indicates that the Amazon Kinesis Firehose destination is invalid. See the
	// error message for details.
	ErrCodeInvalidFirehoseDestinationException = "InvalidFirehoseDestination"

	// ErrCodeInvalidLambdaFunctionException for service response error code
	// "InvalidLambdaFunction".
	//
	// Indicates that the provided AWS Lambda function is invalid, or that Amazon
	// SES could not execute the provided function, possibly due to permissions
	// issues. For information about giving permissions, see the Amazon SES Developer
	// Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html).
	ErrCodeInvalidLambdaFunctionException = "InvalidLambdaFunction"

	// ErrCodeInvalidPolicyException for service response error code
	// "InvalidPolicy".
	//
	// Indicates that the provided policy is invalid. Check the error stack for
	// more information about what caused the error.
	ErrCodeInvalidPolicyException = "InvalidPolicy"

	// ErrCodeInvalidRenderingParameterException for service response error code
	// "InvalidRenderingParameter".
	//
	// Indicates that one or more of the replacement values you provided is invalid.
	// This error may occur when the TemplateData object contains invalid JSON.
	ErrCodeInvalidRenderingParameterException = "InvalidRenderingParameter"

	// ErrCodeInvalidS3ConfigurationException for service response error code
	// "InvalidS3Configuration".
	//
	// Indicates that the provided Amazon S3 bucket or AWS KMS encryption key is
	// invalid, or that Amazon SES could not publish to the bucket, possibly due
	// to permissions issues. For information about giving permissions, see the
	// Amazon SES Developer Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html).
	ErrCodeInvalidS3ConfigurationException = "InvalidS3Configuration"

	// ErrCodeInvalidSNSDestinationException for service response error code
	// "InvalidSNSDestination".
	//
	// Indicates that the Amazon Simple Notification Service (Amazon SNS) destination
	// is invalid. See the error message for details.
	ErrCodeInvalidSNSDestinationException = "InvalidSNSDestination"

	// ErrCodeInvalidSnsTopicException for service response error code
	// "InvalidSnsTopic".
	//
	// Indicates that the provided Amazon SNS topic is invalid, or that Amazon SES
	// could not publish to the topic, possibly due to permissions issues. For information
	// about giving permissions, see the Amazon SES Developer Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html).
	ErrCodeInvalidSnsTopicException = "InvalidSnsTopic"

	// ErrCodeInvalidTemplateException for service response error code
	// "InvalidTemplate".
	//
	// Indicates that the template that you specified could not be rendered. This
	// issue may occur when a template refers to a partial that does not exist.
	ErrCodeInvalidTemplateException = "InvalidTemplate"

	// ErrCodeInvalidTrackingOptionsException for service response error code
	// "InvalidTrackingOptions".
	//
	// Indicates that the custom domain to be used for open and click tracking redirects
	// is invalid. This error appears most often in the following situations:
	//
	//    * When the tracking domain you specified is not verified in Amazon SES.
	//
	//    * When the tracking domain you specified is not a valid domain or subdomain.
	ErrCodeInvalidTrackingOptionsException = "InvalidTrackingOptions"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceeded".
	//
	// Indicates that a resource could not be created because of service limits.
	// For a list of Amazon SES limits, see the Amazon SES Developer Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/limits.html).
	ErrCodeLimitExceededException = "LimitExceeded"

	// ErrCodeMailFromDomainNotVerifiedException for service response error code
	// "MailFromDomainNotVerifiedException".
	//
	// Indicates that the message could not be sent because Amazon SES could not
	// read the MX record required to use the specified MAIL FROM domain. For information
	// about editing the custom MAIL FROM domain settings for an identity, see the
	// Amazon SES Developer Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from-edit.html).
	ErrCodeMailFromDomainNotVerifiedException = "MailFromDomainNotVerifiedException"

	// ErrCodeMessageRejected for service response error code
	// "MessageRejected".
	//
	// Indicates that the action failed, and the message could not be sent. Check
	// the error stack for more information about what caused the error.
	ErrCodeMessageRejected = "MessageRejected"

	// ErrCodeMissingRenderingAttributeException for service response error code
	// "MissingRenderingAttribute".
	//
	// Indicates that one or more of the replacement values for the specified template
	// was not specified. Ensure that the TemplateData object contains references
	// to all of the replacement tags in the specified template.
	ErrCodeMissingRenderingAttributeException = "MissingRenderingAttribute"

	// ErrCodeProductionAccessNotGrantedException for service response error code
	// "ProductionAccessNotGranted".
	//
	// Indicates that the account has not been granted production access.
	ErrCodeProductionAccessNotGrantedException = "ProductionAccessNotGranted"

	// ErrCodeRuleDoesNotExistException for service response error code
	// "RuleDoesNotExist".
	//
	// Indicates that the provided receipt rule does not exist.
	ErrCodeRuleDoesNotExistException = "RuleDoesNotExist"

	// ErrCodeRuleSetDoesNotExistException for service response error code
	// "RuleSetDoesNotExist".
	//
	// Indicates that the provided receipt rule set does not exist.
	ErrCodeRuleSetDoesNotExistException = "RuleSetDoesNotExist"

	// ErrCodeTemplateDoesNotExistException for service response error code
	// "TemplateDoesNotExist".
	//
	// Indicates that the Template object you specified does not exist in your Amazon
	// SES account.
	ErrCodeTemplateDoesNotExistException = "TemplateDoesNotExist"

	// ErrCodeTrackingOptionsAlreadyExistsException for service response error code
	// "TrackingOptionsAlreadyExistsException".
	//
	// Indicates that the configuration set you specified already contains a TrackingOptions
	// object.
	ErrCodeTrackingOptionsAlreadyExistsException = "TrackingOptionsAlreadyExistsException"

	// ErrCodeTrackingOptionsDoesNotExistException for service response error code
	// "TrackingOptionsDoesNotExistException".
	//
	// Indicates that the TrackingOptions object you specified does not exist.
	ErrCodeTrackingOptionsDoesNotExistException = "TrackingOptionsDoesNotExistException"
)
