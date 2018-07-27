package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 12-06-2018.
 */

public class APIAuthRequesetModel {
    @SerializedName("username")
    public String UserName;

    @SerializedName("password")
    public String Password;
}
