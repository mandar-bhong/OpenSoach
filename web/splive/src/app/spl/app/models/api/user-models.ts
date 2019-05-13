import { USER_CATEGORY } from '../../../../shared/app-common-constants';
import { USER_STATE } from '../../../../shared/app-common-constants';

export class UserFilterRequest {
    cpmid: number;
    usrname: string;
    usrcategory: USER_CATEGORY;
    usrstate: USER_STATE;
}
