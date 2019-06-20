import { USER_CATEGORY } from '../../../../shared/app-common-constants';
import { USER_STATE } from '../../../../shared/app-common-constants';
import { UserFilterRequest } from '../api/user-models';
export class UserFilterModel {
    cpmid: number;
    usrname: string;
    usrcategory: USER_CATEGORY;
    usrstate: USER_STATE;

    copyTo(userFilterRequest: UserFilterRequest) {
        userFilterRequest.cpmid = this.cpmid;
        userFilterRequest.usrname = this.usrname;
        userFilterRequest.usrcategory = this.usrcategory;
        userFilterRequest.usrstate = this.usrstate;
    }
}
