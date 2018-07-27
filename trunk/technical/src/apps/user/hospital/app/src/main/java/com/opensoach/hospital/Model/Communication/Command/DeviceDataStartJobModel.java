package com.opensoach.hospital.Model.Communication.Command;

import android.os.Parcelable;

import com.google.auto.value.AutoValue;
import com.opensoach.hospital.Model.Communication.DeviceDataBaseModel;

import java.util.Date;

/**
 * Created by Mandar on 26-11-2017.
 */

@AutoValue
public abstract class DeviceDataStartJobModel extends DeviceDataBaseModel implements Parcelable {

    public abstract Integer LocationID();
    public abstract Integer JobID();
    public abstract String OperatorCode();
    public abstract Date StartTime();

    public static DeviceDataStartJobModel create(Integer locationID,Integer jobID,String operatorCode,Date startTime)
    {
        return new AutoValue_DeviceDataStartJobModel(locationID,jobID,operatorCode, startTime);
    }
}
