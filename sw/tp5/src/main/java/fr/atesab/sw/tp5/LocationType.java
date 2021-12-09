package fr.atesab.sw.tp5;

import org.apache.jena.rdf.model.Model;
import org.apache.jena.rdf.model.Resource;

public record LocationType(int id, Resource resource) {
    /**
     * create an array mapping the id of https://gtfs.org/reference/static/#stopstxt
     * location_type to a resource
     * 
     * @param extNamespace the ext namespace
     * @param model        the model to create resources
     * @return the array
     */
    public static LocationType[] buildMap(String extNamespace, Model model) {
        var resources = new LocationType[5];

        resources[0] = new LocationType(0, model.createResource(extNamespace + "Stop"));
        resources[1] = new LocationType(1, model.createResource(extNamespace + "Station"));
        resources[2] = new LocationType(2, model.createResource(extNamespace + "Entrance"));
        resources[3] = new LocationType(3, model.createResource(extNamespace + "GenericNode"));
        resources[4] = new LocationType(4, model.createResource(extNamespace + "BoardingArea"));

        return resources;
    }
}
