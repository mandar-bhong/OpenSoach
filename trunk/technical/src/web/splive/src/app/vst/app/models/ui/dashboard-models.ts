import {
    ComplaintSummaryResponse,
    ComplaintTrendResponse,
    DeviceSummaryResponse,
    FeedbackSummaryResponse,
    FeedbackTrendResponse,
    ServicePointSummaryResponse,
    TaskSummaryResponse,
    TaskTrendResponse,
    SnapShotRequest,
    SnapShotResponse,
    TimeRequest,
    TimeResponse,
    ServiceTimeAvrResponse,
    VehicleServiceTrendMontlyResponse,
    VehicleServiceTrendWeeklyResponse,
} from '../api/dashboard-models';
import { SNAPSHOT_STATE } from '../../app-constants';

export class DeviceSummaryModel {
    total: number;
    online: number;
    offline: number;
    onlinepercentage: number;
    offlinepercentage: number;

    copyFrom(response: DeviceSummaryResponse) {
        this.total = response.total;
        this.online = response.online;
        this.offline = response.offline;
        this.onlinepercentage = 0;
        this.offlinepercentage = 0;
        if (this.total > 0) {
            this.onlinepercentage = (this.online / this.total) * 100;
            this.offlinepercentage = (this.offline / this.total) * 100;
        }
    }
}

export class ServicePointSummaryModel {
    total: number;
    inuse: number;
    inusepercentage: number;

    copyFrom(response: ServicePointSummaryResponse) {
        this.total = response.total;
        this.inuse = response.inuse;
        this.inusepercentage = 0;
        if (this.total > 0) {
            this.inusepercentage = (this.inuse / this.total) * 100;
        }
    }
}

export class ComplaintSummaryModel {
    total: number;
    open: number;
    closed: number;
    inprogress: number;
    active: number;
    activepercentage: number;

    copyFrom(response: ComplaintSummaryResponse) {
        this.open = response.open;
        this.closed = response.closed;
        this.inprogress = response.inprogress;
        this.total = this.open + this.closed + this.inprogress;
        this.active = this.open + this.inprogress;
        this.activepercentage = 0;
        if (this.total > 0) {
            this.activepercentage = (this.active / this.total) * 100;
        }
    }
}

export class FeedbackSummaryModel {
    rating1: number;
    rating2: number;
    rating3: number;
    rating4: number;
    rating5: number;
    rating1percentage: number;
    rating2percentage: number;
    rating3percentage: number;
    rating4percentage: number;
    rating5percentage: number;
    total: number;
    averagerating: number;

    copyFrom(response: FeedbackSummaryResponse) {
        this.rating1 = response.rating1;
        this.rating2 = response.rating2;
        this.rating3 = response.rating3;
        this.rating4 = response.rating4;
        this.rating5 = response.rating5;

        this.total = this.rating1 + this.rating2 + this.rating3 + this.rating4 + this.rating5;
        this.averagerating = 0;
        this.rating1percentage = 0;
        this.rating2percentage = 0;
        this.rating3percentage = 0;
        this.rating4percentage = 0;
        this.rating5percentage = 0;

        if (this.total > 0) {
            this.averagerating = (this.rating1 * 1
                + this.rating2 * 2
                + this.rating3 * 3
                + this.rating4 * 4
                + this.rating5 * 5) / this.total;

            this.rating1percentage = (this.rating1 / this.total) * 100;
            this.rating2percentage = (this.rating2 / this.total) * 100;
            this.rating3percentage = (this.rating3 / this.total) * 100;
            this.rating4percentage = (this.rating4 / this.total) * 100;
            this.rating5percentage = (this.rating5 / this.total) * 100;
        }
    }
}

export class TaskSummaryModel {
    ontime: number;
    delayed: number;
    ontimepercentage: number;
    delayedpercentage: number;
    total: number;

