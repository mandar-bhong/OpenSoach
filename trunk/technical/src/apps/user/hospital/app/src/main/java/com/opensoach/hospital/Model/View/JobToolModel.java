package com.opensoach.hospital.Model.View;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 9/16/2017.
 */

public class JobToolModel {
    @SerializedName("toolid")
    public String ID;

    @SerializedName("toolname")
    public String Name;

    @SerializedName("specs")
    public String Description;
}
