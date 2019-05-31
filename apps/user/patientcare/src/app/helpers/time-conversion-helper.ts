import * as moment from "moment";
export class TimeConversion {
    static timeConvert(value: number) {
        const date = new Date(new Date().setHours(0, value, 0, 0));
        var rhours = date.getHours();
        var rminutes = date.getMinutes();
        if (rhours > 0) {
            if (rminutes > 0) {
                return rhours + " hr and " + rminutes + " mins";
            } else {
                return rhours + " hours";
            }
        } else {
            return rminutes + " minutes";
        }
    }
    static getStartTime(startTime: string): number {
        let t = startTime.toString();
        const time = t.split('.');
        const minutes = 60 * Number(time[0]);
        const totalminutes = minutes + Number(time[1]);
        return totalminutes;
    }
    
    //This method will take local date. UTC format convertion done internally
    static getServerShortTimeFormat(date: Date) {
        return moment.utc(date).format('YYYY-MM-DDTHH:mm:ss') + 'Z'
    }
}