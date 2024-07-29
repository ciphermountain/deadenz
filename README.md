# DeadEnz
A random situation generator and game

## Components

### Core
The core of the game functions as a state mutator. Old state is provided with a game command
to mutate the state and produce an effect. No game state is saved as this is the responsibility
of a game client to do.

The core game service is only responsible for game assets and state mutation mechanics. Running
as a service allows hot loading of game assets.

#### Run the Core Service

The core service can run with or without a multiverse server where a multiverse service simply
satisfies the multiverse interface. Using a multiverse service with core, the multiverse service
must be started before the core service.

*Basic Core Service*

```
$ deadenz run core -l 127.0.0.1 -p 8000
```

*Connect to Multiverse Service*

```
$ deadenz run core -l 127.0.0.1 -p 8000 --with-multiverse --multiverse-host=127.0.0.1:9001
```

### Client
A client should interact with the core game service and is responsible for feedback to the user,
user state tracking, etc.

### Multiverse
The multiverse server addon tracks and distributes multiverse events.

## Console Version
The default version of the game provided runs directly on a console and is currently
a single player game. This version currently supports the basic functions: spawnin,
walk, backpack, xp, and currency.

## Game Commands

### Spawnin
This is the entry point of the game.
