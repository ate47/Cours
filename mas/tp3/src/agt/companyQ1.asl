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
      

+!start : true
  <- 
    for (client(Ag, Lc)) {
      // requests to the client what is its acceptable delay (i.e., delay_limit belief, wtf in english?) to perform the assigned task
      .send(Ag, askOne, delay_limit(X), delay_limit(L));
      if ( L > Lc ) {
        // if the delay limit is greater than the delay needed by the company, then request the delay to the client
        // delay already in delay_limit(L) so useless to request it? TODO: check
      } elif ( L < Lc  ) {
        // if the delay limit is less than the delay needed by the company, print an error message informing failure in completing the task
        println("Can't complete the task for ", Ag, ", the delay needed by the company is too important Client=", L, " < Company=", Lc);
      } // no = case asked
    };
    .


!start.

{ include("$jacamoJar/templates/common-cartago.asl") }
{ include("$jacamoJar/templates/common-moise.asl") }

// uncomment the include below to have an agent compliant with its organisation
//{ include("$moiseJar/asl/org-obedient.asl") }
