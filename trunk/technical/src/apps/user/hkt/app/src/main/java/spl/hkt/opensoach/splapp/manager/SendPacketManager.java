package spl.hkt.opensoach.splapp.manager;

import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.os.Message;
import android.os.Parcelable;
import android.util.Log;

import com.google.gson.Gson;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.Date;
import java.util.List;
import java.util.TimeZone;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.communication.CommunicationManager;
import spl.hkt.opensoach.splapp.dal.DatabaseManager;
import spl.hkt.opensoach.splapp.helper.AppHelper;
import spl.hkt.opensoach.splapp.helper.ApplicationConstants;
import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.model.ChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.DeviceChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.DeviceDataBaseModel;
import spl.hkt.opensoach.splapp.model.communication.PacketChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketServiceInstanceTxnModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableRowModel;

/**
 * Created by samir.s.bukkawar on 3/25/2017.
 */

public class SendPacketManager extends Thread {

    private static SendPacketManager singleton;

    private Handler sendPacketHandler;


    /* A private Constructor prevents any other
     * class from instantiating.
     */
    private SendPacketManager() {

    }

    public boolean Init() {
        start();//TODO: Handle error
        return true;
    }

    public void DeInit() {
        //stop();//TODO: Deint this class
    }

    /* Static 'instance' method */
    public static SendPacketManager Instance() {
        if (singleton == null)
            singleton = new SendPacketManager();
        return singleton;
    }

    @Override
    public void run() {
        Looper.prepare();
        sendPacketHandler = new SendPacketHandler();
        Looper.loop();
    }

    public void send(DeviceDataBaseModel deviceDataBaseModel) {
        Message message = new Message();
        Bundle b = new Bundle();
        b.putParcelable("CHART_DATA", (Parcelable) deviceDataBaseModel);
        message.setData(b);

        sendPacketHandler.sendMessage(message);
    }
}

class SendPacketHandler extends Handler {
    @Override
    public void handleMessage(Message msg) {
        super.handleMessage(msg);

        Bundle b = msg.getData();
        DeviceDataBaseModel deviceDataBaseModel = (DeviceDataBaseModel) b.get("CHART_DATA");

        try {

            switch (deviceDataBaseModel.getCommandType()) {
                case CommandConstants.DEVICE_DATA_COMMAND_CHART_DATA:
                    DeviceChartDataModel devideChartDataModel = (DeviceChartDataModel) deviceDataBaseModel;
                    ProcessSendMessage(devideChartDataModel);
                    break;
                case CommandConstants.DEVICE_DATA_COMMAND_CHART_UNSYNC_DATA:
                    break;
            }


        } catch (Exception ex) {
            Log.d("ParsingSendChartData", ex.getMessage());
        }
    }

    private void ProcessSendMessage(DeviceChartDataModel devideChartDataModel) {

        ArrayList<ChartDataModel> chartDataList = devideChartDataModel.getChartDataModels();

        //TODO: Save packet to database
        try {

            List<DBChartDataTableRowModel> dbChartDataItems = updateTableChartData(devideChartDataModel);

            AppHelper.NotifyChartDataStatusUpdate(dbChartDataItems);

        } catch (ParseException e) {
            e.printStackTrace();
        }

        if (AppRepo.getInstance().IsServerConnected()) {

            SimpleDateFormat UTCEntryTimeFormat = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSS'Z'");
            UTCEntryTimeFormat.setTimeZone(TimeZone.getTimeZone("GMT"));

            ArrayList<PacketServiceInstanceTxnModel> txns = new ArrayList<>();

            for (ChartDataModel model : devideChartDataModel.getChartDataModels()) {
                PacketServiceInstanceTxnModel txnModel = new PacketServiceInstanceTxnModel();
                PacketChartDataModel packetChartDataModel = new PacketChartDataModel();
                txnModel.servinid = model.getChartId();
                txnModel.txndate = UTCEntryTimeFormat.format(model.getEntryDate());
                txnModel.fopcode = model.getAuthCode();
                if (model.getEntryDate().getTime() < model.getSlotEndTime().getTime()) {
                    txnModel.status = 1;//On time
                } else {
                    txnModel.status = 2; // Delay
                }

                packetChartDataModel.taskName = model.getTaskName();

                Calendar calChartStart = Calendar.getInstance();
                calChartStart.setTime(model.getSlotStartTime());
                packetChartDataModel.slotStartTime = calChartStart.get(Calendar.HOUR_OF_DAY) * 60 + calChartStart.get(Calendar.MINUTE);

                Calendar calChartEnd = Calendar.getInstance();
                calChartEnd.setTime(model.getSlotEndTime());
                packetChartDataModel.slotEndTime = calChartEnd.get(Calendar.HOUR_OF_DAY) * 60 + calChartEnd.get(Calendar.MINUTE);

                txnModel.txndata = new Gson().toJson(packetChartDataModel);

                txns.add(txnModel);
            }

            String packet = PacketHelper.GetChartDataPacket(txns);

            CommunicationManager.getInstance().SendPacket(packet);
        }
    }

    private List<DBChartDataTableRowModel> updateTableChartData(DeviceChartDataModel devideChartDataModel) throws ParseException {

        List<DBChartDataTableRowModel> dbChartDataItems = new ArrayList<>();

        for (ChartDataModel model : devideChartDataModel.getChartDataModels()) {
            DBChartDataTableRowModel dbChartDataTableRowModel = new DBChartDataTableRowModel();

            dbChartDataTableRowModel.setChartId(model.getChartId());
            dbChartDataTableRowModel.setTaskName(model.getTaskName());
            dbChartDataTableRowModel.setSlotId(model.getSlotId());
            dbChartDataTableRowModel.setEntryTime(model.getEntryDate());
            dbChartDataTableRowModel.setSlotStartTime(model.getSlotStartTime());
            dbChartDataTableRowModel.setSlotEndTime(model.getSlotEndTime());
            dbChartDataTableRowModel.setSynced(false);
            dbChartDataTableRowModel.setAuthCode(model.getAuthCode());

            SimpleDateFormat chartStartDate = new SimpleDateFormat("yyyy-MM-dd");
            Date dateWithoutTime = chartStartDate.parse(chartStartDate.format(model.getEntryDate()));

            dbChartDataTableRowModel.setChartDay(dateWithoutTime);

            if (model.getSlotEndTime().getTime() < model.getEntryDate().getTime()) {
                dbChartDataTableRowModel.setCellState(ApplicationConstants.DB_CHART_STATE_DELAYED);
            } else {
                dbChartDataTableRowModel.setCellState(ApplicationConstants.DB_CHART_STATE_ON_TIME);
            }

            DatabaseManager.InsertRow(dbChartDataTableRowModel);

            dbChartDataItems.add(dbChartDataTableRowModel);
        }

        return dbChartDataItems;
    }
}