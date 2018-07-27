package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 8/27/2017.
 */

public class PacketLocationDataModel {
    @SerializedName("locationid")
    public Integer LocationId;
    @SerializedName("name")
    public String Name;
}
