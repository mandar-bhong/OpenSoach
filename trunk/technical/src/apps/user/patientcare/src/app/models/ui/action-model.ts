import { Time } from "@angular/common/src/i18n/locale_data_api";

export class SchedularData {
    uuid: string;
    admission_uuid: string;
    conf_type_code: string;
    conf: SchedularConfigData;
}

export class SchedularConfigData {
    name: string;
    desc: string;
    freqMorn: boolean;
    freqAftrn: boolean;
    freqNight: boolean;
    quantity: number;
    startDate: string;
    frequency: number;
    foodInst: number;
    duration: string;
    intervalHrs: string;
    startTime: string;
}
export class ActionList {
    uuid: string;
    admission_id: string;
    chart_conf_id: number;
    exec_time: Date;
}
export class ActionItems {
    dateAction: Date;
    dayAction: DayActionValue[];
}
export enum ProcessTime {
    DayStartTime = "8.00",
    DayEndTime = "11.00",
    foodInstAfterMeal = 1,
    foodInstBeforeMeal = 0,
}


export class TimeConstants {
    morningTime: any;
    afternoonTime: any
    nightTime: any;
}

export class DayActionValue {
    time: number;
}

export enum BeforeMealTimeInMinutes {
    MorningBeforeMeal = 540,
    AfternoonbeBeforeMeal = 690,
    NightBeforemeal = 1170,
}
export enum AfterMealTimeInMinutes {
    MorningAfterMeal = 570,
    AfternoonAfterMeal = 750,
    NightAfteremeal = 1230,
}
export enum BeforeMealTime {
    MorningBeforeMeal = "9.00 AM",
    AfternoonbeBeforeMeal = "11.30 AM",
    NightBeforemeal = "7.30 PM",
}
export enum AfterMealTime {
    MorningAfterMeal = "9.30 AM",
    AfternoonAfterMeal = "12.30 PM",
    NightAfteremeal = "8.00 PM",
}
export enum DayTimes {
    dayStartTime = 0, // minutes
    dayEndTime = 1440 // minutes
}
export enum Medicinefrequency {
    xTimesInDay = 0,
    AfterXTimeInterval = 1
}
export enum Monitorfrequency {
    AfterXTimeInterval = 0,
    speficicTime = 1
}
export enum Intakefrequency {
    AfterXTimeInterval = 0,
    speficicTime = 1
}
export enum dayTime {
    Morning = 'morning',
    Afternoon = 'Afternoon',
    Night = 'night'
}

export class AfterXtimeIntervl {
    dateAction: Date;
    time: number;
}