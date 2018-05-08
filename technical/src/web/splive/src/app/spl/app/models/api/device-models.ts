export class DeviceAddRequest {
    devid: number;
    custid: number;
    serialno: string;
    devstate: number;
}

export class DeviceAddDetailsRequest {
    devid: number;
    devname: string;
    make: string;
    technology: string;
    techversion: string;
    shortdesc: string;
}

export class DeviceFilterRequest {
    serialno: string;
    custid: number;
    custname: string;
    devname: string;
}

export class DeviceDataListResponse {
    devid: number;
    custid: number;
    custname: string;
    serialno: string;
    devstate: number;
    devstatesince: Date;
    devname: string;
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
    devname: string;
    make: string;
    technology: string;
    techversion: string;
    shortdesc: string;
    createdon: Date;
    updatedon: Date;
}

