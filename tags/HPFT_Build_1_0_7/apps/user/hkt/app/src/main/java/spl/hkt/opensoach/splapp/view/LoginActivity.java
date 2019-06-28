package spl.hkt.opensoach.splapp.view;

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.EditText;
import android.widget.TextView;

import spl.hkt.opensoach.splapp.R;
import spl.hkt.opensoach.splapp.logger.AppLogger;
import spl.hkt.opensoach.splapp.scheduler.ScheduleManager;

public class LoginActivity extends Activity implements View.OnClickListener {

    private EditText mIpAddress;
    TextView testString;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        setContentView(R.layout.activity_login);

        mIpAddress = (EditText) findViewById(R.id.etIpAddress);

        //init Scheuler
        ScheduleManager mScheduleManager = new ScheduleManager();
        mScheduleManager.init(this);
        mScheduleManager.startScheduler(this, 1, 30, 1);


    }

    public void doSubmit(View v) {

        if (mIpAddress.getText() != null && mIpAddress.getText().toString() != null) {

            String strIPAddress = mIpAddress.getText().toString();

            Intent intent = new Intent(this, DashboardActivity.class);
            startActivity(intent);
        }
    }

    public void showData(View view) {
        testString = (TextView) findViewById(R.id.testString);
        String str = "";

       /* ChartTableManager chartTableManager = new ChartTableManager();
        List<DBChartTableRowModel> chartTableRowModelList = chartTableManager.getAllChartData();
        if (chartTableRowModelList != null && chartTableRowModelList.size() > 0) {
            DBChartTableRowModel chartTableRowModel;

            for (int i = 0; i < chartTableRowModelList.size(); i++) {
                chartTableRowModel = chartTableRowModelList.get(i);
                String result = " < " + chartTableRowModel.getChartName() + " : " + chartTableRowModel.getCustomerId() + " : " + chartTableRowModel.getLocationCategoryId() + " : " + chartTableRowModel.getSlotInterval() + " >";
                str = str + "  \n  " + result;
            }
        } else {
            str = "No Data";
        }*/

        //List<DBChartDataTableRowModel> tasksRow = DatabaseManager.SelectAll(new DBChartDataTableQueryModel(), new DBChartDataTableRowModel());
/*
        int i = 0;
        for (DBChartDataTableRowModel model : tasksRow) {
            Log.i("Cell : " + i++, "chart ID " + model.getChartId() +
                    "task ID " + model.getTaskId() +
                    "slot ID " + model.getSlotId()
            );
        }

        testString.setText(str);*/
    }

    @Override
    protected void onDestroy() {
        super.onDestroy();

        //Stop The Alarm Manager & Cancel the PendingIntent
        new ScheduleManager().stopSchedular(this);
    }

    @Override
    public void onClick(View view) {

    }
}
