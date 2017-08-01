# countdown-timer
A small webapp to create a countdown to an event.

# usage

Create the app with:

```
make
```

And export `COUNTDOWN_DESTINATION_TIMER` and `COUNTDOWN_DESTINATION_EVENT` before running the app:

```
export COUNTDOWN_DESTINATION_TIMER='Jul 4, 2018 at 8:00pm (EST)'
export COUNTDOWN_DESTINATION_EVENT='Independence Day!'
./countdown-timer
```

Or create a Docker container with:

```
make container
```
