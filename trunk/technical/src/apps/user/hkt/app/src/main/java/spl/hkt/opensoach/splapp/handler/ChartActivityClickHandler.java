package spl.hkt.opensoach.splapp.handler;

import android.util.Log;
import android.view.View;
import android.widget.Toast;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;

import spl.hkt.opensoach.splapp.R;
import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.communication.CommunicationManager;
import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.manager.SendPacketManager;
import spl.hkt.opensoach.splapp.model.ChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.DeviceChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketUserComplaintDataModel;
import spl.hkt.opensoach.splapp.view.DialogHelper;
import spl.hkt.opensoach.splapp.viewModels.MainViewModel;

/**
 * Created by Mandar on 4/1/2017.
 */

public class ChartActivityClickHandler implements View.OnClickListener {
    @Override
    public void onClick(final View clickedView) {

        switch (clickedView.getId()) {
            case R.id.uploadData: {
                if (AppRepo.getInstance().isAuthCodeRequired()) {
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
                            });
                } else {
                    processChartData(null);
                }

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
                            public void onSucess(String complaintsBy, String userComments) {
                                processUserComplaint(complaintsBy, userComments);
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

        DeviceChartDataModel deviceChartDataModel = new DeviceChartDataModel(chartDataList);
        deviceChartDataModel.setCommandType(CommandConstants.DEVICE_DATA_COMMAND_CHART_DATA);
        SendPacketManager.Instance().send(deviceChartDataModel);

        MainViewModel.getInstance().createNewCurrenClickeCellModelMap();
    }

    private void processUserComplaint(String complaintsBy, String userComments) {
        //TODO Proces User Comments
        Log.i("ClickHandler", "complaintsBy : " + complaintsBy + " userComments : " + userComments);

        SimpleDateFormat raiseOnDateFormat = new SimpleDateFormat("yyyy:MM:dd");

        PacketUserComplaintDataModel packetUserComplaintDataModel = new PacketUserComplaintDataModel();
        packetUserComplaintDataModel.Description = userComments;
        packetUserComplaintDataModel.ComplaintBy = complaintsBy;
        packetUserComplaintDataModel.LocationId = AppRepo.getInstance().getCurrentLocationId();
        packetUserComplaintDataModel.RaisedOn = raiseOnDateFormat.format(new Date());
        packetUserComplaintDataModel.EmailId = "";
        packetUserComplaintDataModel.EmployeeID = "";
        packetUserComplaintDataModel.MobileNo = "";

        String JSONPacket = PacketHelper.GetComplaintPacket(packetUserComplaintDataModel);
        CommunicationManager.getInstance().SendPacket(JSONPacket);
    }
}
