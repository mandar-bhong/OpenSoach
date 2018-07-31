export class DeviceAddRequest {
    devid: number;
    custid: number;
    serialno: string;
    devstate: number;
}
export class DeviceMasterUpdateRequest {
    devid: number;
    devstate: number;
}
export class DeviceMasterUpdateResponse {
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
    devstate: number;
}

export class DeviceDataListResponse {
    devid: number;
    custid: number;
    custname: string;
    serialno: string;
    devstate: number;
    devstatesince: Date;
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
