import { DeviceFilterRequest, DeviceAddRequest, DeviceAddDetailsRequest, DeviceDetailsResponse } from '../api/device-models';
import { DEVICE_STATE } from '../../../../shared/app-common-constants';

export class DeviceFilterModel {
    serialno: string;
    custid: number;
    custname: string;
    devname: string;

    copyTo(deviceFilterRequest: DeviceFilterRequest) {
        deviceFilterRequest.serialno = this.serialno;
        deviceFilterRequest.custid = this.custid;
        deviceFilterRequest.custname = this.custname;
        deviceFilterRequest.devname = this.devname;
    }
}

export class DeviceAddModel {
    devid: number;
    custid: number;
    serialno: string;
    devstate: DEVICE_STATE;

    copyTo(deviceAddRequest: DeviceAddRequest) {
        deviceAddRequest.devid = this.devid;
        deviceAddRequest.custid = this.custid;
        deviceAddRequest.serialno = this.serialno;
        deviceAddRequest.devstate = this.devstate;
    }
}

export class DeviceDetailsModel {
    devid: number;
    devname: string;
    make: string;
    technology: string;
    techversion: string;
    shortdesc: string;
    copyTo(deviceAddDetailsRequest: DeviceAddDetailsRequest) {
        deviceAddDetailsRequest.devid = this.devid;
        deviceAddDetailsRequest.devname = this.devname;
        deviceAddDetailsRequest.make = this.make;
        deviceAddDetailsRequest.technology = this.technology;
        deviceAddDetailsRequest.techversion = this.techversion;
        deviceAddDetailsRequest.shortdesc = this.shortdesc;
    }
    copyFrom(deviceDetailsResponse: DeviceDetailsResponse) {
        this.devid = deviceDetailsResponse.devid;
        this.devname = deviceDetailsResponse.devname;
        this.make = deviceDetailsResponse.make;
        this.technology = deviceDetailsResponse.technology;
        this.techversion = deviceDetailsResponse.techversion;
        this.shortdesc = deviceDetailsResponse.shortdesc;
    }
}
