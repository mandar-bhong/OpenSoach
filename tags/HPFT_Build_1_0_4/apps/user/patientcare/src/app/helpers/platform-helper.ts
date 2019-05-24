import { PlatformHelperAPI } from "~/app/helpers/platform-helper-api.js";
import * as application from "tns-core-modules/application";
import { PlatformAndroidHelper } from "./platform-android-helper.js";
import { PlatformiOSHelper } from "./platform-ios-helper.js";

export class PlatformHelper {
    public static API: PlatformHelperAPI;
    public static init() {
        if (application.android) {
            PlatformHelper.API = new PlatformAndroidHelper();
        } else if (application.ios) {
            PlatformHelper.API = new PlatformiOSHelper();
        }
    }
}