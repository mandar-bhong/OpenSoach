package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

/**
 * Created by Mandar on 8/27/2017.
 */

public class PacketAuthCodeDataModel {

    public PacketAuthCodeDataModel(){
        AuthCodes = new ArrayList<>();
    }
    @SerializedName("opcodes")
    public ArrayList<String> AuthCodes;
}
