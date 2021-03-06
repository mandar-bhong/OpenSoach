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


