package spl.hkt.opensoach.splapp.model.communication;

import android.os.Parcel;
import android.os.Parcelable;

import java.util.ArrayList;

import spl.hkt.opensoach.splapp.model.ChartDataModel;

/**
 * Created by samir.s.bukkawar on 3/25/2017.
 */

public class DeviceChartDataModel extends DeviceDataBaseModel implements Parcelable {


    ArrayList<ChartDataModel> chartDataModels;

    public DeviceChartDataModel() {
        chartDataModels = new ArrayList<ChartDataModel>();
    }

    public DeviceChartDataModel(ArrayList<ChartDataModel> data) {
        chartDataModels = data;
    }

    public ArrayList<ChartDataModel> getChartDataModels() {
        return chartDataModels;
    }

    protected DeviceChartDataModel(Parcel in) {
        if (in.readByte() == 0x01) {
            chartDataModels = new ArrayList<ChartDataModel>();
            in.readList(chartDataModels, ChartDataModel.class.getClassLoader());
        } else {
            chartDataModels = null;
        }
    }

    @Override
    public int describeContents() {
        return 0;
    }

    @Override
    public void writeToParcel(Parcel dest, int flags) {
        if (chartDataModels == null) {
            dest.writeByte((byte) (0x00));
        } else {
            dest.writeByte((byte) (0x01));
            dest.writeList(chartDataModels);
        }
    }

    @SuppressWarnings("unused")
    public static final Parcelable.Creator<DeviceChartDataModel> CREATOR = new Parcelable.Creator<DeviceChartDataModel>() {
        @Override
        public DeviceChartDataModel createFromParcel(Parcel in) {
            return new DeviceChartDataModel(in);
        }

        @Override
        public DeviceChartDataModel[] newArray(int size) {
            return new DeviceChartDataModel[size];
        }
    };
}
