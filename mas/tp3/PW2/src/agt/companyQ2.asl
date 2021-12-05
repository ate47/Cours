// Agent companyQ2

/* Initial beliefs and rules */
client(giacomo, 20).
client(andrei, 30).
skill(plumbing).
delay_limit(giacomo, 30).
/* Initial goals */

!start.

/* Plans */

+!start : client(C,V) & skill(S) & delay_limit(C, D)
  <-
  .print("sending a request");
  .send(C, achieve, requestManagement(update,limit_delay,S,D)).

+!inform_done
  <-
    .print("inform_done")
    .

+!failure
  <-
    .print("failure")
    .

{ include("$jacamoJar/templates/common-cartago.asl") }
{ include("$jacamoJar/templates/common-moise.asl") }

// uncomment the include below to have an agent compliant with its organisation
//{ include("$moiseJar/asl/org-obedient.asl") }
