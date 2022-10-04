# Alien invasion
## Description
Mad aliens are about to invade the earth and there is a task to simulate the invasion.
## Configuration
Configuration is done through the environment variables. Supported configuration parameters:  
- ```LOG_LEVEL``` - determines the logs level of the [logrus](https://github.com/sirupsen/logrus) logger. Could have values such as  ```panic```, ```fatal```, ```error```,  ```warn```, ```info```, ```debug```, ```trace```. By default it is set to ```debug```.
- ```MAP_FILE_PATH``` - determines the path to the map text file.
- ```ALIEN_MAX_STEPS_NUMBER``` - determines the number of max steps for alien before he could stop moving.
## Usage
### Command line
Direct usage could be done with the following command:  

```go run main.go --aliens-count 2 --env ./config/.env.dev```  

```--aliens-count``` (shot version is ```-c```) parameter should take number of aliens participated in the simulation.  
```--env``` (shot version is ```-e```) parameter should take the path to the config file.
### Makefile
Use makefile for a quick run of the program. There are the following rules:
- ```run``` - runs simulation with 2 aliens.  
- ```test``` - runs unit-tests.
- ```godoc``` - generates docs and launches it on http://localhost:6060.  
- ```lint``` - runs linting.  