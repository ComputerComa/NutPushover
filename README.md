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

## Optional
PUSHOVER_Sound=SOUNDNAME (!!! MUST MATCH A PUSHOVER BUILTIN SOUND OR A CUSTOM UPLOADED SOUND !!!)
PUSHOVER_Priority=0


### Using Priorities

From the Pushover.net [[API Docs](https://pushover.net/api#priority)]
```

Lowest Priority (-2)

When the priority parameter is specified with a value of -2, messages will be considered lowest priority and will not generate any notification. On iOS, the application badge number will be increased.

Low Priority (-1)

Messages with a priority parameter of -1 will be considered low priority and will not generate any sound or vibration, but will still generate a popup/scrolling notification depending on the client operating system. Messages delivered during a user's quiet hours are sent as though they had a priority of (-1).

Normal Priority (0)

Messages sent without a priority parameter, or sent with the parameter set to 0, will have the default priority. These messages trigger sound, vibration, and display an alert according to the user's device settings. On iOS, the message will display at the top of the screen or as a modal dialog, as well as in the notification center. On Android, the message will scroll at the top of the screen and appear in the notification center.

If a user has quiet hours set and your message is received during those times, your message will be delivered as though it had a priority of -1.

High Priority (1)

Messages sent with a priority of 1 are high priority messages that bypass a user's quiet hours. These messages will always play a sound and vibrate (if the user's device is configured to) regardless of the delivery time. High-priority should only be used when necessary and appropriate.

High-priority messages are highlighted in red in the device clients.

Emergency Priority (2)

Emergency-priority notifications are similar to high-priority notifications, but they are repeated until the notification is acknowledged by the user. These are designed for dispatching and on-call situations where it is critical that a notification be repeatedly shown to the user (or all users of the group that the message was sent to) until it is acknowledged. The first user in a group to acknowledge a message will cancel retries for all other users in the group. 

```