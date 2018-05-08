export class CorporateShortDataResponse {
    corpid: number;
    corpname: string;
}
export class CorporateAddRequest {
    corpname: string;
    corpmobileno: string;
    corpemailid: string;
    corplandlineno: string;
}

export class CorporateFilterRequest {
    corpname: string;
    corpmobileno: string;
    corpemailid: string;
}

export class CorporateDataListingItemResponse {
    corpid: number;
    corpname: string;
    corpmobileno: string;
    corpemailid: string;
    corplandlineno: string;
    createdon: Date;
    updatedon: Date;
}
export class CorporateUpdateRequest {
    corpid: number;
    corpname: string;
    corpmobileno: string;
    corpemailid: string;
    corplandlineno: string;
}
export class CorpDetailsResponse {
    corpid: number;
    corpname: string;
    corpmobileno: string;
    corpemailid: string;
    corplandlineno: string;
    createdon: Date;
    updatedon: Date;
}
