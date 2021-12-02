// Agent clientQ2 in project test

/* Initial beliefs and rules */
knowledge(limit_delay,plumbing, 50).
knowledge(penality, plumbing, 100).
action(update).
/* Initial goals */


/* Plans */
+requestManagement(Action, Knowledge, Skill, Value)
    <-
      .

{ include("$jacamoJar/templates/common-cartago.asl") }
{ include("$jacamoJar/templates/common-moise.asl") }

// uncomment the include below to have an agent compliant with its organisation
//{ include("$moiseJar/asl/org-obedient.asl") }
