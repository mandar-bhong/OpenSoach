package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

public class APIResponseModel<T>  {

    @SerializedName("issuccess")
    public boolean IsSuccess;

    @SerializedName("error")
    public APIResponseErrorModel Error;

    @SerializedName("data")
    public T Data;
}
