import { PlatformHelperAPI } from "./platform-helper-api";

export class PlatformAndroidHelper implements PlatformHelperAPI {
    getRandomUUID(): string {
        return java.util.UUID.randomUUID().toString();
    }

    getSerialNumber(): string {
        // return android.os.Build.SERIAL;
        let serialNumber = '';
        try {
            var plugin = require("nativescript-uuid");
            serialNumber = plugin.getUUID();         
        } catch (err) {
            console.error('error occured while getting serial number', err);
        }
        return serialNumber;
    }
}