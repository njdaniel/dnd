# dnd 

Initial goal is to provide a tool for helping dm a pathfinder/dnd 
role-playing session.

With new release of Pathfinder2.0 and my own plethora of house rules I 
wanted to be able to reference/search/list the rules, items, traits, spells, feats, creatures.
But from my own forked version. 
Build trader inventories at an instance, create new npc's, and encounters.
Keep track of attacking monsters, and npc's created.
Secret communication of players, access to players info and help with new players
immerse into the role without worrying about the math and game mechanics.

But small steps first :camel:

## Design

There were a few ways I thought about approaching this design: 
traditional SPA-REST-DB stack, scrape the pathfinder2.0 data push into a database
and run db queries and have a REST API wrapping it to interface with.
Once I scraped some data to test with and added some house rule changes, 
I was left with a directory of files that were actually quite usable for myself. 
Using grep and jq I could pretty much find the data I was looking for. Might be 
a little long to type but I could create an alias for the cmds maybe..
Or I could use parse with go and run simplified commands for what I need. 
I could have a CLI tool to simplify the cmds and leave the design more decoupled for 
the other desired features. The rules and information can be stored and versioned 
as directories and files. 

With the data being decoupled from the commands. I wanted the same ease of access of 
data with parsing through directories with linux. So without ETLing the data into a 
database I wanted to call the cmds to search the data remotely. The interaction between 
the two could be .... grpc? Also a good excuse to learn some grpc deploy to k8s cluster 
and accomplish the goals. 



## How to Contribute

contact me through slack 
https://join.slack.com/t/rtp-gophers/shared_invite/enQtNTE3NjIyMTgyODgyLWEyZWVkYWE0OGE3ZmRjNTFkNWJkODdiOWJjY2JjMDcyZTJlNmJiNjU0YjFlN2NmMTNjNGVjY2FmMGUzYzk5NTg

Message @Nick