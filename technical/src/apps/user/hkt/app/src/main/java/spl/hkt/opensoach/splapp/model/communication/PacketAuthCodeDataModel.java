package spl.hkt.opensoach.splapp.model.communication;

import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;
import java.util.List;


public class PacketAuthCodeDataModel {

    public PacketAuthCodeDataModel(){
        AuthCodes = new ArrayList<>();
    }

    @SerializedName("locationid")
    public  int LocationId;
    @SerializedName("authcodes")
    public ArrayList<String> AuthCodes;
}
