package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 9/5/2017.
 */

public class PacketEnggPartDataModel {

    @SerializedName("partid")
    public Integer PartID;
    @SerializedName("partno")
    public String PartNo;
    @SerializedName("partrevision")
    public String PartRevision;
    @SerializedName("internalpartno")
    public String InternalPartNo;
    @SerializedName("process")
    public String Process;



}
