// Agent companyQ1

/* Initial beliefs and rules */
client(giacomo, 20).
client(andrei, 30).
skill(plumbing).
/* Initial goals */

/* Plans */

+!needDelay(S,D) : skill(X) & S == X & client(C,V)
  <-
  !proposeDelay(C,S,D).
      

{ include("$jacamoJar/templates/common-cartago.asl") }
{ include("$jacamoJar/templates/common-moise.asl") }

// uncomment the include below to have an agent compliant with its organisation
//{ include("$moiseJar/asl/org-obedient.asl") }
