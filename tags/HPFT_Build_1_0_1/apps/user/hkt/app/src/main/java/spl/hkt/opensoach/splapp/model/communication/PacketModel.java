package spl.hkt.opensoach.splapp.model.communication;
import com.google.gson.annotations.SerializedName;

public  class  PacketModel<T>{

    @SerializedName("header")
    public  PacketHeaderModel Header;
    @SerializedName("payload")
    public T Payload;

}



