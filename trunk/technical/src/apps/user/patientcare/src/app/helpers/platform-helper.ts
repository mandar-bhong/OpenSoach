import { PlatformHelperAPI } from "~/app/helpers/platform-helper-api";
import * as application from "tns-core-modules/application";
import { PlatformAndroidHelper } from "./platform-android-helper";
import { PlatformiOSHelper } from "./platform-ios-helper";

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