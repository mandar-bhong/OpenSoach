import { OPERATOR_AREA, OPERATOR_STATE } from '../../../shared/app-common-constants';
import {
    OperatorAddRequest,
    OperatorAssociateListResponse,
    OperatorDetailsResponse,
    OperatorFiltrRequest,
    OperatorServicePointListModel,
    OperatorUpdateRequest,
} from '../api/operator-models';

export class OperatorFilterModel {
    fopid: number;
    fopname: string;
    mobileno: string;
    emailid: string;
    shortdesc: string;
    fopstate: OPERATOR_STATE;
    foparea: OPERATOR_AREA;

    copyTo(operatorFiltrRequest: OperatorFiltrRequest) {
        operatorFiltrRequest.fopid = this.fopid;
        operatorFiltrRequest.fopname = this.fopname;
        operatorFiltrRequest.mobileno = this.mobileno;
        operatorFiltrRequest.emailid = this.emailid;
        operatorFiltrRequest.shortdesc = this.shortdesc;
        operatorFiltrRequest.fopstate = this.fopstate;
        operatorFiltrRequest.foparea = this.foparea;
    }
}

export class OperatorAddModel {
    fopid: number;
    fopcode: string;
    fopname: string;
    mobileno: string;
    emailid: string;
    shortdesc: string;
    fopstate: number;
    foparea: number;
    spid: number;
    spname: string;
    copyTo(operatorAddRequest: OperatorAddRequest) {
        operatorAddRequest.fopcode = this.fopcode;
        operatorAddRequest.fopname = this.fopname;
        operatorAddRequest.mobileno = this.mobileno;
        operatorAddRequest.emailid = this.emailid;
        operatorAddRequest.shortdesc = this.shortdesc;
        operatorAddRequest.fopstate = this.fopstate;
        operatorAddRequest.foparea = this.foparea;
    }
    copyToUpdateRequest(operatorUpdateRequest: OperatorUpdateRequest) {
        operatorUpdateRequest.fopid = this.fopid;
        operatorUpdateRequest.fopname = this.fopname;
        operatorUpdateRequest.mobileno = this.mobileno;
        operatorUpdateRequest.emailid = this.emailid;
        operatorUpdateRequest.shortdesc = this.shortdesc;
        operatorUpdateRequest.fopstate = this.fopstate;
        operatorUpdateRequest.foparea = this.foparea;
        operatorUpdateRequest.fopcode = this.fopcode;
    }
    copyFrom(operatorDetailsResponse: OperatorDetailsResponse) {
        this.fopid = operatorDetailsResponse.fopid;
        this.fopname = operatorDetailsResponse.fopname;
        this.mobileno = operatorDetailsResponse.mobileno;
        this.emailid = operatorDetailsResponse.emailid;
        this.shortdesc = operatorDetailsResponse.shortdesc;
        this.fopstate = operatorDetailsResponse.fopstate;
        this.foparea = operatorDetailsResponse.foparea;
        this.fopcode = operatorDetailsResponse.fopcode;
    }
}
export class OperatorServicePointsDataModel {
    splist: OperatorServicePointListModel[];
    fopid: number;
    associatedsplist: OperatorServicePointListModel[];
    availablesplist: OperatorServicePointListModel[];

}
