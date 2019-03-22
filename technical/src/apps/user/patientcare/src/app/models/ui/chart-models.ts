import { ScheduleDatastoreModel } from "../db/schedule-model.js";

export class ChartListViewModel {
    dbmodel: any;
    // serverdbmodal: any;
    conf: ConfigData;
}

export class ChartDBModel {
    conf_type_code: string;
    conf: any;
    uuid: string;
    admission_uuid: string;

}

export class MonitorChartModel {
    name: string;
    foodInst: number;
    frequency: number;
    duration: number;
    startDate: Date;
    intervalHrs: number;
    startTime: string;
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
    intervalHrs: number;
    startDate: Date;
    duration: number;
    startTime: string;
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
    startDate: string;
    intervalHrs: number;
    startTime: string;
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
}

export class ConfigData {
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
    mornFreqQuantity: number;
}

export class AftrnFreqInfo {
    freqAftrn: boolean;
    aftrnFreqQuantity: number;
}

export class NightFreqInfo {
    freqNight: boolean;
    nightFreqQuantity: number;
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
    intervalHrs: number;
    startDate: string;
    duration: number;
    startTime: string;
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
    startDate: string;
    desc: string;
    remark: string;
    name: string;
}