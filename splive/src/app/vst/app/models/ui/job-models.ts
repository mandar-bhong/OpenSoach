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

    vehicleid: number;
    vehicleno: string;
    details: VehicleFullDetails;
    mobileno: string;
    name: string;
    kms: string;
    petrol: string;
    status: number;
    fopcode: number;
    txndata: any;
    txndate: Date;
    tokenid: number;
    transactions: JobTransaction[];

    copyFromDetails(vehicleDetailsResponse: VehicleDetailsResponse) {
        this.vehicleno = vehicleDetailsResponse.vehicleno;
        this.mobileno = vehicleDetailsResponse.mobileno;
        this.kms = vehicleDetailsResponse.kms;
        this.petrol = vehicleDetailsResponse.petrol;
        this.name = vehicleDetailsResponse.name;
        // this.details = new VehicleFullDetails();
        // Object.assign(this.details, JSON.parse(vehicleDetailsResponse.details));
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

