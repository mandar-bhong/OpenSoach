package spl.hkt.opensoach.splapp.model.communication;

import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

/**
 * Created by Sanjay on 22/05/2018.
 */

public class PacketLocationDataModel {
//@SerializedName("payload")
//    public ArrayList<LocationDataModel> Locations;

    @SerializedName("spid")
    public int SPID;
    @SerializedName("spname")
    public String LocationName;
    @SerializedName("spcname")
    public String CatgoryName;
}
