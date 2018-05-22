package spl.hkt.opensoach.splapp.model.view;

import java.lang.reflect.Field;
import java.util.ArrayList;
import java.util.HashMap;

/**
 * Created by Mandar on 3/27/2017.
 */

public class ChartConfigModel
{
    private int chartId;
    private String chartName;
    private int serverChartId;
    private int locationId;
    private int slotInterval;
    private HashMap<Integer, ChartConfigSlotModel> slots;
    private HashMap<Integer, ChartConfigTaskModel> tasks;

    public ChartConfigModel(){
        slots = new HashMap<>();
        tasks =new HashMap<>();
    }

    public int getChartId() {
        return chartId;
    }

    public void setChartId(int chartId) {
        this.chartId = chartId;
    }

    public String getChartName() {
        return chartName;
    }

    public void setChartName(String chartName) {
        this.chartName = chartName;
    }

    public int getServerChartId() {
        return serverChartId;
    }

    public void setServerChartId(int serverChartId) {
        this.serverChartId = serverChartId;
    }

    public int getLocationId() {
        return locationId;
    }

    public void setLocationId(int locationId) {
        this.locationId = locationId;
    }

    public int getSlotInterval() {
        return slotInterval;
    }

    public void setSlotInterval(int slotInterval) {
        this.slotInterval = slotInterval;
    }

    public HashMap<Integer, ChartConfigSlotModel> getSlots() {
        return slots;
    }

    public void setSlots(HashMap<Integer, ChartConfigSlotModel> slots) {
        this.slots = slots;
    }

    public HashMap<Integer, ChartConfigTaskModel> getTasks() {
        return tasks;
    }

    public void setTasks(HashMap<Integer, ChartConfigTaskModel> tasks) {
        this.tasks = tasks;
    }
}
