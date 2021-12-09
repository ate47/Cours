package fr.atesab.sw.tp5.model;

public class RouteData {
    public final String routeId;
    public final String agencyId;
    public final String routeShortName;
    public final String routeLongName;
    public final String routeDesc;
    public final String routeType;
    public final String routeUrl;
    public final String routeColor;
    public final String routeTextColor;

    public RouteData(String[] raw) {
        // route_id,agency_id,route_short_name,route_long_name,route_desc,route_type,route_url,route_color,route_text_color
        var i = 0;
        routeId = raw[i++].replace(" ", "_");
        agencyId = raw[i++];
        routeShortName = raw[i++];
        routeLongName = raw[i++];
        routeDesc = raw[i++];
        routeType = raw[i++];
        routeUrl = raw[i++];
        routeColor = raw[i++];
        routeTextColor = raw[i++];
    }
}
