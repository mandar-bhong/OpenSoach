import { ScheduleDatastoreModel } from "../db/schedule-model.js";

export class ChartListViewModel {
    dbmodel: any;
    // serverdbmodal: any;
    conf: ConfigData;
    expanded: boolean;
}

export class ChartDBModel {
    conf_type_code: string;
    conf: any;
    uuid: string;
    admission_uuid: string;
    start_date: string;

}

export class MonitorChartModel {
    name: string;
    foodInst: number;
    frequency: number;
    duration: number;
    startDate: Date;
    interval: number;
    startTime: number;
    endTime: string;
    specificTime: string;
    specificTimes: Array<any>;
    desc: string;
    endDate: Date
    numberofTimes: number;
    remark: string;
}

export class MedChartModel {
    name: string;
    quantity: number;
    foodInst: number;
    frequency: number;
    mornFreqInfo: MornFreqInfo;
    aftrnFreqInfo: AftrnFreqInfo;
    nightFreqInfo: NightFreqInfo
    interval: number;
    startDate: Date;
    duration: number;
    startTime: number;
    desc: string;
    endDate: Date
    numberofTimes: number;
    medicinetype: string;
    splinstruction: string;
    remark: string;
}

export class IntakeChartModel {
    name: string;
    quantity: string;
    frequency: number;
    duration: number;
    startDate: Date;
    interval: number;
    startTime: number;
    specificTime: string;
    specificTimes: Array<any>;
    desc: string;
    endDate: Date
    numberofTimes: number;
    splinstruction: string;
    remark: string;
    intakeType: string;
}

export class Schedulardata {
    data: ScheduleDatastoreModel;
    uuid: string;
    admission_uuid: string;
    conf_type_code: string;
    conf: SchedularConfigData;
    updated_by: number;
    start_date: string;
}

export class ConfigData {
    name: string;
    desc: string;
    freqMorn: boolean;
    freqAftrn: boolean;
    freqNight: boolean;
    quantity: number;
    startDate: Date;
    frequency: number;
    foodInst: number;
    duration: string;
    interval: number;
    startTime: number;
    medicinetype: string;
    splinstruction: string;
    remark: string;

}
export class ActionListModel {
    uuid: string;
    admission_id: number;
    chart_conf_id: number;
    scheduled_time: string;
}

export class MornFreqInfo {
    freqMorn: boolean;
    mornFreqQuantity: string;
}

export class AftrnFreqInfo {
    freqAftrn: boolean;
    aftrnFreqQuantity: string;
}

export class NightFreqInfo {
    freqNight: boolean;
    nightFreqQuantity: string;
}
export class
    SchedularConfigData {
    name: string;
    quantity: number;
    foodInst: number;
    frequency: number;
    mornFreqInfo: MornFreqInfo;
    aftrnFreqInfo: AftrnFreqInfo;
    nightFreqInfo: NightFreqInfo
    interval: number;
    startDate: Date;
    duration: number;
    startTime: number;
    desc: string;
    endTime: string;
    specificTime: string;
    specificTimes: Array<any>;
    numberofTimes: number;
    endDate: Date
    splinstruction: string;
    remark: string;
}
export class FrequencyValues {
    name: string;
    value: number;
}
export class PickerValues {
    name: string;
    value: string;
}
export class OutputChartModel {
    outputType: string;
    duration: number;
    startDate: Date;
    desc: string;
    remark: string;
    name: string;
    //frequency: number;
}