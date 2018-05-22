package spl.hkt.opensoach.splapp.model.communication;
import com.google.gson.annotations.SerializedName;
/**
 * Created by Mandar on 2/26/2017.
 */

public class PacketTaskModel {
    @SerializedName("taskid")
    public int TaskID;
    @SerializedName("taskname")
    public  String  TaskName;
    @SerializedName("taskorder")
    public  int  TaskOrder  ;
}
