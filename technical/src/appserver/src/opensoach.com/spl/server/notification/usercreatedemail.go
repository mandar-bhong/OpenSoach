package notification

import (
	"strings"

	"opensoach.com/core/logger"
	engnotemail "opensoach.com/engine/notification/email"
	gnotmodels "opensoach.com/models/notification"
	"opensoach.com/spl/server/constants"
	"opensoach.com/spl/server/dbaccess"
	lmodels "opensoach.com/spl/server/models"
	repo "opensoach.com/spl/server/repository"
)

type UserEmailNotification struct {
	gnotmodels.EmailOptions
	gnotmodels.NotificationConfigModel
}

func SendUserAssociatedEmailNotification(toEmail, code string) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Debug, "Executing SendUserAssociatedEmailNotification")

	SendEmailNotification(toEmail, code, constants.DB_EMAIL_TML_USER_ASSOCIATED)

}

func SendUserOtpEmailNotification(toEmail, otpcode string) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Debug, "Executing SendUserOtpEmailNotification")

	SendEmailNotification(toEmail, otpcode, constants.DB_EMAIL_TML_USER_OTP)

}

func SendEmailNotification(toEmail, replacablecode, templatecode string) {

	userEmailNotification := &UserEmailNotification{}

	dbErr, templateData := dbaccess.GetEmailTemplate(repo.Instance().Context.Master.DBConn, templatecode)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting email template from database", dbErr)
		return
	}

	userEmailNotification.To = []string{toEmail}
	userEmailNotification.From = repo.Instance().Config.EmailConfig.From

	userEmailNotification.Subject = templateData.Subject
	userEmailNotification.Body = strings.Replace(templateData.Body, "$Code$", replacablecode, -1)

	userEmailNotification.SMTPAddress = repo.Instance().Config.EmailConfig.SMTPAddress
	userEmailNotification.SMTPUsername = repo.Instance().Config.EmailConfig.SMTPUsername
	userEmailNotification.SMTPPassword = repo.Instance().Config.EmailConfig.SMTPPassword
	userEmailNotification.SMTPPort = repo.Instance().Config.EmailConfig.SMTPPort

	dbEmailRowModel := lmodels.DBEmailRowModel{}
	dbEmailRowModel.Subject = userEmailNotification.Subject
	dbEmailRowModel.Body = userEmailNotification.Body
	dbEmailRowModel.TemplateID = templateData.ID

	dbEmailSaveErr, insertedID := dbaccess.SaveEmail(repo.Instance().Context.Master.DBConn, dbEmailRowModel)

	if dbEmailSaveErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while saving email", dbEmailSaveErr)
	}

	emailSendErr := engnotemail.SendEmail(userEmailNotification)

	if emailSendErr != nil {
		dbEmailRowModel.Status = constants.DB_EMAIL_SEND_FAILED
		sendErr := emailSendErr.Error()
		dbEmailRowModel.Comment = &sendErr
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to send user created email", emailSendErr)
		dbaccess.UpdateEmailStatus(repo.Instance().Context.Master.DBConn, dbEmailRowModel)
		return
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Debug, "Email sent successfully")

	dbEmailRowModel.Status = constants.DB_EMAIL_SEND_SUCCESS
	dbEmailRowModel.ID = insertedID
	statusUpdateErr := dbaccess.UpdateEmailStatus(repo.Instance().Context.Master.DBConn, dbEmailRowModel)

	if statusUpdateErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to save email send success status", emailSendErr)
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Debug, "Email status updated successfully")
}

func (r *UserEmailNotification) GetEmailOptions() gnotmodels.EmailOptions {
	return r.EmailOptions
}

func (r *UserEmailNotification) GetNotificationConfig() gnotmodels.NotificationConfigModel {
	return r.NotificationConfigModel
}
