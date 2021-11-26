package fr.atesab.sw.tp5;

import java.util.Arrays;

public class StopData {
    public final String stopId;
    public final String stopName;
    public final String stopDesc;
    public final String stopLat;
    public final String stopLon;
    public final String zoneId;
    public final String stopUrl;
    public final String locationType;
    public final String parentStation;

    public StopData(String[] raw) {
        // stop_id,stop_name,stop_desc,stop_lat,stop_lon,zone_id,stop_url,location_type,parent_station
        var i = 0;
        stopId = raw[i++].replace(" ", "_");
        stopName = raw[i++];
        stopDesc = raw[i++];
        stopLat = raw[i++];
        stopLon = raw[i++];
        zoneId = raw[i++];
        stopUrl = raw[i++];
        locationType = raw[i++];
        parentStation = raw[i++];
    }
}
