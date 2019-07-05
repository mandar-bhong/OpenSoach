package com.opensoach.hpft.Model.Communication;

import com.google.gson.annotations.SerializedName;

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
