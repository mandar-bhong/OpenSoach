
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
    scheduled_time: Date;
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
    MorningBeforeMeal = 540, // minutes
    AfternoonbeBeforeMeal = 690, // minutes
    NightBeforemeal = 1170, // minutes
}
export enum AfterMealTimeInMinutes {
    MorningAfterMeal = 570, // minutes
    AfternoonAfterMeal = 750, // minutes
    NightAfteremeal = 1230, // minutes
}
export enum DayTimes {
    dayStartTime = 0, // minutes
    dayEndTime = 1440 // minutes
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

export class DataActionItem {
	uuid: string;
	admission_uuid: string;
	conf_type_code: string;
	schedule_uuid: string;
	scheduled_time: string;
	name: string;
	desc: string;
	status: number;
	document_uuid: string;
	doctors_orders: string;
	doctor_id: number;
	type: number;
	actionStatus: any;
	txn_state: number;
	txn_data: GetJsonModel;
    client_updated_at: string;
    fname: string;
    lname:string;
    value: BloodPressureValueModel;
}
export class GetJsonModel {
	comment: string;
    value: any;
 }
export class BloodPressureValueModel {
	systolic: number;
	diastolic: number;
}
