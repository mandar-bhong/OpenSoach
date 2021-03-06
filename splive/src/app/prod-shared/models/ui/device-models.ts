import { DeviceDetailsResponse, DeviceFilterRequest, DeviceDetailsUpdateRequest } from '../api/device-models';

export class DeviceFilterModel {
    serialno: string;
    devname: string;
    connectionstate: string;

    copyTo(deviceFilterRequest: DeviceFilterRequest) {
        deviceFilterRequest.serialno = this.serialno;
        deviceFilterRequest.devname = this.devname;
        deviceFilterRequest.connectionstate = this.connectionstate;
    }
}

export class DeviceDetailsModel {
    devid: number;
    devname: string;
    copyFrom(deviceDetailsResponse: DeviceDetailsResponse) {
        this.devid = deviceDetailsResponse.devid;
        this.devname = deviceDetailsResponse.devname;
    }
    copyTo(deivceDetailsUpdateRequest: DeviceDetailsUpdateRequest) {
        deivceDetailsUpdateRequest.devid = this.devid;
        deivceDetailsUpdateRequest.devname = this.devname;
    }
}
