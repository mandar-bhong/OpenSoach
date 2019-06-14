import { BuildMode } from "../app-constants.js";

declare var process: any;
export class CommonHelper {
    static getEnvironmentVars(key: string): string {
        if (typeof process !== 'undefined' && process && process.env) {
            return process.env[key];
        } else {
            return BuildMode.DEVELOPMENT
        }
    }

    static getBuildEnvironment(): string {
        if (typeof process !== 'undefined' && process && process.env) {
            return process.env["buildmode"];
        } else {
            switch (process.env["buildmode"]) {
                case "prod"://this flag need to pass from command line
                    return BuildMode.PRODUCTION;
               
                default:
                    return BuildMode.DEVELOPMENT
            }

        }
    }
}