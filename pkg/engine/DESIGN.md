# Deviant Engine Design

## Introduction

Deviant Engine houses all of the logic to validate and execute any actions that any client wants to perform it can be thought of loosely as the core 'engine' of the game.  Thomas itself is comprised of 3 smaller component packages actions, rules and test which I will outline the functionality of in this document. The Deviant Engine itself is vended out as a general purpose package for consumption in multiple applications, however it's primary application is within the deviant-instance-shard. It's important to note that due to the nature of Deviant being an MMO all client actions MUST be processed by the action and rules engines no matter what otherwise we risk compromising the integrity of the game.



## Design

### Rules Engine

The rules package contains sets of rules that cover turns, entity actions and encounters. Each rule itself is a Boolean function that validates a particular constraint that we would like to impose I.E. (entity must have X action points to perform Y action). When you combine all the rule sets within the rules package you should have an all encompassing view of Deviant's ruleset. The rules package also contains a processor for rules, this processor contains both the logic on how to process rules as well as the declarative definition of what rules should be processed when (our rulesets).



### Actions Engine

The action package contains a set of rules that covers turns, entity actions and encounters. Each action itself is a Boolean function that mutates the current game state in some way, I.E. (entity moves from x to y). By the time an action makes it to the actions engine it is expected to have completed all validation from the rules engine. If something is sent to the actions engine without having been validated by the rules engine all of our games 'anti-cheat' is disabled and arbitrary states in the game can occur. Similarly to the rules engine the actions engine contains action sets for each turn phase and encounter state. The action package also contains a processor for actions, the actions processor contains logic on how to process actions, how those actions being processed changes the game state/turn state flow and what to do when mutations to the game state dramatically modify the defined order of gameplay.



### Test Engine

The test package contains a set of mock game state data that is required to test both the rules and actions engine. This same mock data can be used by other packages to simply implement fake game instances and validate behavior. It also contains some exported helper functions to make testing rules/actions related state changes easier.