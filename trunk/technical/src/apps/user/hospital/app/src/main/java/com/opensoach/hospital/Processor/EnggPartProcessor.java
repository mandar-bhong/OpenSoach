package com.opensoach.hospital.Processor;

import com.google.gson.JsonArray;
import com.google.gson.JsonElement;
import com.google.gson.JsonParser;
import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Helper.DataConverter.DataModelConverter;
import com.opensoach.hospital.Model.Communication.PacketEnggPartToolDataModel;
import com.opensoach.hospital.Model.Communication.PacketEnggPartToolsDataModel;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.DB.DBEnggPartTableRowModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Utility.AppLogger;

import java.util.List;

/**
 * Created by Mandar on 9/4/2017.
 */

public class EnggPartProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            PacketModel<PacketEnggPartToolsDataModel> packetEnggPartsDataModel = (PacketModel<PacketEnggPartToolsDataModel>)resultModel.Packet.Payload;
            List<PacketEnggPartToolDataModel> enggPartToolModels = packetEnggPartsDataModel.Payload.EnggPartTools;

            JsonParser parser = new JsonParser();
            JsonElement root = parser.parse(resultModel.JSONPacket);
            JsonArray enggPartToolJSON = root.getAsJsonObject().getAsJsonObject("payload").getAsJsonArray("engparts");

            Integer loopCount = 0;

            for (PacketEnggPartToolDataModel enggPartModel:enggPartToolModels) {

                String enggPartTools = enggPartToolJSON.get(loopCount).getAsJsonObject().get("tools").toString();

                DBEnggPartTableRowModel dbEnggPartTableRowModel = DataModelConverter.ConvertToDBEnggPart(enggPartModel,enggPartTools);

                DatabaseManager.InsertRow(dbEnggPartTableRowModel);

                loopCount++;
            }

            packetProcessResultModel.IsSuccess = true;

        }catch (Exception exeception){
            packetProcessResultModel.IsSuccess = false;
            AppLogger.getInstance().Log(exeception,"Error occured in EnggPartProcessor");
        }

        return packetProcessResultModel;
    }
}
