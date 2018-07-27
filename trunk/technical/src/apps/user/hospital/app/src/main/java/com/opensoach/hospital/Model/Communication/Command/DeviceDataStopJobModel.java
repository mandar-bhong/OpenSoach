package com.opensoach.hospital.Model.Communication.Command;

import android.os.Parcelable;

import com.google.auto.value.AutoValue;
import com.opensoach.hospital.Model.Communication.DeviceDataBaseModel;

import java.util.Date;

/**
 * Created by Mandar on 26-11-2017.
 */

@AutoValue
public abstract class DeviceDataStopJobModel extends DeviceDataBaseModel implements Parcelable {

    public abstract Integer LocationID();
    public abstract Integer JobID();
    public abstract String OperatorCode();
    public abstract Date EndTime();

    public static DeviceDataStopJobModel create(Integer locationID,Integer jobID,String operatorCode,Date endTime)
    {
        return new AutoValue_DeviceDataStopJobModel(locationID,jobID,operatorCode, endTime);
    }
}
