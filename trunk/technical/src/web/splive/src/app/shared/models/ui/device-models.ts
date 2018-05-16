import { DeviceDetailsResponse, DeviceFilterRequest } from '../api/device-models';

export class DeviceFilterModel {
    serialno: string;
    devname: string;

    copyTo(deviceFilterRequest: DeviceFilterRequest) {
        deviceFilterRequest.serialno = this.serialno;
        deviceFilterRequest.devname = this.devname;
    }
}

export class DeviceDetailsModel {
    devid: number;
    devname: string;
    make: string;
    technology: string;
    techversion: string;
    shortdesc: string;
    copyFrom(deviceDetailsResponse: DeviceDetailsResponse) {
        this.devid = deviceDetailsResponse.devid;
        this.devname = deviceDetailsResponse.devname;
        this.make = deviceDetailsResponse.make;
        this.technology = deviceDetailsResponse.technology;
        this.techversion = deviceDetailsResponse.techversion;
        this.shortdesc = deviceDetailsResponse.shortdesc;
    }
}
