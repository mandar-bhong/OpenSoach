package com.opensoach.hospital.Communication;

/**
 * Created by Mandar on 2/26/2017.
 */

public interface IWebSocketConnection {

    void OnConnect();

    void OnMessage(String strRespone);

    void OnDisconnect(int code, String reason);

    void OnError(Exception error);
}
