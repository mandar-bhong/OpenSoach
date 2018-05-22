package spl.hkt.opensoach.splapp.model.communication;

import com.google.gson.annotations.SerializedName;


public class LocationDataModel {
    @SerializedName("spid")
    public int SPID;
    @SerializedName("spname")
    public String LocationName;
    @SerializedName("spcname")
    public String CatgoryName;
}
