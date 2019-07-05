package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

public class PacketServiceTaskItemDataModel {

    @SerializedName("taskname")
    public String taskName;

    @SerializedName("comment")
    public String Comment;

    @SerializedName("note")
    public String Note;//Operator comment

    @SerializedName("cost")
    public String Cost;
}
