package spl.hkt.opensoach.splapp.model.communication;

import com.google.gson.annotations.SerializedName;

public class APIResponseModel<T>  {

    @SerializedName("issuccess")
    public boolean IsSuccess;

    @SerializedName("error")
    public APIResponseErrorModel Error;

    @SerializedName("data")
    public T Data;
}
