export class DeviceAddRequest {
    devid: number;
    custid: number;
    serialno: string;
    devstate: number;
}

export class DeviceAddDetailsRequest {
    devid: number;
    make: string;
    technology: string;
    techversion: string;
    shortdesc: string;
}

export class DeviceFilterRequest {
    serialno: string;
    custid: number;
    custname: string;
}

export class DeviceDataListResponse {
    devid: number;
    custid: number;
    custname: string;
    serialno: string;
    devstate: number;
    devstatesince: Date;
    connectionstate: number;
    connectionstatesince: Date;
    syncstate: number;
    syncstatesince: Date;
    batterylevel: number;
    batterylevelsince: Date;
    createdon: Date;
    updatedon: Date;
}
export class DeviceDetailsResponse {
    devid: number;
    make: string;
    technology: string;
    techversion: string;
    shortdesc: string;
    createdon: Date;
    updatedon: Date;
}

export class DeviceAssociateProductRequest {
    cpmid: number;
    devid: number;

}

export class DeviceAssociateProductListItemResponse {
    custname: string;
    prodcode: string;
}
