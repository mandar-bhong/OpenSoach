
import { USER_LAB } from '../../../shared/app-common-constants';


export class HPFTRouteHelper {
   static getUserHomeRoute(userrole: string) {
       
        if (userrole === USER_LAB) {
            return "/hospitals";
        }
        else {
            return "";
        }
    }
}