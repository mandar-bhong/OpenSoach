package spl.hkt.opensoach.splapp.manager;

import android.os.AsyncTask;
import android.util.Log;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.util.List;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.dal.DatabaseManager;
import spl.hkt.opensoach.splapp.helper.AppHelper;
import spl.hkt.opensoach.splapp.model.db.DBLocationTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBLocationTableRowModel;

/**
 * Created by Mandar on 4/8/2017.
 */

public class LocationChangeManager implements PropertyChangeListener {

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



