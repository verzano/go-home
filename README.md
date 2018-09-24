# gome

gome is a project intended to simplify IoT discovery and management, with the intended use being for the purposes of home automation.  At first it will be entirely terminal based, with an ASCII/ANSI based UI for running on linux systems.

As IoT devices could have a very wide range of features (e.g. a speaker light combo) the mix-in style interfaces of go might be well suited to composing a device as a collections of features instead of a more object inheritance style that Java would do.

## Ideas
- create commonly used commands (i.e. turn on/off, light settings, device status)
    - will need adapters for common systems, plus a nice interface for implementing new ones
- allow for custom commands to be sent through to specific devices
- DB for data storage
    - local DB
    - cloud DB
- REST api
- auto detection of IoT devices
- manual adding of IoT devices
- room/home layout mapper
    - doors/windows
    - 3 dimensional room info
    - curved walls
    - maybe just allow a formula to define a room???
    - single room first
    - single floor
    - multi-floor
    - non connected rooms
    - non connected buildings
- allow positioning of IoT devices within rooms