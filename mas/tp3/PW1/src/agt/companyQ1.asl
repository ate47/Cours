// Agent companyQ1

/* Initial beliefs and rules */
// client(c, v): allocation of v time units to the client c (Initial belief)
client(giacomo, 20).
client(andrei, 30).
// ????
skill(plumbing).
// delay_limit(d): maximum delay tolerance of d time units for the execution of a task (Received belief)

/* Initial goals */

!start.

/* Plans */

+!needDelay(S,D) : skill(X) & S == X & client(C,V)
  <-
  !proposeDelay(C,S,D).


+delay_limit(LimitCompany)[source(Client)]
  <-
      if (LimitClient > LimitCompany) {
        // Q: if the delay limit is greater than the delay needed by the company, 
        //    then request the delay to the client (???)
        .send(Client, tell, delay(LimitCompany));
      } elif (LimitClient < LimitCompany) {
        // Q: if the delay limit is less than the delay needed by the company, 
        //    print an error message informing failure in completing the task
        println("Can't complete the task for ", Client, 
                ", the delay needed by the company is too important Client=", 
                LimitClient, " < Company=", LimitCompany);
      } // no = case asked
      .

+!start : true
  <- 
    for (client(Client, LimitCompany)) {
      // Q: requests to the client what is its acceptable delay
      //    (i.e., delay_limit belief) to perform the assigned task
      .send(Client, askOne, delay_limit(X), delay_limit(L));
      // use the sync version to be sure all the delay_limit are here before the end
      +delay_limit(L)[source(Client)];
    };
    // Update the companyQ1.asl file so that the company sends a delay request to 
    // all clients and choose the client with greatest delay_limit
    .findall(
      delay_limit(LimitClient)[source(Client)], 
      delay_limit(LimitClient)[source(Client)],
      ListClient
    );
    .max(ListClient, delay_limit(LimitClient)[source(Client)]);
    println("I choose ", Client, " his delay_limit is ", LimitClient);
    .

{ include("$jacamoJar/templates/common-cartago.asl") }
{ include("$jacamoJar/templates/common-moise.asl") }

// uncomment the include below to have an agent compliant with its organisation
//{ include("$moiseJar/asl/org-obedient.asl") }
