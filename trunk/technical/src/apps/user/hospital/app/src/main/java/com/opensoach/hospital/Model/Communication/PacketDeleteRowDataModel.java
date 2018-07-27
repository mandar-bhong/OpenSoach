package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.List;

/**
 * Created by Mandar on 9/23/2017.
 */

public class PacketDeleteRowDataModel {

    @SerializedName("tableid")
    public Integer TableID;
    @SerializedName("rowids")
    public List<Integer> RowIDs;

}
