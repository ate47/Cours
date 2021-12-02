# 19/11/2021 MAS Report

## Links

- [Eclipse .gitignore](https://github.com/github/gitignore/blob/master/Global/Eclipse.gitignore) and [Gradle .gitignore](https://github.com/github/gitignore/blob/master/Gradle.gitignore) (can be useful to the reader)
- [Course](https://ci.mines-stetienne.fr/cps2/mac/)
- [Practical 3](https://ci.mines-stetienne.fr/cps2/mac/usecase-interaction-centric.html)

## Table of contents

- [19/11/2021 MAS Report](#19112021-mas-report)
  - [Links](#links)
  - [Table of contents](#table-of-contents)
- [3.1 QUESTIONS](#31-questions)
- [3.2 Management of delay for performing a task](#32-management-of-delay-for-performing-a-task)
- [4.1 QUESTION](#41-question)

# 3.1 QUESTIONS

<!-- if you are reading this, I have no idea how to answer those questions, so I've put random answers -->

## 3.1.1 <!-- omit in toc -->

> What are the hidden hypothesis of coordination by messages like in this scenario?

Messages are handled by the agents

## 3.1.2 <!-- omit in toc -->

> Discuss the pros and cons of this approach versus the agent-centric coordination approach.

- **pros**:
  - No need to have an agent in the middle to handle all the tasks.
- **cons**:
  - The messages should follow the same protocol.

## 3.1.3 <!-- omit in toc -->

> What is missing in the description of coordination by messages? Why?

How the whole set of agents should behave in common.

# 3.2 Management of delay for performing a task

## 3.2.6 <!-- omit in toc -->

I was already knowing their behaviors from [the doc](http://jason.sourceforge.net/doc/api/jason/stdlib/send.html), because we read it last time.

- `tell` add a belief to another agent base.
- `achieve` add an event to another agent event queue.
- `askOne` ask a belief from another agent base, if no 4th argument is put, the task is asynchronous and it will trigger the event `+the_belief(...)` when the belief is received, if a 4th argument is added, the task is synchronous and will unify the belief with the 4th argument, we can set a timeout in ms if a 5th argument is added.

# 4.1 QUESTION

## 4.1.1 <!-- omit in toc -->

> What are the pros and cons of coordination with protocol?

**Pro**: Everybody can join or leave the system if they respect the protocol.

**Cons**: The agents are restricted in their communications.
