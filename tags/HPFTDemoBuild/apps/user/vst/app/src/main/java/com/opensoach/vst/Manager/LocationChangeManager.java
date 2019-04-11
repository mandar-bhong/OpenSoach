package com.opensoach.vst.Manager;

import android.util.Log;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;

import com.opensoach.vst.AppRepo.AppRepo;

/**
 * Created by Mandar on 4/8/2017.
 */

public class LocationChangeManager implements PropertyChangeListener {

    private static LocationChangeManager singleton;

    private LocationChangeManager() {

    }

    /* Static 'instance' method */
    public static LocationChangeManager Instance() {
        if (singleton == null)
            singleton = new LocationChangeManager();
        return singleton;
    }


    @Override
    public void propertyChange(PropertyChangeEvent evt) {
        switch (evt.getPropertyName()) {
            case AppRepo.CurrentLocationIdPropName:
                Log.d("CurrentLocationId: ", evt.getNewValue().toString());
                Integer locationId = (Integer) evt.getNewValue();
                new LocationChartRunnable(locationId).run();
                break;
        }
    }
}



