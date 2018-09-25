# gome

gome is a project intended to simplify IoT discovery and management, with the intended use being for the purposes of home automation.  At first it will be entirely terminal based, with an ASCII/ANSI based UI for running on linux systems.

As IoT devices could have a very wide range of features (e.g. a speaker light combo) the mix-in style interfaces of go might be well suited to composing a device as a collections of features instead of a more object inheritance style that Java would do.

## Concepts
### Spaces
- building
- floor
- room
    - 3 dimensional specs
    - curved walls
    - windows
    - doors
- outside

### IoT Devices
- any network device capable of sending/receiving communication
- can be positioned within spaces

## Server
In general there will need to be adapters that convert server input into inputs that specific devices will understand.

### IoT Device Discovery
Should be able to automatically detect devices on the network. Should also be able to manually add devices.

### REST Api
#### Spaces
- CRUD endpoints

#### IoT Devices
- CRUD endpoints
- interact with device
    - commonly used actions
    - allow pass through of commands

### DB Storage
- local db
- remote db

## Clients
### Terminal

### Web