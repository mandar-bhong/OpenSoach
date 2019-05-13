package com.opensoach.hpft.Manager;

import java.util.HashMap;

import com.opensoach.hpft.Model.Communication.CommandRequest;

/**
 * Created by Mandar on 3/2/2017.
 */

public class RequestManager {

    private static RequestManager singleton;

    private int _currentId = 0;
    private HashMap<Integer, CommandRequest> _requestMap;

    /* Static 'instance' method */
    public static RequestManager Instance() {
        if (singleton == null)
            singleton = new RequestManager();
        return singleton;
    }

    private RequestManager() {
        _requestMap = new HashMap<Integer, CommandRequest>();
    }

    public int GenerateRequestID() {
        synchronized (this) {
            _currentId++;

            boolean isRequestIdGenerated = false;
            do {
                if (!_requestMap.containsKey(_currentId)) {
                    isRequestIdGenerated = true;
                    //_requestMap[_currentId] = request;
                } else {
                    if (_currentId == Integer.MAX_VALUE)
                        _currentId = 1;
                    else
                        _currentId++;
                }
            }
            while (!isRequestIdGenerated);

            return _currentId;
        }
    }

    public void AddRequest(int requestID, CommandRequest packetPayloadModel) {

        _requestMap.put(requestID, packetPayloadModel);
    }

    public void CompleteRequest(int requestID) {
        if (_requestMap.containsKey(requestID))
            _requestMap.remove(requestID);
    }

    public CommandRequest GetRequest(int requestID) {

        if(_requestMap.containsKey(requestID)){
            return  _requestMap.get(requestID);
        }
        return null;
    }


}
