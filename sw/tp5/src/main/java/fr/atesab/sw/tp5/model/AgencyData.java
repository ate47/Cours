package fr.atesab.sw.tp5.model;

public class AgencyData {
    public final String agencyId;
    public final String agencyName;
    public final String agencyUrl;
    public final String agencyTimezone;
    public final String agencyLang;

    public AgencyData(String[] raw) {
        // agency_id,agency_name,agency_url,agency_timezone,agency_lang
        var i = 0;
        agencyId = raw[i++].replace(" ", "_");
        agencyName = raw[i++];
        agencyUrl = raw[i++];
        agencyTimezone = raw[i++];
        agencyLang = raw[i++];
    }
}
