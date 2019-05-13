package spl.hkt.opensoach.splapp.manager;

import android.os.AsyncTask;
import android.util.Log;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.communication.CommunicationManager;
import spl.hkt.opensoach.splapp.communication.IWebSocketConnection;
import spl.hkt.opensoach.splapp.helper.AppAction;

/**
 * Created by Mandar on 2/26/2017.
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
        CommunicationManager.getInstance().Connect(AppRepo.getInstance().getServerWebSocketURL());
    }

    @Override
    public void OnConnect() {
        AppRepo.getInstance().IsServerConnected(true);
        SendPacketManager.Instance().send(AppAction.ON_CONNECTION, AppRepo.getInstance().getAuthToken());
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

        AsyncTask.execute(updateConnectionStatus);
    }


    //TODO: Need to implement timer for every miniute re-connection on disconnect
}
