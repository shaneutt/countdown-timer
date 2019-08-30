# countdown-timer
A small webapp to create a countdown to an event.
Default running port 8080 - you can change it in main.go file.

# usage

Create the app with:

```
make
```

And export `COUNTDOWN_DESTINATION_TIME` and `COUNTDOWN_DESTINATION_EVENT` before running the app:

```
export COUNTDOWN_DESTINATION_TIME='Jul 4, 2018 at 8:00pm (EST)'
export COUNTDOWN_DESTINATION_EVENT='Independence Day!'
./countdown
```

# Docker usage:

If you want to edit web layout
```
vi countdown-time.html
```

Create docker image
```
make container
```

Start image
```
docker run -d -e COUNTDOWN_DESTINATION_EVENT='Lauch web Day!' -e COUNTDOWN_DESTINATION_TIME='Sep 2, 2019 at 9:00am (EST)' -p 8888:8080 --name countdown01 countdown-timer
```
