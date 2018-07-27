package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.List;

/**
 * Created by Mandar on 9/5/2017.
 */

public class PacketJobCardsDataModel {
    @SerializedName("jobcards")
    public List<PacketJobCardDataModel> JobCards;
}