    copyFrom(response: TaskSummaryResponse) {
        this.ontime = response.ontime;
        this.delayed = response.delayed;

        this.total = this.ontime + this.delayed;
        this.ontimepercentage = 0;
        this.delayedpercentage = 0;
        if (this.total > 0) {
            this.ontimepercentage = (this.ontime / this.total) * 100;
            this.delayedpercentage = (this.delayed / this.total) * 100;
        }
    }
}

export class FeedbackTrendModel {
    month: number;
    year: number;
    rating1: number;
    rating2: number;
    rating3: number;
    rating4: number;
    rating5: number;
    total: number;
    averagerating: number;

    copyFrom(response: FeedbackTrendResponse) {
        this.rating1 = response.rating1;
        this.rating2 = response.rating2;
        this.rating3 = response.rating3;
        this.rating4 = response.rating4;
        this.rating5 = response.rating5;
        this.month = response.month;
        this.year = response.year;

        this.total = this.rating1 + this.rating2 + this.rating3 + this.rating4 + this.rating5;
        this.averagerating = 0;

        if (this.total > 0) {
            this.averagerating = (this.rating1 * 1
                + this.rating2 * 2
                + this.rating3 * 3
                + this.rating4 * 4
                + this.rating5 * 5) / this.total;
        }
    }
}

export class TrendChartPerMonthXaxis {
    year: number;
    month: number;
}

export class TaskTrendModel {
    month: number;
    year: number;
    ontime: number;
    delayed: number;

    copyFrom(response: TaskTrendResponse) {
        this.ontime = response.ontime;
        this.delayed = response.delayed;

        this.month = response.month;
        this.year = response.year;
    }
}

export class ComplaintTrendModel {
    month: number;
    year: number;
    open: number;
    closed: number;
    inprogress: number;

    copyFrom(response: ComplaintTrendResponse) {
        this.open = response.open;
        this.inprogress = response.inprogress;
        this.closed = response.closed;

        this.month = response.month;
        this.year = response.year;
    }
}

export class SnapshotModel {
    startdate: Date;
    enddate: Date;
    lastactiontime: Date;
    count: number;
    status: SNAPSHOT_STATE;
    copyTo(snapShotRequest: SnapShotRequest) {
        snapShotRequest.startdate = this.startdate;
        snapShotRequest.enddate = this.enddate;
    }
    copyFrom(response: SnapShotResponse) {
        this.lastactiontime = response.lastactiontime;
        this.count = response.count;
        this.status = response.status;
    }
}
export class TimeModel {
    startdate: Date;
    enddate: Date;
    waittime: number;
    jobcreationtime: number;
    jobexetime: number;
    deliverytime: number;
    copyTo(timeRequest: TimeRequest) {
        timeRequest.startdate = this.startdate;
        timeRequest.enddate = this.enddate;
    }
    copyFrom(response: TimeResponse) {
        this.waittime = response.waittime;
        this.jobcreationtime = response.jobcreationtime;
        this.jobexetime = response.jobexetime;
        this.deliverytime = response.deliverytime;
    }
}

export class SeriveTimeAvrModel {
    month: number;
    year: number;
    waittime: number;
    jobcreationtime: number;
    jobexetime: number;
    deliverytime: number;

    copyFrom(response: ServiceTimeAvrResponse) {
        this.waittime = response.waittime;
        this.jobcreationtime = response.jobcreationtime;
        this.jobexetime = response.jobexetime;
        this.deliverytime = response.deliverytime;
        this.month = response.month;
        this.year = response.year;
    }
}


export class VehicleServiceTrendMonthlyModel {
    month: number;
    year: number;
    vehicleserviced: number;

    copyFrom(response: VehicleServiceTrendMontlyResponse) {
        this.vehicleserviced = response.vehicleserviced;
        this.month = response.month;
        this.year = response.year;
    }
}
export class VehicleServiceTrendWeeklyModel {
    vehicleserviced: number;
    servicedate: Date;

    copyFrom(response: VehicleServiceTrendWeeklyResponse) {
        this.vehicleserviced = response.vehicleserviced;
        this.servicedate = response.servicedate;
    }
}
export class VehicleChartPerWeeklyYaxis {
    year: number;
    month: number;
    weekly: number;
}
