package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 9/10/2017.
 */

public class PacketPartDrawingDataModel {

    @SerializedName("drawingid")
    public Integer DrawingID;
    @SerializedName("partid")
    public Integer PartID;
    @SerializedName("path")
    public String Path;
}
