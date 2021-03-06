export class DeviceFilterRequest {
    serialno: string;
    devname: string;
    connectionstate: string;
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
    createdon: Date;
    updatedon: Date;
}

export class DeviceListItemResponse {
    devid: number;
    devname: string;
    serialno: string;
}
export class DeviceDetailsUpdateRequest {
    devid: number;
    devname: string;
}