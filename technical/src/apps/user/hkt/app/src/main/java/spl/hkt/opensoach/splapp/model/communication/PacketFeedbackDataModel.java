package spl.hkt.opensoach.splapp.model.communication;

import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

/**
 * Created by Mandar on 8/14/2017.
 */

public class PacketFeedbackDataModel extends DeviceDataBaseModel {

    @SerializedName("rating")
    public int Rating;

}
