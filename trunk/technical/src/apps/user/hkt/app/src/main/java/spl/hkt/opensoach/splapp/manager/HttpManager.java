package spl.hkt.opensoach.splapp.manager;

import android.os.AsyncTask;

import com.google.gson.Gson;
import com.google.gson.JsonElement;
import com.google.gson.JsonParser;

import okhttp3.MediaType;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;
import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.helper.AppHelper;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.logger.AppLogger;
import spl.hkt.opensoach.splapp.model.communication.APIResponseAuthDataModel;
import spl.hkt.opensoach.splapp.model.communication.APIResponseModel;

public class HttpManager {
    public static void ProcessWebSocketURL(String authURL, String deviceSerialNumber) {
        HttpHandler httpHandler = new HttpHandler();
        httpHandler.execute(authURL,deviceSerialNumber);
    }
}


class HttpHandler extends AsyncTask {

    OkHttpClient client = new OkHttpClient();

    @Override
    protected Object doInBackground(Object[] objects) {
        Request.Builder builder = new Request.Builder();

        MediaType JSON = MediaType.parse("application/json; charset=utf-8");
        String jsonData =  PacketHelper.GetAPIAuthRequestJson((String) objects[1]);
        //String jsonData =  PacketHelper.GetAPIAuthRequestJson(Build.SERIAL);
        RequestBody body = RequestBody.create(JSON, jsonData);

        Request request = builder.addHeader("Content-Type", "application/json")
                .post(body)
                .url((String) objects[0])
                .build();

        try {
            Response response = client.newCall(request).execute();

            switch (response.code()) {
                case 200:
                    String responseJson = response.body().string();
                    APIResponseModel responseModel = new Gson().fromJson(responseJson, APIResponseModel.class);

                    if (responseModel.IsSuccess == true) {
                        JsonParser parser = new JsonParser();
                        JsonElement root = parser.parse(responseJson);
                        String authJSON = root.getAsJsonObject().get("data").toString();

                        APIResponseAuthDataModel responseDataModel = new Gson().fromJson(authJSON, APIResponseAuthDataModel.class);

                        return responseDataModel;

                    } else {
                        Thread.sleep(3 * 1000);
                        //HttpManager.ProcessWebSocketURL();
                    }
                    break;
                default:
                    Thread.sleep(5 * 1000);
                    //HttpManager.ProcessWebSocketURL();
                    break;
            }
        } catch (Exception e) {
            e.printStackTrace();
        }

        return null;
    }

    @Override
    protected void onPostExecute(Object o) {
        super.onPostExecute(o);

        try {
            if (o == null) {
                AppLogger.getInstance().Log(AppLogger.LogLevel.Error, "HttpManager:onPostExecute: Received object is null");
                return;
            }

            APIResponseAuthDataModel responseDataModel = (APIResponseAuthDataModel) o;
            AppRepo.getInstance().setServerWebSocketURL(responseDataModel.LocationURL);
            AppRepo.getInstance().setAuthToken(responseDataModel.Token);
            AppHelper.OnWebSocketURLReceived();

        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
