import { Injectable } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import { ACTION_STATUS, ACTION_MISSED_WINDOW, ACTION_FUTURE_AFTER } from "~/app/app-constants";
import * as moment from "moment";
import { ActionStatusHelper } from "~/app/helpers/action-status-helper";
import { NextActionTimes } from "~/app/models/ui/next-action-times";
import { setInterval, clearInterval } from "tns-core-modules/timer";
import { WorkerService } from "../worker.service";
import { Subject } from "rxjs";

@Injectable()
export class NextActionService {

    nextActionTimesMap = new Map<string, NextActionTimes>();
    timerReference: number;
    nextActionMapChanged = new Subject<{ admission_uuid: string, nextActionTimes: NextActionTimes }>();
    constructor(private database: DatabaseService,
        private workerService: WorkerService) {

        this.workerService.patientAdmissionDataReceivedSubject.subscribe((dataStoreModel) => {
            this.getNextActionsForAdmission(dataStoreModel.uuid);
        });
        this.workerService.scheduleDataReceivedSubject.subscribe((dataStoreModel) => {
            this.getNextActionsForAdmission(dataStoreModel.admission_uuid);
        });
        this.workerService.actionDataReceivedSubject.subscribe((dataStoreModel) => {
            this.getNextActionsForAdmission(dataStoreModel.admission_uuid);
        });
        this.workerService.actionTxnDataReceivedSubject.subscribe((dataStoreModel) => {
            this.getNextActionsForAdmission(dataStoreModel.admission_uuid);
        });

        this.startTimer();
    }

    getNextActionsForAllPatients(): Promise<Map<string, NextActionTimes>> {
        return new Promise((resolve, reject) => {
            const currentDate = new Date().toISOString();
            this.database.selectByID("action_tbl_next_actions_all", [currentDate, currentDate]).then(
                (val) => {
                    console.log('action_tbl_next_actions_all', val);
                    val.forEach(element => {
                        const nextActionTimes = this.nextActionTimesMap.get(element.admission_uuid) || new NextActionTimes();
                        nextActionTimes.times.push(new Date(element.scheduled_time));
                        this.nextActionTimesMap.set(element.admission_uuid, nextActionTimes);
                        this.setNextActionTimeAndStatus(nextActionTimes);
                    });
                    resolve(this.nextActionTimesMap);
                }, (error) => {
                    console.error(error);
                    reject(error);
                });
        });
    }

    getNextActionsForAdmission(admission_uuid: string) {
        const currentDate = new Date().toISOString();
        this.database.selectByID("action_tbl_next_actions_for_admission", [admission_uuid, currentDate, admission_uuid, currentDate]).then(
            (val) => {
                console.log('action_tbl_next_actions_for_admission', val);
                const nextActionTimes = new NextActionTimes();
                val.forEach(element => {
                    nextActionTimes.times.push(new Date(element.scheduled_time));
                    this.nextActionTimesMap.set(element.admission_uuid, nextActionTimes);
                    this.setNextActionTimeAndStatus(nextActionTimes);
                });

                this.nextActionMapChanged.next({ admission_uuid: admission_uuid, nextActionTimes: nextActionTimes });
            }, (error) => {
                console.error(error);
            });

    }

    setNextActionTimeAndStatus(nextActionTimes: NextActionTimes) {

        nextActionTimes.currentTime = null;
        nextActionTimes.status = ACTION_STATUS.NONE;
        if (nextActionTimes.times.length > 0
            && nextActionTimes.currentIndex < nextActionTimes.times.length) {
            const referenceTimeBefore = moment().subtract(ACTION_MISSED_WINDOW, 'minutes');
            const referenceTimeAfter = moment().add(ACTION_FUTURE_AFTER, 'minutes');
            let i = nextActionTimes.currentIndex || 0;
            for (i = 0; i < nextActionTimes.times.length; i++) {
                if (moment(nextActionTimes.times[i]).isBetween(referenceTimeBefore, referenceTimeAfter, null, '[]')) {
                    nextActionTimes.currentTime = nextActionTimes.times[i];
                    nextActionTimes.currentIndex = i;
                    nextActionTimes.status = ActionStatusHelper.getActionStatus(nextActionTimes.currentTime);
                    break;
                }
                else {
                    continue;
                }
            }
        }
    }

    updateStatus(actionTimes: NextActionTimes) {
        // if none  do nothing
        const status = ActionStatusHelper.getActionStatus(actionTimes.currentTime);
        if (status == ACTION_STATUS.MISSED) {
            // if the action is moved to missed, then the next action time to be considered.
            this.setNextActionTimeAndStatus(actionTimes);
        }
    }

    startTimer() {
        this.timerReference = setInterval(() => {
            console.log('timer ticked');
            this.nextActionTimesMap.forEach(item => {
                this.updateStatus(item);
            });
        }, 1000 * 60);
    }

    clearTimer() {
        clearInterval(this.timerReference);
    }


}