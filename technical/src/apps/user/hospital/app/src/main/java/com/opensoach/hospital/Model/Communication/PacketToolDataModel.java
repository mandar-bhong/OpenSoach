package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 9/5/2017.
 */

public class PacketToolDataModel {
    @SerializedName("toolid")
    public Integer ToolID;
    @SerializedName("toolname")
    public String Name;
    @SerializedName("specs")
    public String Specs;
}
