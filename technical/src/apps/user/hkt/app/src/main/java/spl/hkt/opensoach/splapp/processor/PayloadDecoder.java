package spl.hkt.opensoach.splapp.processor;


import android.util.Log;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;

import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.model.PacketDecodeResultModel;
import spl.hkt.opensoach.splapp.model.communication.PacketAuthCodeDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketChartConfigurationModel;
import spl.hkt.opensoach.splapp.model.communication.PacketLocationDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.communication.PacketSimpleAckModel;

/**
 * Created by Mandar on 2/25/2017.
 */

public class PayloadDecoder {

    public static void Decode(PacketDecodeResultModel packetDecodeResultModel, String packet) {

        packetDecodeResultModel.IsSuccess = false;
        try {

            PacketChartConfigurationModel packetChartConfigurationModel;
            Log.i("PayloadDecoder", "CategoryID : " + packetDecodeResultModel.Packet.Header.Category + " CommandID : " + packetDecodeResultModel.Packet.Header.CommandID);

            switch (packetDecodeResultModel.Packet.Header.Ack){
                case 1: // Ack Response
                    Type packetType = new TypeToken<PacketModel<PacketSimpleAckModel>>() {
                    }.getType();
                    packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                    packetDecodeResultModel.Processor = new AcknowledgementProcessor();
                    packetDecodeResultModel.IsSuccess = true;
                    break;
                case 0:
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
                                    TypeToken<PacketModel<PacketLocationDataModel>> typeToken = new TypeToken<PacketModel<PacketLocationDataModel>>() {
                                    };
                                    packetType = typeToken.getType();
                                    packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                                    packetDecodeResultModel.Processor = new LocationDataProcessor();
                                    packetDecodeResultModel.IsSuccess = true;
                                    break;
                                }

                                case CommandConstants.CMD_CONFIG_CHART_CONFIG: {

                                    TypeToken<PacketModel<PacketChartConfigurationModel>> typeToken = new TypeToken<PacketModel<PacketChartConfigurationModel>>() {
                                    };
                                    packetType = typeToken.getType();
                                    packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                                    packetDecodeResultModel.Processor = new ChartDataProcessor();
                                    packetDecodeResultModel.IsSuccess = true;
                                    break;
                                }

                                case CommandConstants.CMD_CONFIG_SERVER_SYNC_COMPLETED: {
                                    //TODO
                                    break;
                                }

                                case CommandConstants.CMD_CONFIG_LOCATION_HCODE: {
                                    TypeToken<PacketModel<PacketAuthCodeDataModel>> typeToken = new TypeToken<PacketModel<PacketAuthCodeDataModel>>() {
                                    };
                                    packetType = typeToken.getType();
                                    packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                                    packetDecodeResultModel.Processor = new AuthCodeDataProcessor();
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
                            switch (packetDecodeResultModel.Packet.Header.CommandID) {
                                case CommandConstants.CMD_ACK_CHART_DATA:
                                case CommandConstants.CMD_ACK_DEVICE_REG: {
                                    packetType = new TypeToken<PacketModel<PacketSimpleAckModel>>() {
                                    }.getType();
                                    packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                                    packetDecodeResultModel.Processor = new AcknowledgementProcessor();
                                    packetDecodeResultModel.IsSuccess = true;
                                    break;
                                }

                                case CommandConstants.CMD_ACK_DEVICE_COMPLAINT_REGISTRATION: {
                                    //TODO
                                    break;
                                }

                                case CommandConstants.CMD_ACK_DEVICE_NOTIFICATION: {
                                    //TODO
                                    break;
                                }
                            }
                            break;
                        }

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
                    break;
            }

        } catch (Exception ex) {
            //result.Error =new ErrorModel();
            packetDecodeResultModel.IsSuccess = false;
            //TODO: Set Error Model
            //TODO: Log exception error
        }
    }
}
