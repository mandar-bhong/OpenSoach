import {
    DeviceFilterRequest, DeviceAddRequest, DeviceAddDetailsRequest, DeviceDetailsResponse,
    DeviceAssociateProductRequest, DeviceAssociateProductListItemResponse, DeviceMasterUpdateRequest, DeviceMasterUpdateResponse
} from '../api/device-models';
import { DEVICE_STATE } from '../../../../shared/app-common-constants';
import { CustomerAssociateProductListItemResponse } from '../../../../spl/app/models/api/customer-models';
export class DeviceFilterModel {
    serialno: string;
    custid: number;
    custname: string;

    copyTo(deviceFilterRequest: DeviceFilterRequest) {
        deviceFilterRequest.serialno = this.serialno;
        deviceFilterRequest.custid = this.custid;
        deviceFilterRequest.custname = this.custname;
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
    copyToUpdateRequest(deviceMasterUpdateRequest: DeviceMasterUpdateRequest) {
        deviceMasterUpdateRequest.devid = this.devid;
        deviceMasterUpdateRequest.devstate = this.devstate;
    }
    copyFrom(customerMasterResponse: DeviceMasterUpdateResponse) {
        this.custid = customerMasterResponse.custid;
        this.devid = customerMasterResponse.devid;
        this.serialno = customerMasterResponse.serialno;
        this.devstate = customerMasterResponse.devstate;
    }
}

export class DeviceDetailsModel {
    devid: number;
    make: string;
    technology: string;
    techversion: string;
    shortdesc: string;
    copyTo(deviceAddDetailsRequest: DeviceAddDetailsRequest) {
        deviceAddDetailsRequest.devid = this.devid;
        deviceAddDetailsRequest.make = this.make;
        deviceAddDetailsRequest.technology = this.technology;
        deviceAddDetailsRequest.techversion = this.techversion;
        deviceAddDetailsRequest.shortdesc = this.shortdesc;
    }
    copyFrom(deviceDetailsResponse: DeviceDetailsResponse) {
        this.devid = deviceDetailsResponse.devid;
        this.make = deviceDetailsResponse.make;
        this.technology = deviceDetailsResponse.technology;
        this.techversion = deviceDetailsResponse.techversion;
        this.shortdesc = deviceDetailsResponse.shortdesc;
    }
}
export class DeviceAssociateProductModel {
    cpmid: number;
    cpm: CustomerAssociateProductListItemResponse;
    devid: number;
    custname: string;
    prodcode: string;
    custid: number;

    copyToAddRequest(request: DeviceAssociateProductRequest) {
        request.cpmid = this.cpm.cpmid;
        request.devid = this.devid;
    }

    copyFrom(details: DeviceAssociateProductListItemResponse) {
        this.custname = details.custname;
        this.prodcode = details.prodcode;
    }
}
