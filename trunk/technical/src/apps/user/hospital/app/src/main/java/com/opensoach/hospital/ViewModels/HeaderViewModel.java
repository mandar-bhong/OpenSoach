package com.opensoach.hospital.ViewModels;

import android.databinding.BaseObservable;
import android.databinding.Bindable;
import android.databinding.InverseBindingMethod;
import android.databinding.InverseBindingMethods;
import android.widget.Spinner;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.BR;
import com.opensoach.hospital.Model.DB.DBLocationTableRowModel;
import com.opensoach.hospital.Utility.AppLogger;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Samir
 */


@InverseBindingMethods({
        @InverseBindingMethod(type = Spinner.class, attribute = "android:selectedItemPosition"),
})
public class HeaderViewModel extends BaseObservable {

    private ArrayList<String> locationList;
    private int position;
    private String locationName;
    private List<DBLocationTableRowModel> locations;
    private List<Integer> locationIDLList;

    public HeaderViewModel() {
        //Temp Adding Location

        this.locationList = new ArrayList<>();
        locations = new ArrayList<>();
    }

    public List<DBLocationTableRowModel> getLocations() {

        if(locations.size()== 1) {
            locationName = locations.get(0).getLocationName();
            notifyPropertyChanged(BR.locationName);
        }

        return AppRepo.getInstance().getLocationList();
    }

    public void setLocations(List<DBLocationTableRowModel> locations) {
        this.locations = locations;

        notifyPropertyChanged(BR.locationList);
    }

    @Bindable
    public ArrayList<String> getLocationList() {

        ArrayList<String> locationList = new ArrayList<>();

        locationIDLList = new ArrayList<>();

        try {
            for (DBLocationTableRowModel location : AppRepo.getInstance().getLocationList()) {
                locationList.add(location.getLocationName());
                locationIDLList.add(location.getLocationId());
            }

            if (AppRepo.getInstance().getLocationList().size() == 1) {
                locationName = locations.get(0).getLocationName();
                notifyPropertyChanged(BR.locationName);
            }
        }catch (Exception ex){
            AppLogger.getInstance().Log(AppLogger.LogLevel.Error,ex);
        }

        return locationList;
    }

    @Bindable
    public int getPosition() {
        return position;
    }

    public void setPosition(int position) {
        this.position = position;
        Integer selectedLocationID = locationIDLList.get(position);

        AppRepo.getInstance().setCurrentLocationId(selectedLocationID);
    }

    @Bindable
    public String getLocationName() {
       return locationName;
    }

    public void setLocationName(String locationName) {
        this.locationName = locationName;
        notifyPropertyChanged(BR.locationName);
    }

    public int getPosition(Spinner spinner) {
        return spinner.getSelectedItemPosition();
    }

    @Override
    public void notifyChange() {
        notifyPropertyChanged(BR._all);

        super.notifyChange();
    }
}
