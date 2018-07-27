package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.List;

/**
 * Created by Mandar on 9/2/2017.
 */

public class PacketLocationsDataModel {
    @SerializedName("locations")
    public List<PacketLocationDataModel> Locations;
}
