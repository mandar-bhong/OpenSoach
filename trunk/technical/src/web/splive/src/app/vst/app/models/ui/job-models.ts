import {
    JobFiltrRequest, JobDetailsDataListResponse,
    VehicleDetailsResponse, OwnerResponse,
    VehicleFullDetails, JobDetailslistResponse, JobTxndata
} from '../api/job-models';

export class JobFilterModel {
    vehicleno: number;
    status: number;
    startdate: Date;
    enddate: Date;

    copyTo(jobFiltrRequest: JobFiltrRequest) {
        jobFiltrRequest.vehicleno = this.vehicleno;
        jobFiltrRequest.status = this.status;
        jobFiltrRequest.startdate = this.startdate;
        jobFiltrRequest.enddate = this.enddate;
    }
}

export class JobDetailsModel {
    time: string;
    servicepersonnel: string;
    acitivty: string;
    notes: string;
    price: string;
    vehicleid: number;
    vehicleno: number;
    details: VehicleFullDetails;
    // details: {};
    // ownerdetails: OwnerResponse;
    // vehicleetails: VehicleResponse;
    status: number;
    fopcode: number;
    txndata: any;
    txndate: Date;
    tokenid: number;
    transactions: JobTransaction[];

    copyFromDetails(vehicleDetailsResponse: VehicleDetailsResponse) {
        this.vehicleno = vehicleDetailsResponse.vehicleno;
        this.details = new VehicleFullDetails();
        Object.assign(this.details, JSON.parse(vehicleDetailsResponse.details));
    }
    copyFormList(jobDetailslistResponse: JobDetailslistResponse[]) {
        this.transactions = [];
        jobDetailslistResponse.forEach(item => {
            const transaction = new JobTransaction();
            transaction.copyFrom(item);
            this.transactions.push(transaction);
            // console.log('this.transactions', this.transactions);
        });
    }


}

export class JobTransaction {
    status: number;
    fopcode: number;
    txndata: any;
    txndate: Date;
    // tokenid: number;

    copyFrom(jobDetailslistResponse: JobDetailslistResponse) {
        this.status = jobDetailslistResponse.status;
        this.fopcode = jobDetailslistResponse.fopcode;
        this.txndate = jobDetailslistResponse.txndate;
        this.txndata = {};
        Object.assign(this.txndata, JSON.parse(jobDetailslistResponse.txndata));
        // console.log('this.txndata', this.txndata);
    }
}

