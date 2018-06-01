export class DeviceFilterRequest {
    serialno: string;
    devname: string;
}

export class DeviceDataListResponse {
    devid: number;
    serialno: string;
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

export class DeviceListItemResponse {
    devid: number;
    devname: string;
    serialno: string;
}