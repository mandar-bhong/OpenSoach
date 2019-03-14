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
}