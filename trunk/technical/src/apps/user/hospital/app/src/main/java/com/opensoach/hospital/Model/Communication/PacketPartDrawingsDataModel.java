package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.List;

/**
 * Created by Mandar on 9/10/2017.
 */

public class PacketPartDrawingsDataModel {
    @SerializedName("drawings")
    public List<PacketPartDrawingDataModel> PartDrawings;
}
