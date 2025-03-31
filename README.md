# Nut Pushover

Uses Pushover to send notifications for Network Upstools Monitor
[![Go](https://github.com/ComputerComa/NutPushover/actions/workflows/go.yml/badge.svg)](https://github.com/ComputerComa/NutPushover/actions/workflows/go.yml)

## upsnotify



Called by UPSmon, in turn calls the Pushover.Net API.

See `upsmon.conf` at `/etc/nut/upsmon.conf`
Place this binary in the path defined as `NOTIFYCMD`:

```
NOTIFYCMD /usr/local/bin/upsnotify
```

UPSmon calls the script with the environment variables `NOTIFYTYPE` and `UPSNAME`. The only arg provided is the event message.

## config

Provided by environment variables or `/etc/default/upsnotifyPushover`


## Required
```
PUSHOVER_API_KEY=A000000000000000000000000000000
PUSHOVER_USER_KEY=u00000000000000000000000000000
```

