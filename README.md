# Liquidator-Game
First fully functional version of my Arcade Shooter Game. You are a mercenary soldier guarding the border of a quarantined area. Do not let anyone through!

IMPORTANT! Leaderboards are currently running only on local host.

#
### Version Notes:
##### v1.0 beta
- First beta version
##### v0.9 alpha
- Implemented hiscore leaderboards
- Removed unused assets and bits of code
##### v0.8 alpha
- Made closing the program more convenient while in Setup launcher window
- Set target FPS back to 60
- Fixed a game breaking bug
#

# Running
Client:
```
go get Liquidator
go build
./Liquidator
```
Server:
```
cd backend
go build
./backend
```
### Controls:
- Arrow keys = movement
- Spacebar = fire weapon
- Numeric keys 1-5 = select weapon
- R = reload weapon
- P = pause game
- F1 = toggle frame limiter on/off
- End = commit suicide
- Esc = Exit game

### Credits:
- Valiant (Programming, Art & Game Design)
- xXAshuraXx (Art)
- Jarpdzonson (Art)
- SuperPhat (Sound FX)
