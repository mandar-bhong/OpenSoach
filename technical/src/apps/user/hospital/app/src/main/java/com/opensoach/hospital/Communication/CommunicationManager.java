package com.opensoach.hospital.Communication;

import android.util.Log;

import java.io.IOException;
import java.util.List;
import java.util.Map;

import com.neovisionaries.ws.client.*;
import com.opensoach.hospital.Utility.AppLogger;

/**
 * Created by Mandar on 2/25/2017.
 */

public class CommunicationManager {

    private static CommunicationManager singleton;
    private IWebSocketConnection _webSocketEventHandler;
    private WebSocketFactory _factory;
    private WebSocket _webSocket;

    private CommunicationManager() {
        _factory = new WebSocketFactory().setConnectionTimeout(5000);
    }

    /* Static 'instance' method */
    public static CommunicationManager getInstance() {
        if (singleton == null)
            singleton = new CommunicationManager();
        return singleton;
    }

    public void DeInit() {
        //stop();//TODO: Deint this class
    }

    public boolean Init(IWebSocketConnection handler) {
        _webSocketEventHandler = handler;
        return true;
    }

    public boolean Connect(String WebSocketURL) {

        try {

            _webSocket=null;
            // Create a WebSocket. The timeout value set above is used.
            _webSocket = _factory.createSocket(WebSocketURL);

            _webSocket.addListener(new WebSocketAdapter() {
                @Override
                public void onTextMessage(WebSocket websocket, String message) throws Exception {
                    Log.d("Packet Received: ",message);
                    _webSocketEventHandler.OnMessage(message);                }

                @Override
                public void onConnectError(WebSocket websocket, WebSocketException exception) throws Exception {
                    _webSocketEventHandler.OnError(exception);
                }

                @Override
                public void onDisconnected(WebSocket websocket, WebSocketFrame serverCloseFrame, WebSocketFrame clientCloseFrame, boolean closedByServer) throws Exception {
                    Log.d("WS Connected: ","WS Connected");
                    _webSocketEventHandler.OnDisconnect(0,"");
                }

                @Override
                public void onConnected(WebSocket websocket, Map<String, List<String>> headers) throws Exception {
                    Log.d("WS DisConnected: ","WS DisConnected");
                    _webSocketEventHandler.OnConnect();
                }
            });

            _webSocket.connect();

        } catch (IOException e) {
            e.printStackTrace();
        } catch (WebSocketException e) {
            e.printStackTrace();
            _webSocketEventHandler.OnError(e);
        }

        return true;
    }

    public boolean SendPacket(String packet) {

        AppLogger.getInstance().Log(AppLogger.LogLevel.Debug,"Sending JSON Packet Server: "+ packet);

        if(_webSocket != null){
            _webSocket.sendText(packet);

            return true;
        }
        return false;
    }

    public boolean IsServerConnected() {
        //According websocket connection state return true or false
        return true;
    }

    public IWebSocketConnection get_webSocketEventHandler() {
        return _webSocketEventHandler;
    }
}
