import { Injectable } from '@angular/core';

import { EnvironmentProvider } from '../../environment-provider';
import { TRANSLATIONS, TranslationValue } from './translations';

@Injectable()
export class TranslateService {
  private currentLanguage = 'en';
  public translate: Function;
  translateMap: Map<string, TranslationValue>;
  constructor() {
    if (EnvironmentProvider.production) {
      this.translate = this.translateProdEnv;
    } else {
      this.translate = this.translateDevEnv;
    }

    // create a map of TRANSLATIONS
    this.translateMap = new Map(TRANSLATIONS.map(x => [x.key, x] as [string, TranslationValue]));
  }

  translateProdEnv(str) {
    const resourceValue = this.translateMap.get(str);
    return resourceValue ? resourceValue[this.currentLanguage] : str;
  }

  translateDevEnv(str) {
    const resourceValue = this.translateMap.get(str);
    return resourceValue ? resourceValue[this.currentLanguage] : '';
  }

  selectLanguage(language) {
    this.currentLanguage = language;
  }

  addAppSpecificTranslations(appTranslations: TranslationValue[]) {
    console.log('in transaltions');
    appTranslations.forEach(element => {

      this.translateMap.set(element.key, element);
    });
  }
}
