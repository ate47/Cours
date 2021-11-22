# 19/11/2021 MAS Report

## Links

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

## 3.1.1 <!-- omit in toc -->

> What are the hidden hypothesis of coordination by messages like in this scenario?

TODO: Fill this

## 3.1.2 <!-- omit in toc -->

> Discuss the pros and cons of this approach versus the agent-centric coordination approach.

TODO: Fill this

## 3.1.3 <!-- omit in toc -->

> What is missing in the description of coordination by messages? Why?

TODO: Fill this

# 3.2 Management of delay for performing a task

## 3.2.6 <!-- omit in toc -->

I was already knowing their behaviors from [the doc](http://jason.sourceforge.net/doc/api/jason/stdlib/send.html), because we read it last time.

- `tell` add a belief to another agent base.
- `achieve` add an event to another agent event queue.
- `askOne` ask a belief from another agent base, can be synchronously executed or not depending on if we add a 4th argument.

# 4.1 QUESTION

## 4.1.1 <!-- omit in toc -->

> What are the pros and cons of coordination with protocol?

**Pro**: Everybody can join or leave the system if they respect the protocol.

**Cons**: The agents are restricted in their communications.
