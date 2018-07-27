package com.opensoach.hospital.ViewModels;

import android.widget.BaseAdapter;

import com.opensoach.hospital.Views.Interfaces.IList;
import com.opensoach.hospital.Views.Interfaces.IGridView;
import com.opensoach.hospital.Views.Miscellaneous.JobBriefViewAdaptor;

import java.util.Iterator;
import java.util.List;

/**
 * Created by Mandar on 8/25/2017.
 */

public class JobGridViewModel extends BaseViewModel implements IList<JobBriefViewModel>,IGridView {

    private JobBriefViewAdaptor dataAdapter;


    public JobGridViewModel(){
        dataAdapter = new JobBriefViewAdaptor();
    }

    @Override
    public BaseAdapter getDataAdaptor() {
        return dataAdapter;
    }

    @Override
    public BaseAdapter setItemsSource(Iterator item) {
        return null;
    }

    @Override
    public List<JobBriefViewModel> getItemsSource() {
      return   dataAdapter.getItemsSource();
    }

    @Override
    public void setItemsSource(List<JobBriefViewModel> source) {
        dataAdapter.setItemsSource(source);
    }
}
