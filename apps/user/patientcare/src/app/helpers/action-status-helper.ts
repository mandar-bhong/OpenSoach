import * as moment from "moment";
import { ACTION_DELAYED_AFTER, ACTION_NEEDS_ATTENTION, ACTION_MISSED_WINDOW, ACTION_STATUS, ACTION_FUTURE_AFTER } from "../app-constants";

export class ActionStatusHelper {
    static getActionStatus(actionTime: Date): ACTION_STATUS {
        let status = ACTION_STATUS.NONE
        if (actionTime) {
            const diff = moment(actionTime).diff(moment(), 'minutes', true);
            if (diff < -ACTION_MISSED_WINDOW) {
                status = ACTION_STATUS.MISSED
            }
            if (diff < -ACTION_DELAYED_AFTER) {
                status = ACTION_STATUS.ACTIVE_DELAYED;
            } else if (diff <= ACTION_NEEDS_ATTENTION) {
                status = ACTION_STATUS.ACTIVE_NEEDS_ATTENTION;
            }
            else if (diff <= ACTION_FUTURE_AFTER) {
                status = ACTION_STATUS.ACTIVE_NORMAL;
            }
            else {
                status = ACTION_STATUS.ACTIVE_FUTURE;
            }
        }

        return status;
    }
}