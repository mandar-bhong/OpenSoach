package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.List;

/**
 * Created by Mandar on 9/24/2017.
 */

public class PacketJobCardsStatusChangedDataModel {

    @SerializedName("jobcards")
    public List<PacketJobCardStatusChangedDataModel> JobCards;
}
