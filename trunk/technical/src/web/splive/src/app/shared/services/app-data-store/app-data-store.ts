
export interface AppDataStore {
    getObject<T>(key: string): T | null;
    setObject<T>(key: string, data: T): void;
    removeObject(key: string): void;
    clear(): void;
}

export class AppLocalStorage implements AppDataStore {
    getObject<T>(key: string): T | null {
        const valueJSON = localStorage.getItem(key);
        if (valueJSON) {
            return JSON.parse(valueJSON);
        }
    }
    setObject<T>(key: string, data: T): void {
        if (data) {
            localStorage.setItem(key, JSON.stringify(data));
        }
    }
    removeObject(key: string): void {
        localStorage.removeItem(key);
    }
    clear(): void {
        localStorage.clear();
    }
}

export class AppInMemoryStore implements AppDataStore {
    private appDataStore = new Map<string, string>();
    getObject<T>(key: string): T | null {
        const valueJSON = this.appDataStore.get(key);
        if (valueJSON) {
            return JSON.parse(valueJSON);
        }
    }
    setObject<T>(key: string, data: T): void {
        if (data) {
            this.appDataStore.set(key, JSON.stringify(data));
        }
    }
    removeObject(key: string): void {
        this.appDataStore.delete(key);
    }
    clear(): void {
        this.appDataStore.clear();
    }
}
