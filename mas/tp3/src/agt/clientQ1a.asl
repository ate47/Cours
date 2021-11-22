// Agent clientQ1a

{ include("common_clientQ1.asl") }

/* Initial beliefs and rules */
// delay_limit(v): the total client’s delay tolerance in v time units (Initial belief)
delay_limit(10).
// ???
penalty(plumbing, 50).
// delay(d): delay of d time units needed by the client (Received belief)

/* Initial goals */

/* Plans */

+!updateTolerance(V) :
    delay_limit(L) & L > V
	<-
	.print("update tolerance");
	-+delay_limit(L-V).
   
{ include("$jacamoJar/templates/common-cartago.asl") }
{ include("$jacamoJar/templates/common-moise.asl") }

// uncomment the include below to have an agent compliant with its organisation
//{ include("$moiseJar/asl/org-obedient.asl") }
