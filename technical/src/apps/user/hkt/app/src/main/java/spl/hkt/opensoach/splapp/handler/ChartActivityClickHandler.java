package spl.hkt.opensoach.splapp.handler;

import android.util.Log;
import android.view.View;
import android.widget.Toast;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.TimeZone;

import spl.hkt.opensoach.splapp.R;
import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.helper.AppAction;
import spl.hkt.opensoach.splapp.manager.SendPacketManager;
import spl.hkt.opensoach.splapp.model.ChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.DeviceChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketUserComplaintDataModel;
import spl.hkt.opensoach.splapp.view.DialogHelper;
import spl.hkt.opensoach.splapp.viewModels.MainViewModel;

import static spl.hkt.opensoach.splapp.helper.ApplicationConstants.PACKET_DATE_FORMAT;

/**
 * Created by Mandar on 4/1/2017.
 */

public class ChartActivityClickHandler implements View.OnClickListener {
    @Override
    public void onClick(final View clickedView) {

        switch (clickedView.getId()) {
            case R.id.uploadData: {
                DialogHelper.showSingleLineEditTextAlert(
                        clickedView.getContext(),
                        clickedView.getContext().getResources().getString(R.string.dialog_enter_auth_code),
                        new DialogHelper.DialogCallBack() {

                            @Override
                            public void onSucess(String authText) {

                                if (AppRepo.getInstance().getAuthCodeList().contains(authText)) {
                                    processChartData(authText);
                                } else {
                                    Toast.makeText(
                                            clickedView.getContext(),
                                            clickedView.getContext().getResources().getString(R.string.invalid_auth_code),
                                            Toast.LENGTH_LONG).show();
                                }
                            }

                            @Override
                            public void onSucess(String strData1, String strData2) {

                            }

                            @Override
                            public void onSucess(String strData1, String strData2, String strData3) {

                            }
                        });
            }
            break;
            case R.id.imgCommentView: {
                DialogHelper.showComplaintDialog(clickedView.getContext(),
                        clickedView.getContext().getResources().getString(R.string.dialog_complaint_title),
                        new DialogHelper.DialogCallBack() {

                            @Override
                            public void onSucess(String strData) {

                            }

                            @Override
                            public void onSucess(String strData1, String strData2) {

                            }

                            @Override
                            public void onSucess(String complaintBy, String title, String details) {
                                processUserComplaint(complaintBy, title, details);
                            }
                        });
            }
            break;

        }
    }

    private void processChartData(String authCode) {
        ArrayList<ChartDataModel> chartDataList = new ArrayList<ChartDataModel>(MainViewModel.getInstance().getCurrenChartDataModelMap().values());

        ChartDataModel chartDataModel;
        if (authCode != null) {
            for (int i = 0; i < chartDataList.size(); i++) {
                chartDataModel = chartDataList.get(i);
                chartDataModel.setAuthCode(authCode);
                chartDataList.set(i, chartDataModel);
            }
        }

        final DeviceChartDataModel deviceChartDataModel = new DeviceChartDataModel(chartDataList);
        deviceChartDataModel.setUserActionType(AppAction.CHART_DATA);
        SendPacketManager.Instance().send(AppAction.CHART_DATA, deviceChartDataModel);

        MainViewModel.getInstance().createNewCurrenClickeCellModelMap();
    }

    private void processUserComplaint(String complaintBy, String title, String details) {
        //TODO Proces User Comments
        Log.i("ClickHandler", "complaintBy : " + complaintBy + " complaintDetails : " + details);


        SimpleDateFormat raiseOnDateFormat = new SimpleDateFormat(PACKET_DATE_FORMAT);
        raiseOnDateFormat.setTimeZone(TimeZone.getTimeZone("GMT"));

        ArrayList<PacketUserComplaintDataModel> complaints = new ArrayList<>();
        PacketUserComplaintDataModel packetUserComplaintDataModel = new PacketUserComplaintDataModel();
        packetUserComplaintDataModel.ComplaintBy = complaintBy;
        packetUserComplaintDataModel.ComplaintTitle = title;
        packetUserComplaintDataModel.Description = details;
        packetUserComplaintDataModel.RaisedOn = raiseOnDateFormat.format(new Date());

        complaints.add(packetUserComplaintDataModel);
        SendPacketManager.Instance().send(AppAction.COMPLAINT_DATA, complaints);
    }
}
