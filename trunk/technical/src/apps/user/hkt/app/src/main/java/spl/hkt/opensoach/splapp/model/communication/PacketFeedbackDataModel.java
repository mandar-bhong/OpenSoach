package spl.hkt.opensoach.splapp.model.communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 8/14/2017.
 */

public class PacketFeedbackDataModel extends DeviceDataBaseModel {

    @SerializedName("feedback")
    public int Feedback;

    @SerializedName("comment")
    public String Comment;

    @SerializedName("raisedon")
    public String RaisedOn;
}