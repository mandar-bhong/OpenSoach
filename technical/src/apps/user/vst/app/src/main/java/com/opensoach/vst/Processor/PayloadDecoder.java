package com.opensoach.vst.Processor;


import android.util.Log;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.ArrayList;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.Communication.PacketCardListConfigurationModel;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketLocationDataModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketSimpleAckModel;

/**
 * Created by Mandar on 2/25/2017.
 */

public class PayloadDecoder {

    public static void Decode(PacketDecodeResultModel packetDecodeResultModel, String packet) {

        packetDecodeResultModel.IsSuccess = false;
        try {

            Log.i("PayloadDecoder", "CategoryID : " + packetDecodeResultModel.Packet.Header.Category + " CommandID : " + packetDecodeResultModel.Packet.Header.CommandID);
            Type packetType;

            switch (packetDecodeResultModel.Packet.Header.Category) {
                case CommandConstants.CMD_CAT_DEVICE_REG: {
                    switch (packetDecodeResultModel.Packet.Header.CommandID) {
                        case CommandConstants.CMD_DEVICE_REGISTRATION: {
                            //TODO
                            break;
                        }
                    }

                    break;
                }
                case CommandConstants.CMD_CAT_CONFIG: {
                    switch (packetDecodeResultModel.Packet.Header.CommandID) {

                        case CommandConstants.CMD_CONFIG_DEVICE_SYNC_COMPLETED: {
                            //TODO
                            break;
                        }

                        case CommandConstants.CMD_CONFIG_LOCATION_SYNC: {
                            TypeToken<PacketModel<PacketLocationDataModel[]>> typeToken = new TypeToken<PacketModel<PacketLocationDataModel[]>>() {
                            };
                            packetType = typeToken.getType();
                            packetDecodeResultModel.Packet = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new LocationDataProcessor();
                            packetDecodeResultModel.IsSuccess = true;
                            break;
                        }

//                        case CommandConstants.CMD_CONFIG_CHART_CONFIG: {
//
//                            TypeToken<PacketModel<ArrayList<PacketChartConfigurationModel>>> typeToken = new TypeToken<PacketModel<ArrayList<PacketChartConfigurationModel>>>() {
//                            };
//                            packetType = typeToken.getType();
//                            packetDecodeResultModel.Packet = new Gson().fromJson(packet, packetType);
//                            packetDecodeResultModel.Processor = new ChartDataProcessor();
//                            packetDecodeResultModel.IsSuccess = true;
//                            break;
//                        }

                        case CommandConstants.CMD_CONFIG_CARD_LIST_CONFIG:{
                            TypeToken<PacketModel<ArrayList<PacketCardListConfigurationModel>>> typeToken = new TypeToken<PacketModel<ArrayList<PacketCardListConfigurationModel>>>() {
                            };
                            packetType = typeToken.getType();
                            packetDecodeResultModel.Packet = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new CardListProcessor();
                            packetDecodeResultModel.IsSuccess = true;
                            break;
                        }

                        case CommandConstants.CMD_CONFIG_SERVER_SYNC_COMPLETED: {
                            //TODO
                            break;
                        }


                        case CommandConstants.CMD_CONFIG_LOCATION_HCODE: {
                            TypeToken<PacketModel<ArrayList<String>>> typeToken = new TypeToken<PacketModel<ArrayList<String>>> () {
                            };
                            packetType = typeToken.getType();
                            packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new AuthCodeDataProcessor();
                            packetDecodeResultModel.IsSuccess = true;
                            break;
                        }

                        case CommandConstants.CMD_CONFIG_LOCATION_AUTH_CODE_ASSOCIATED:
                        case CommandConstants.CMD_CONFIG_LOCATION_AUTH_CODE_ADDED:{
                            TypeToken<PacketModel<ArrayList<String>>> typeToken = new TypeToken<PacketModel<ArrayList<String>>> () {
                            };
                            packetType = typeToken.getType();
                            packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new AuthCodeAddedProcessor();
                            packetDecodeResultModel.IsSuccess = true;

                            break;
                        }

                        case CommandConstants.CMD_CONFIG_LOCATION_AUTH_CODE_REMOVED:{
                            TypeToken<PacketModel<ArrayList<String>>> typeToken = new TypeToken<PacketModel<ArrayList<String>>> () {
                            };
                            packetType = typeToken.getType();
                            packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new AuthCodeRemovedProcessor();
                            packetDecodeResultModel.IsSuccess = true;
                            break;
                        }
                    }
                    break;
                }
                case CommandConstants.CMD_CAT_DATA: {
                    switch (packetDecodeResultModel.Packet.Header.CommandID) {
                        case CommandConstants.CMD_DATA_CHART_DATA: {
                            //TODO
                            break;
                        }
                        case CommandConstants.CMD_DATA_COMPLAINT_DATA: {
                            //TODO
                            break;
                        }
                    }
                    break;

                }
                case CommandConstants.CMD_CAT_ACK: {
                    packetType = new TypeToken<PacketModel<PacketSimpleAckModel>>() {
                    }.getType();
                    packetDecodeResultModel.Packet = new Gson().fromJson(packet, packetType);
                    CommandRequest request =  RequestManager.Instance().GetRequest(packetDecodeResultModel.Packet.Header.SeqID);
                    if(request!=null) {
                        packetDecodeResultModel.Processor = request.AckProcessor;
                    }
                    packetDecodeResultModel.IsSuccess = true;
                }
                break;

                case CommandConstants.CMD_CAT_DEVICE_STATUS: {
                    switch (packetDecodeResultModel.Packet.Header.CommandID) {

                        case CommandConstants.CMD_DEVICE_STATUS_BATTERY_STAUS: {
                            //TODO
                            break;
                        }
                    }
                    break;
                }
            }


        } catch (Exception ex) {
            //result.Error =new ErrorModel();
            packetDecodeResultModel.IsSuccess = false;
            //TODO: Set Error Model
            //TODO: Log exception error
        }
    }
}
