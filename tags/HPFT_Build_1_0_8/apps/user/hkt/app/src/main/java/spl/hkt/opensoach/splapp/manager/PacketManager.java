package spl.hkt.opensoach.splapp.manager;

import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.os.Message;
import android.os.Parcelable;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.communication.WebSocketHandler;
import spl.hkt.opensoach.splapp.model.communication.DeviceChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.processor.PacketProcessor;

/**
 * Created by Mandar on 2/25/2017.
 */

public class PacketManager extends Thread {

    private static PacketManager singleton ;

    private Handler packetHandler;


    /* A private Constructor prevents any other
     * class from instantiating.
     */
    private PacketManager() {

    }

    public boolean Init(){
        start();//TODO: Handle error
        return  true;
    }

    public void DeInit(){
        //stop();//TODO: Deint this class
    }

    /* Static 'instance' method */
    public static PacketManager getInstance( ) {
        if(singleton == null)
            singleton = new PacketManager( );
        return singleton;
    }

    public void handleReceivedPacket(String packet) {
        Message message = new Message();
        Bundle b = new Bundle();
        b.putString("Server_Received_Packet",packet);
        message.setData(b);

        packetHandler.sendMessage(message);
    }

    @Override
    public void run() {
        Looper.prepare();
        packetHandler = new PacketProcessor();
        Looper.loop();
    }
}
