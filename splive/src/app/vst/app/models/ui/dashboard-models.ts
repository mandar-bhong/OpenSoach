import {
    ComplaintSummaryResponse,
    ComplaintTrendResponse,
    DeviceSummaryResponse,
    FeedbackSummaryResponse,
    FeedbackTrendResponse,
    ServicePointSummaryResponse,
    TaskSummaryResponse,
    TaskTrendResponse,
} from '../api/dashboard-models';

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
    tokantime = 4;
    tokancount = 3;
    createdtime = 1;
    createdcount = 4;
    inprogtime = 2;
    inprogconut = 5;
    compltime = 7;
    complcount = 8;
    vehicletime = 4;
    vehiclecount = 6;
}
