export class OperatorFiltrRequest {
    fopid: number;
    fopname: string;
    mobileno: string;
    emailid: string;
    shortdesc: string;
    fopstate: number;
    foparea: number;
}
export class OperatorDataListResponse {
    fopid: number;
    fopcode: string;
    fopname: string;
    mobileno: string;
    emailid: string;
    shortdesc: string;
    fopstate: number;
    foparea: number;
    createdon: Date;
    updatedon: Date;
}
export class OperatorAddRequest {
    fopcode: string;
    fopname: string;
    mobileno: string;
    emailid: string;
    shortdesc: string;
    fopstate: number;
    foparea: number;
}
export class OperatorUpdateRequest {
    fopid: number;
    fopname: string;
    mobileno: string;
    emailid: string;
    shortdesc: string;
    fopstate: number;
    foparea: number;
    fopcode: string;
}
export class OperatorDetailsResponse {
    fopid: number;
    fopcode: string;
    fopname: string;
    mobileno: string;
    emailid: string;
    shortdesc: string;
    fopstate: number;
    foparea: number;
    createdon: Date;
    updatedon: Date;
}
export class OperatorAssociateListResponse {
    fopid: number;
    spid: number;
}
export class OperatorServicepointListResponse {
    spid: number;
    spname: string;
}
export class OperatorServicePointListModel {
    spid: number;
    spname: string;
    ischecked: boolean;
}
