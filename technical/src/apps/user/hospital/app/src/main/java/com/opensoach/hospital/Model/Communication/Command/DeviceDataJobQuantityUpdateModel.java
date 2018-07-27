package com.opensoach.hospital.Model.Communication.Command;

import android.os.Parcelable;

import com.google.auto.value.AutoValue;
import com.opensoach.hospital.Model.Communication.DeviceDataBaseModel;

import java.util.Date;

/**
 * Created by Mandar on 25-11-2017.
 */

@AutoValue
public abstract class DeviceDataJobQuantityUpdateModel extends DeviceDataBaseModel implements Parcelable {

    public abstract Integer LocationID();
    public abstract Integer JobID();
    public abstract String OperatorCode();
    public abstract Integer FinishedPartCount();
    public abstract Date CompletionTime();
    public abstract String Comment();

    public static DeviceDataJobQuantityUpdateModel create(Integer locationID,Integer jobID,String operatorCode,Integer finishedPartCount, Date completionTime,String comment)
    {
        return new AutoValue_DeviceDataJobQuantityUpdateModel(locationID,jobID,operatorCode,finishedPartCount, completionTime,comment);
    }

}
