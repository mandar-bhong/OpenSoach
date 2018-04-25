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
    key: 'InfoMessageDetailsNotAvailable',
    en: 'Details are not available, please fill in the details and submit.'
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
];

