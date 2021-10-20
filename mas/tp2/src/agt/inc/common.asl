/* auxiliary rules for agents */

i_am_winning(Art)   // check if I placed the current best bid on auction artifact Art
   :- currentWinner(W)[artifact_id(Art)] &
      .my_name(Me) & .term2string(Me,MeS) & W == MeS.

/* auxiliary plans for Cartago */

+!start : true
   <- 
      // send to giacomo the available tasks
      for ( my_task(Task) ) {
         .send(giacomo, tell, my_task(Task));
      };
      .

+!auction_to_use(Task, ArtName)
   <-
      // find art id and focus on it
      lookupArtifact(ArtName,ArtId);
      focus(ArtId);
      .
