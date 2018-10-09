export class JobFiltrRequest {
    vehicleno: number;
    status: number;
    startdate: Date;
    enddate: Date;
}
export class JobDataListResponse {
    vehicleno: number;
    generatedon: Date;
    intime: Date;
    outtime: Date;
    token: number;
    state: number;
    tokenid: number;
    vehicleid: number;
    // mobileno: number;
}
export class JobDetailsDataListResponse {
    time: string;
    servicepersonnel: string;
    acitivty: string;
    notes: string;
    price: string;
}
export class JobDetailslistResponse {
    status: number;
    fopcode: number;
    txndata: string;
    txndate: Date;
}
export class JobTxndata {
    tasks: JobTrndatalist[];
    tokenid: number;
    vehicledetails: JobTrnVehicleResponse;
}
export class JobTrndatalist {
    cost: string;
    comment: string;
    taskname: string;
}
export class JobTrnVehicleResponse {
    petrol: string;
    kms: string;
}
export class StatusChangeRequest {
    state: number;
    tokenid: number;
    amount: number;
}
export class VehicleDetailsResponse {
    vehicleno: number;
    vehicleid: number;
    details: string;
}
export class VehicleFullDetails {
    ownerdetails: OwnerResponse;
    // vehicledetails: VehicleResponse;
}

export class OwnerResponse {
    firstname: string;
    lastname: string;
    mobileno: number;
}
