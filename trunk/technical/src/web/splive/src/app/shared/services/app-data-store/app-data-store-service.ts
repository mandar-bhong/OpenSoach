import { Injectable } from '@angular/core';

import { APP_SHARED_IN_MEMORY_STORE_KEYS, APP_SHARED_LOCAL_STORAGE_KEYS } from '../../app-common-constants';
import { AppDataStore, AppInMemoryStore, AppLocalStorage } from './app-data-store';

@Injectable()
export class AppDataStoreService {
    private strategy = new Map<string, AppDataStore>();
    appInMemoryStoreKeys: string[];
    appLocalStorageKeys: string[];
    constructor() {
    }

    getDataStore(key: string) {
        return this.strategy.get(key);
    }

    init() {
        const inMemoryStore = new AppInMemoryStore();
        const localStorageStore = new AppLocalStorage();
        APP_SHARED_IN_MEMORY_STORE_KEYS.forEach(item => {
            this.strategy.set(item, inMemoryStore);
        });
        this.appInMemoryStoreKeys.forEach(item => {
            this.strategy.set(item, inMemoryStore);
        });

        APP_SHARED_LOCAL_STORAGE_KEYS.forEach(item => {
            this.strategy.set(item, localStorageStore);
        });

        this.appLocalStorageKeys.forEach(item => {
            this.strategy.set(item, localStorageStore);
        });
    }
}
