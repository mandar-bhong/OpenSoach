export class TimeConversionHelper
{
    static timeConvert(time:number) {
        var num = time;
        var hours = (num / 60);
        var rhours = Math.floor(hours);
        var minutes = (hours - rhours) * 60;
        var rminutes = Math.round(minutes);
        if (rhours > 0) {
          if (rminutes > 0) {
            return rhours + " hour & " + rminutes + " minute";
          } else {
            return rhours + " hour";
          }
        } else {
          return rminutes + " minute";
        }
      }
}