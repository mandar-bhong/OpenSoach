package spl.hkt.opensoach.splapp.model.communication;
import com.google.gson.annotations.SerializedName;
import java.util.List;


public class PacketChartConfigurationModel {

    @SerializedName("chartid")
    public int ChartID;
    @SerializedName("locationid")
    public  int  LocationID;
    @SerializedName("customerid")
    public  int  CustomerID  ;
    @SerializedName("locationcategoryid")
    public int  LocationCategoryID ;
    @SerializedName("chartname")
    public  String  ChartName  ;
    @SerializedName("starttime")
    public  int  StartTime  ;
    @SerializedName("endtime")
    public  int  EndTime  ;
    @SerializedName("slotinterval")
    public  int  SlotInterval  ;
    @SerializedName("tasks")
    public  List<PacketTaskModel>  Tasks  ;

}
