package com.opensoach.hpft.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.List;

/**
 * Created by Mandar on 06-08-2018.
 */

public class PacketTaskConfigModel {

    @SerializedName("tasks")
    public List<PacketTaskModel> Tasks;

}
