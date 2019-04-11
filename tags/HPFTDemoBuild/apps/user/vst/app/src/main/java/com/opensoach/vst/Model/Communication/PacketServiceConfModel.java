package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 06-08-2018.
 */

public class PacketServiceConfModel {

    @SerializedName("taskconf")
    public PacketTaskConfigModel TaskConfig;

    @SerializedName("timeconf")
    public PacketTimeConfigModel TimeConfig;
}
