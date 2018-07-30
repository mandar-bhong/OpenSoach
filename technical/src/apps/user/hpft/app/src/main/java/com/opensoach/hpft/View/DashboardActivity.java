package com.opensoach.hpft.View;

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import android.os.Handler;
import android.view.View;
import android.widget.TextView;

import java.util.Timer;
import java.util.TimerTask;

import com.opensoach.hpft.R;
import com.opensoach.hpft.SPLApplication;

public class DashboardActivity extends Activity implements UpdateChartListner {

    private TextView tvResponseData;
    private String strResponse = "";
    private SPLApplication mSPLApplication;

    private Timer mRefreshViewTimer;
    private Handler mHandler;
    private static int updateinterval = 1000 * 10;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_dashboard);

        tvResponseData = (TextView) findViewById(R.id.tvResponseData);

        mSPLApplication = (SPLApplication) getApplication();

        mSPLApplication.registerListner(this, DashboardActivity.this);

        mHandler = new Handler();
    }

    @Override
    protected void onResume() {
        super.onResume();
        //mDeviceListAdapter.notifyDataSetChanged();
        mRefreshViewTimer = new Timer();
        //Schedule a runnalbe at interval for fetching data.
        mRefreshViewTimer.scheduleAtFixedRate(new TimerTask() {
            @Override
            public void run() {
                mHandler.post(new Runnable() {
                    @Override
                    public void run() {
                        // fetchApplinaceList();
                        //getCpassDeviceData();
                    }
                });
            }
        }, 1000, updateinterval);
    }


    @Override
    public void callback(String result) {

        strResponse = strResponse + "\n \n\n" + result;
        tvResponseData.setText(strResponse);

    }

    @Override
    protected void onDestroy() {
        super.onDestroy();
        mRefreshViewTimer.cancel();
    }


    public void doNext(View view) {

        Intent intent = new Intent(this, ChartActivity.class);
        startActivity(intent);
    }
}
