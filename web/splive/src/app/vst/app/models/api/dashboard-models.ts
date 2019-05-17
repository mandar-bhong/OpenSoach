export class DeviceSummaryResponse {
    total: number;
    online: number;
    offline: number;
}

export class ServicePointSummaryResponse {
    total: number;
    active: number;
    inuse: number;
}

export class ComplaintSummaryRequest {
    spid: number;
    startdate: Date;
    enddate: Date;
}

export class ComplaintSummaryResponse {
    open: number;
    closed: number;
    inprogress: number;
}

export class FeedbackSummaryRequest {
    spid: number;
    startdate: Date;
    enddate: Date;
}

export class FeedbackSummaryResponse {
    rating1: number;
    rating2: number;
    rating3: number;
    rating4: number;
    rating5: number;
}

export class TaskSummaryRequest {
    spid: number;
    startdate: Date;
    enddate: Date;
}

export class TaskSummaryResponse {
    ontime: number;
    delayed: number;
}

export class FeedbackTrendRequest {
    spid: number;
    startdate: Date;
    enddate: Date;
}

export class FeedbackTrendResponse {
    month: number;
    year: number;
    rating1: number;
    rating2: number;
    rating3: number;
    rating4: number;
    rating5: number;
}

export class TaskTrendRequest {
    spid: number;
    startdate: Date;
    enddate: Date;
}

export class TaskTrendResponse {
    month: number;
    year: number;
    ontime: number;
    delayed: number;
}

export class ComplaintTrendRequest {
    spid: number;
    startdate: Date;
    enddate: Date;
}

export class ComplaintTrendResponse {
    month: number;
    year: number;
    open: number;
    closed: number;
    inprogress: number;
}

export class SnapShotRequest {
    startdate: Date;
    enddate: Date;
}
export class SnapShotResponse {
    lastactiontime: Date;
    count: number;
    status: number;
}
export class TimeRequest {
    startdate: Date;
    enddate: Date;
}
export class TimeResponse {
    waittime: number;
    jobcreationtime: number;
    jobexetime: number;
    deliverytime: number;
}

export class ServiceTimeAvrResponse {
    month: number;
    year: number;
    waittime: number;
    jobcreationtime: number;
    jobexetime: number;
    deliverytime: number;
}
export class ServiceTimeAvrRequest {
    // spid: number;
    startdate: Date;
    enddate: Date;
}

export class VehicleServiceTrendMontlyResponse {
    month: number;
    year: number;
    vehicleserviced: number;
}
export class VehicleServiceTrendMontlyRequest {
    // spid: number;
    startdate: Date;
    enddate: Date;
}
export class VehicleServiceTrendWeeklyRequest {
    startdate: Date;
    enddate: Date;
}
export class VehicleServiceTrendWeeklyResponse {
    // days: string;
    servicedate: Date;
    vehicleserviced: number;
}

