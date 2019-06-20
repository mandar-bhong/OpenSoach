import { PlatformHelperAPI } from "./platform-helper-api";

export class PlatformiOSHelper implements PlatformHelperAPI {
    getRandomUUID(): string {
        return "";
    }
    getSerialNumber(): string {
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