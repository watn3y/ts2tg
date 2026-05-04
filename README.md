# ts2tg

ts2tg notifies you on telegram when users join/leave your teamspeak server

> [!IMPORTANT]  
> The version of this repository you are seeing is likely a mirror.
>
> If you want to contribute any changes please head over to [git.watn3y.de](https://git.watn3y.de/watn3y/ts2tg) and register an account.

## Running with Docker Compose

Docker image: <https://hub.docker.com/r/watn3y/ts2tg>

Example compose file:

```yaml
services:
  ts2tg:
    image: watn3y/ts2tg:latest
    container_name: ts2tg
    restart: unless-stopped
    volumes:
      - /etc/localtime:/etc/localtime:ro
    environment:
      #TS2TG_LOGLEVEL:
      #TS2TG_SLEEPINTERVAL:
      TS2TG_TELEGRAM_APITOKEN:
      TS2TG_TELEGRAM_CHATID:
      TS2TG_TEAMSPEAK_BASEURL:
      TS2TG_TEAMSPEAK_APIKEY:
      #TS2TG_TEAMSPEAK_JOINURL:
```

Available tags: latest, v1.0.0,v1.1.0,etc

## Running on Linux

Grab a release from the [releases page](https://git.watn3y.de/watn3y/ts2tg/releases). Make sure to set your [environment variables](#environment-variables) accordingly.

## HTML redirection

A example HTML file to redirect your Users directly to their TeamSpeak application [is provided here](html/)

## Environment Variables

> [!NOTE]  
> For development purposes, ts2tg supports loading environment variables from a .env file placed in the projects root directory.

| Variable                  | Description                                                                                                                                  | Default  | Required | Example                                       |
| ------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------- | -------- | --------------------------------------------- |
| `TS2TG_LOGLEVEL`          | LogLevel as described [in the zerolog documentation](https://pkg.go.dev/github.com/rs/zerolog@v1.34.0#readme-simple-leveled-logging-example) | 1 (Info) | ❌       | 1                                             |
| `TS2TG_SLEEPINTERVAL`     | Amount of seconds to wait between requests to TeamSpeak                                                                                      | 60       | ❌       | 60                                            |
| `TS2TG_TELEGRAM_APITOKEN` | Telegram BotToken, get it from [@BotFather on Telegram](https://t.me/BotFather)                                                              | None     | ✅       | 1234567890:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw |
| `TS2TG_TELEGRAM_CHATID`   | Chat to notify about users joining/leaving TeamSpeak                                                                                         | None     | ✅       | -1001234567890                                |
| `TS2TG_TEAMSPEAK_BASEURL` | Base URL of your TeamSpeak Server. Including the Virtual Server. For most people this will be /1                                             | None     | ✅       | http://example.com/1                          |
| `TS2TG_TEAMSPEAK_APIKEY`  | TeamSpeak WebQuery API Key which can be generated [like this](https://community.teamspeak.com/t/what-is-apikey-for/6156/2)                   | None     | ✅       | A7B3C9D2E5F1A4B8C6D9E2F5A8B1C4D7E0F3A6B9      |
| `TS2TG_TEAMSPEAK_JOINURL` | HTTP(S) URL of your redirect.html                                                                                                            | None     | ❌       | https://example.com/redirect.html             |

## Nice to know

### Semantic Versioning

This project does it's best to follow [Semantic Versioning](https://semver.org/#semantic-versioning-200), however I can't guarantee anything.

### AI Disclosure

AI was used in the process of creating this project. Mainly for the HTML redirection file because I am shockingly bad at anything that isn't backend.
