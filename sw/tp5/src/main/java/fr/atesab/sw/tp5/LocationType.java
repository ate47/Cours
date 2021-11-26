package fr.atesab.sw.tp5;

import org.apache.jena.rdf.model.Model;
import org.apache.jena.rdf.model.Resource;

public class LocationType {
    public static Resource[] buildMap(String extNamespace, Model model) {
        var resources = new Resource[5];

        resources[0] = model.createResource(extNamespace + "Stop");
        resources[1] = model.createResource(extNamespace + "Station");
        resources[2] = model.createResource(extNamespace + "Entrance");
        resources[3] = model.createResource(extNamespace + "GenericNode");
        resources[4] = model.createResource(extNamespace + "BoardingArea");

        return resources;
    }

    private LocationType() {
    }
}
