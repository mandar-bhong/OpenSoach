import { ServicepointListResponse } from '../../../../prod-shared/models/api/servicepoint-models';

export class ReportContainerModel {
    selecteddateoption: string;
    selectedsp:ServicepointListResponse;
    splist: ServicepointListResponse[];
    startdate: Date;
    enddate:Date;
}