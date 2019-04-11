import {
    CorporateAddRequest, CorporateFilterRequest, CorpDetailsResponse,
    CorporateDataListingItemResponse, CorporateUpdateRequest
} from '../api/corporate-models';

export class CorporateAddModel {
    corpname: string;
    corpmobileno: string;
    corpemailid: string;
    corplandlineno: string;

    copyTo(corporateAddRequest: CorporateAddRequest) {
        corporateAddRequest.corpname = this.corpname;
        corporateAddRequest.corpmobileno = this.corpmobileno;
        corporateAddRequest.corpemailid = this.corpemailid;
        corporateAddRequest.corplandlineno = this.corplandlineno;
    }
}

export class CorporateFilterModel {
    corpname: string;
    corpmobileno: string;
    corpemailid: string;

    copyTo(corporateFilterRequest: CorporateFilterRequest) {
        corporateFilterRequest.corpname = this.corpname;
        corporateFilterRequest.corpmobileno = this.corpmobileno;
        corporateFilterRequest.corpemailid = this.corpemailid;
    }
}

export class CorporateDetailsModel {
    corpid: number;
    corpname: string;
    corpmobileno: string;
    corpemailid: string;
    corplandlineno: string;
    copyToUpdateRequest(corporateUpdateRequest: CorporateUpdateRequest) {
        corporateUpdateRequest.corpid = this.corpid;
        corporateUpdateRequest.corpname = this.corpname;
        corporateUpdateRequest.corpmobileno = this.corpmobileno;
        corporateUpdateRequest.corpemailid = this.corpemailid;
        corporateUpdateRequest.corplandlineno = this.corplandlineno;
    }

    copyToAddRequest(corporateAddRequest: CorporateAddRequest) {
        corporateAddRequest.corpname = this.corpname;
        corporateAddRequest.corpmobileno = this.corpmobileno;
        corporateAddRequest.corpemailid = this.corpemailid;
        corporateAddRequest.corplandlineno = this.corplandlineno;
    }

    copyFrom(corpDetailsResponse: CorpDetailsResponse) {
        this.corpid = corpDetailsResponse.corpid;
        this.corpname = corpDetailsResponse.corpname;
        this.corpmobileno = corpDetailsResponse.corpmobileno;
        this.corpemailid = corpDetailsResponse.corpemailid;
        this.corplandlineno = corpDetailsResponse.corplandlineno;
    }
}
