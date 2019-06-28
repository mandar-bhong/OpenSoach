import { ACTION_STATUS } from "~/app/app-constants";

export class NextActionTimes {
    times: Date[] = []
    currentIndex: number = 0;
    status: ACTION_STATUS = ACTION_STATUS.NONE;
    currentTime: Date;
}