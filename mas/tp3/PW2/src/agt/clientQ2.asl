// Agent clientQ2 in project test

/* Initial beliefs and rules */
knowledge(limit_delay,plumbing, 50).
knowledge(penality, plumbing, 100).
action(update).
/* Initial goals */


/* Plans */
+!requestManagement(Action, limit_delay, Skill, Value)[source(Agent)]:
    action(Action) & knowledge(limit_delay, Skill, CurrentValue) & CurrentValue < Value
    <-
      -+knowledge(Action, Knowledge, Skill, Value);
      .send(Agent, achieve, inform_done);
      .

+!requestManagement(Action, penality, Skill, Value)[source(Agent)]:
    action(Action) & knowledge(penality, Skill, CurrentValue) & CurrentValue > Value
    <-
      -+knowledge(Action, Knowledge, Skill, Value);
      .send(Agent, achieve, inform_done);
      .

+!requestManagement(_, _, _, _)[source(Agent)]
  <-
    .send(Agent, achieve, failure);
    .

{ include("$jacamoJar/templates/common-cartago.asl") }
{ include("$jacamoJar/templates/common-moise.asl") }

// uncomment the include below to have an agent compliant with its organisation
//{ include("$moiseJar/asl/org-obedient.asl") }
