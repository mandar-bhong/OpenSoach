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

// For patient dashbored summery display about hospitalized and discharged
//task completion trend
export class TaskTrendRequest {
    spid: number;
    startdate: Date;
    enddate: Date;
}
//task completion trend
export class TaskTrendResponse {
    month: number;
    year: number;
    hospitalized: number;
    discharged: number;
}

// For patient dashbored summery display about hospitalized and discharged
// request
export class PatientSummaryRequest {
    spid: number;
    startdate: Date;
    enddate: Date;
}

//response
export class PatientSummaryResponse {
    admitted: number;
    discharged: number;
}


