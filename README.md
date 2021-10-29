# dnd 

This is helper tool for a players and DMs to help with flow of the sessions. 

Currently this uses the rules of my current 
[homebrew](https://github.com/njdaniel/dnd/wiki) rules. The spirit of the game is the simular but with a more focus on actually 
playing the game and getting emerged into the characters and the story without worrying about the rules except of the DM. 
The DM oversees the game and enforces the laws of the games physics. Character sheets have two versions the PC and DM. 
The PC version has the description of the character their motivation, personality, items, past history, knowledge, and 
overall skills. The DM version has the math and and details for help running the game.

#### Features:
* Create random character
* Roll dice
* save character to file
* CLI character random generator
* CLI create character

#### Upcoming:
* CLI create store

#### Goals
* Generate stores, towns
* Save the generated files to shared database


#### Quick Use CLI

```console
$ git clone https://github.com/njdaniel/dnd.git
$ cd dnd
$ go build
$ ./dnd 
$ #prints out help of available cmds
```

##### Create New Character

```console
$ ./dnd character create
```

##### Rolling Dice

```console 
$ ./dnd roll 3d6
$ [5,4,1]
$ 
$ ./dnd roll d20
$ 14
$ 
$ Explodes on default 6
$ ./dnd roll d6!
$ [6,3] 
$
$ Keep highest 2 of 4d6
$ ./dnd roll 4kh2d6
$ [5,4] 
$ 
$ Keep lowest 2 of 4d6
$ ./dnd roll 4kl2d6
$ [1,3]
```



#### Homebrew Rules
[wiki](https://github.com/njdaniel/dnd/wiki)

#### How to Contribute

[contributing](contributing.md)
