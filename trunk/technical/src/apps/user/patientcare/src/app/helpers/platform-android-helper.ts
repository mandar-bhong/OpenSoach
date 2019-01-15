import { PlatformHelperAPI } from "./platform-helper-api";

export class PlatformAndroidHelper implements PlatformHelperAPI {
   getRandomUUID(): string {
        return java.util.UUID.randomUUID().toString();
    }
}