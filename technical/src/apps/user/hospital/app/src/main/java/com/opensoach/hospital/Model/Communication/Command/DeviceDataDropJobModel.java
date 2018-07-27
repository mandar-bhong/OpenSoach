package com.opensoach.hospital.Model.Communication.Command;

import android.os.Parcelable;

import com.google.auto.value.AutoValue;
import com.opensoach.hospital.Model.Communication.DeviceDataBaseModel;

import java.util.Date;

/**
 * Created by Mandar on 12-06-2018.
 */

@AutoValue
public abstract class DeviceDataDropJobModel extends DeviceDataBaseModel implements Parcelable {
    public abstract Integer LocationID();
    public abstract Integer JobID();
    public abstract String OperatorCode();
    public abstract Date EndTime();

    public static DeviceDataDropJobModel create(Integer locationID,Integer jobID,String operatorCode,Date startTime)
    {
        return new AutoValue_DeviceDataDropJobModel(locationID,jobID,operatorCode, startTime);
    }
}
