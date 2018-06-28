export class TranslationValue {
  key: string;
  en: string;
}

export const TRANSLATIONS: TranslationValue[] = [
  {
    key: 'LblLoginHeader',
    en: 'Sign In'
  },
  {
    key: 'AppNotificationSuccess',
    en: 'Success'
  },
  {
    key: 'AppNotificationError',
    en: 'Error'
  },
  {
    key: 'AppNotificationAlert',
    en: 'Alert'
  },
  {
    key: 'AppNotificationInformation',
    en: 'Information'
  },
  {
    key: 'AppNotificationWarning',
    en: 'Warning'
  },
  {
    key: 'SERVER_SYSTEM_ERROR',
    en: 'A System error has occurred: Error code'
  },
  {
    key: 'SERVER_ERROR',
    en: 'An Error has occurred: Eror code'
  },
  {
    key: 'SERVER_UNKNOWN_ERROR',
    en: 'An Error has occurred'
  },
  {
    key: 'SERVER_ERROR_10001',
    en: 'SERVER_ERROR_10001'
  },
  {
    key: 'INFO_DETAILS_NOT_AVAILABLE',
    en: 'Details are not available, please fill in the details and save.'
  },
  {
    key: 'USER_INFO_DETAILS_NOT_AVAILABLE',
    en: 'User profile details are not available.'
  },
  {
    key: 'TASK_NOT_AVAILABLE',
    en: 'Charts task are not availbale, please fill in chats task'
  },

  /* CUSTOMER STATES BEGIN */
  {
    key: 'CUSTOMER_STATE_0',
    en: 'NOT SET'
  },
  {
    key: 'CUSTOMER_STATE_1',
    en: 'ACTIVE'
  },
  {
    key: 'CUSTOMER_STATE_2',
    en: 'INACTIVE'
  },
  {
    key: 'CUSTOMER_STATE_3',
    en: 'SUSPENDED'
  },

  /* CUSTOMER STATES END */

  /* OPERATOR STATES  */

  {
    key: 'OPERATOR_STATE_1',
    en: 'ACTIVE'
  },
  {
    key: 'OPERATOR_STATE_2',
    en: 'INACTIVE'
  },
  /* OPERATOR STATES END */

  /* OPERATOR AREA  START */

  {
    key: 'OPERATOR_AREA_1',
    en: 'OPEN'
  },
  {
    key: 'OPERATOR_AREA_2',
    en: 'RESTRICTED'
  },
  /* OPERATOR AREA END */

  /* COMPLAINT STATE  START */
  {
    key: 'COMPLAINT_STATE_1',
    en: 'OPEN'
  },
  {
    key: 'COMPLAINT_STATE_2',
    en: 'CLOSE'
  },
  {
    key: 'COMPLAINT_STATE_3',
    en: 'INPROGRESS'
  },
  /* COMPLAINT STATE END */

  /* SEVERITY STATE  START */
  {
    key: 'SEVERIT_STATE_1',
    en: 'HIGH'
  },
  {
    key: 'SEVERIT_STATE_2',
    en: 'MEDIUM'
  },
  {
    key: 'SEVERIT_STATE_3',
    en: 'LOW'
  },
  {
    key: 'SEVERIT_STATE_4',
    en: 'CRITICAL'
  },
  /* SEVERITY STATE END */

  /* USER STATES BEGIN */
  {
    key: 'USER_STATE_0',
    en: 'NOT SET'
  },
  {
    key: 'USER_STATE_1',
    en: 'ACTIVE'
  },
  {
    key: 'USER_STATE_2',
    en: 'INACTIVE'
  },
  {
    key: 'USER_STATE_3',
    en: 'SUSPENDED'
  },

  /* USER STATES END */

  /* USER PRODUCT MAPPING STATES BEGIN */
  {
    key: 'USER_PRODUCT_MAPPING_STATE_0',
    en: 'NOT SET'
  },
  {
    key: 'USER_PRODUCT_MAPPING_STATE_1',
    en: 'ACTIVE'
  },
  {
    key: 'USER_PRODUCT_MAPPING_STATE_2',
    en: 'INACTIVE'
  },
  {
    key: 'USER_PRODUCT_MAPPING_STATE_3',
    en: 'SUSPENDED'
  },

  /* USER PRODUCT MAPPING STATES  END */

  /* USER CATEGORYS BEGIN */

  {
    key: 'USER_CATEGORY_1',
    en: 'OSU'
  },
  {
    key: 'USER_CATEGORY_2',
    en: 'CU'
  },

  /* USER CATEGORYS END */

  /* USER GENDER BEGIN */
  {
    key: 'USER_GENDER_0',
    en: 'NOT_SELECTED'
  },
  {
    key: 'USER_GENDER_1',
    en: 'MALE'
  },
  {
    key: 'USER_GENDER_2',
    en: 'FEMALE'
  },

  /* USER GENDER END */

  /* DEVICE STATES BEGIN */
  {
    key: 'DEVICE_STATE_0',
    en: 'NOT SET'
  },
  {
    key: 'DEVICE_STATE_1',
    en: 'ACTIVE'
  },
  {
    key: 'DEVICE_STATE_2',
    en: 'INACTIVE'
  },
  {
    key: 'DEVICE_STATE_3',
    en: 'SUSPENDED'
  },

  /* DEVICE STATES END */

  /*SERVICEPOINT STATE  */
  {
    key: 'SERVICEPOINT_STATE_1',
    en: 'ACTIVE'
  },
  {
    key: 'SERVICEPOINT_STATE_2',
    en: 'INACTIVE'
  },
  {
    key: 'SERVICEPOINT_STATE_3',
    en: 'SUSPENDED'
  },
  /*SERVICEPOINT STATE  */

  {
    key: 'SUCCESS_CUSTOMER_DETAILS_SAVED',
    en: 'Successfully add device.'
  },
  {
    key: 'ERROR_LOGIN_INVALID_CATEGORY',
    en: 'Unauthorized Access'
  },



  /* CORPORATE STATES START */
  {
    key: 'SUCCESS_CORPORATE_ADD_SAVED',
    en: 'Corporate add successfully.'
  },
  {
    key: 'ERROR_LOGIN_INVALID_CATEGORY',
    en: 'Unauthorized Access'
  },
  {
    key: 'SUCCESS_CORPORATE_DETAILS_SAVED',
    en: 'Corporate details save successfully'
  },

  /* CORPORATE APP NOTIFICATION END */

  /* DEIVCE APP NOTIFICATION START */
  {
    key: 'CUSTOMER_IS_NOT_BEEN_ASSOCIATED_WITH_ANY_PRODUCT',
    en: 'You need to add a customer before associating a device.'
  },
  {
    key: 'SUCCESS_ADD_DEVICE_ASSOCIATE_SAVED',
    en: 'Successfully added device associate.'
  },
  {
    key: 'SUCCESS_ADD_DEVICE_SAVED',
    en: 'Device added successfully. '
  },
  {
    key: 'SUCCESS_DEVICE_DETAILS_SAVED',
    en: 'Device updated successfully. '
  },
  {
    key: 'INFO_NO_RECORDS_FOUND',
    en: 'No record(s) found. '
  },

  /* DEIVCE APP NOTIFICATION END */


  /* USER APP NOTIFICATION START */
  {
    key: 'CUSTOMER_IS_NOT_BEEN_ASSOCIATED_WITH_ANY_PRODUCT_USER',
    en: 'You need to add a customer before associating a user.'
  },
  {
    key: 'PRODUCT_IS_NOT_ASSOCIATED',
    en: 'You need to add a product before associating a user.'
  },

  /* USER APP NOTIFICATION END */




  /* CUSTOMER APP NOTIFICATION START */

  {
    key: 'CUSTOMER_IS_NOT_ASSOCIATED',
    en: 'Customer not associated with any product.'
  },

  /* CUSTOMER APP NOTIFICATION END */


  /* CUSTOMER PRODUCT MAPPING STATES BEGIN */
  {
    key: 'CUSTOMER_PRODUCT_MAPPING_STATE_0',
    en: 'NOT SET'
  },
  {
    key: 'CUSTOMER_PRODUCT_MAPPING_STATE_1',
    en: 'ACTIVE'
  },
  {
    key: 'CUSTOMER_PRODUCT_MAPPING_STATE_2',
    en: 'INACTIVE'
  },
  {
    key: 'CUSTOMER_PRODUCT_MAPPING_STATE_3',
    en: 'SUSPENDED'
  },

  /* CUSTOMER PRODUCT MAPPING STATES END */


  /* VALIDATION MESSAGES START */
  {
    key: 'VALIDATION_REQUIRED_FIELD',
    en: 'You can\'t leave this empty'
  },

  /* VALIDATION MESSAGES END */

  {
    key: 'CHART_DATA_NO_CHART_CONFIGURED',
    en: 'No Chart has been configured yet. Goto Service Points-> List -> Configure to configure a chart'
  },

];

