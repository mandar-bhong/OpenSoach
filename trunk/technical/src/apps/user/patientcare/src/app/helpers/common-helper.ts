import { BuildMode } from "../app-constants";

declare var process: any;
export class CommonHelper {
    static getEnvironmentVars(key: string): string {    
        if (typeof process !== 'undefined' && process && process.env) {
            return process.env[key];
        } else {
            return BuildMode.DEVELOPMENT
        }
    }
}