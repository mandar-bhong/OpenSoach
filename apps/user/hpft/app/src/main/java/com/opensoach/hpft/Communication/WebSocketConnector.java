package com.opensoach.hpft.Communication;

import android.content.Context;
import android.os.Handler;
import android.util.Log;

import org.apache.http.message.BasicNameValuePair;

import java.net.URI;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by samir.s.bukkawar on 2/18/2017.
 */

public class WebSocketConnector {
    private WebSocketClient mClient;
    private Context mContext;
    private final String TAG = "WebSocketConnector";
    private WebSocketThread mWebSocketThread;
    private IWebSocketConnection eventHandler;


    public WebSocketConnector(IWebSocketConnection handler) {
        eventHandler = handler;
        //URI testUri = URI.create(Constants.WEB_SOCKET_URL);
    }

    public void connectWebSocket(URI testUri) {
        // List<BasicNameValuePair> extraHeaders = Arrays.asList(new BasicNameValuePair("Cookie", "session=abcd"));\
        List<BasicNameValuePair> extraHeaders = new ArrayList<BasicNameValuePair>();

        mClient = new WebSocketClient(testUri, new WebSocketClient.Listener() {

            @Override
            public void onConnect() {
                Log.d(TAG, "Connected!");
                eventHandler.OnConnect();
            }

            @Override
            public void onMessage(String strRespone) {
                Log.d(TAG, String.format("Got string message : " + strRespone));

                eventHandler.OnMessage(strRespone);

               /* mWebSocketThread.handler.obtainMessage(1, strRespone);

                Message message = mWebSocketThread.handler.obtainMessage(1, WebSocketThread.responseString);
                Bundle b = new Bundle();
                b.putString("RESPONSE_MESSAGE", strRespone);
                message.setData(b);
                mWebSocketThread.handler.sendMessage(message);*/
            }

            @Override
            public void onMessage(byte[] data) {
                Log.d(TAG, String.format("Got binary message! %s", data.toString()));
            }

            @Override
            public void onDisconnect(int code, String reason) {
                Log.d(TAG, String.format("Disconnected! Code: %d Reason: %s", code, reason));
                eventHandler.OnDisconnect(code, reason);
            }

            @Override
            public void onError(Exception error) {
                Log.e(TAG, "Error : " + error);
                eventHandler.OnError(error);
            }
        }, extraHeaders);

        mClient.connect();
    }

    public void sendWSData(final String data1, final String data2) {
        final Handler handler = new Handler();
        handler.postDelayed(new Runnable() {
            @Override
            public void run() {

                //Device autheticate
                mClient.send(data1);
                mClient.send(data2);
            }
        }, 1000);
    }

    public void sendWSData(String packet) {
        mClient.send(packet);
    }
}

