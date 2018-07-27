package com.opensoach.hospital.Manager;

import android.os.AsyncTask;
import android.os.Build;
import android.util.Log;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.Communication.CommunicationManager;
import com.opensoach.hospital.Communication.IWebSocketConnection;
import com.opensoach.hospital.Helper.Constants;
import com.opensoach.hospital.Helper.PacketHelper;
import com.opensoach.hospital.Utility.AppLogger;

/**
 * Created by Mandar on 8/26/2017.
 */

public class ServerConnectionManager implements IWebSocketConnection {

    private static ServerConnectionManager singleton;


    private ServerConnectionManager() {

    }

    /* Static 'instance' method */
    public static ServerConnectionManager Instance() {
        if (singleton == null)
            singleton = new ServerConnectionManager();
        return singleton;
    }


    public boolean Init() {
        return true;
    }


    public void Connect() {
        CommunicationManager.getInstance().Connect(Constants.WEB_SOCKET_URL);
    }

    @Override
    public void OnConnect() {

        Runnable sendStartUpPacketRunnable = new Runnable() {
            @Override
            public void run() {
                AppRepo.getInstance().IsServerConnected(true);

                String packetData = PacketHelper.GetStartUpPacket(Build.SERIAL);
                try {
                    CommunicationManager.getInstance().SendPacket(packetData);

                    AppLogger.getInstance().Log(AppLogger.LogLevel.Error,packetData);
                } catch (Exception e) {
                    AppLogger.getInstance().Log(e);
                }
            }
        };

        AsyncTask.execute(sendStartUpPacketRunnable);
    }

    @Override
    public void OnMessage(String strRespone) {
        try {
            PacketManager.getInstance().handleReceivedPacket(strRespone);
        } catch (Exception e) {
            Log.i("Exception", " " + e.getMessage());
        }
    }

    @Override
    public void OnDisconnect(int code, String reason) {

        Runnable updateConnectionStatus = new Runnable() {
            @Override
            public void run() {
                AppRepo.getInstance().IsServerConnected(false);
            }
        };

        AsyncTask.execute(updateConnectionStatus);
    }

    @Override
    public void OnError(Exception error) {

        Runnable updateConnectionStatus = new Runnable() {
            @Override
            public void run() {
                AppRepo.getInstance().IsServerConnected(false);
            }
        };

        AppLogger.getInstance().Log(AppLogger.LogLevel.Debug,"Raising On connection error");
        AsyncTask.execute(updateConnectionStatus);
    }

}
