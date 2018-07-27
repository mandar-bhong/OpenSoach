package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 8/26/2017.
 */

public class PacketHeaderModel {
    @SerializedName("crc")
    public String CRC;
    @SerializedName("category")
    public  int  Category;
    @SerializedName("commandid")
    public  int  CommandID  ;
    @SerializedName("seqid")
    public int  SeqID ;
    @SerializedName("locationid")
    public  int  LocationID  ;
}
