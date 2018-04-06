import { Injectable } from '@angular/core';
import { AppDataStore, AppInMemoryStore, AppLocalStorage } from './app-data-store';

export const APP_IN_MEMORY_STORE_KEYS: string[] = [
    'TEMP_OBJECT_1',
    'TEMP_OBJECT_2'
];

export const APP_LOCAL_STORAGE_KEYS: string[] = [
    'AUTH_TOKEN',
    'USER_ROLE'
];

@Injectable()
export class AppDataStoreService {
    public strategy = new Map<string, AppDataStore>();
    constructor() {
        const inMemoryStore = new AppInMemoryStore();
        const localStorageStore = new AppLocalStorage();
        APP_IN_MEMORY_STORE_KEYS.forEach(item => {
            this.strategy.set(item, inMemoryStore);
        });
        APP_LOCAL_STORAGE_KEYS.forEach(item => {
            this.strategy.set(item, localStorageStore);
        });
    }

    getDataStore(key: string) {
        return this.strategy.get(key);
    }
}
