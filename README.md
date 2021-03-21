# Twitch Chat Archiver

A tool to archive an entire VOD chat. The goal of this tool is to save all relavent data for a complete offline replay of a VOD chat.

### Features
- Archives include embedded badges and emote images / gifs
- Supported JSON and Protocol Buffer output files.

### Supported Image / GIF Archives
- Twitch Robot Emotes
- Twitch Global Emotes
- Channel Subscription Emotes
- BetterTTV Global Emotes
- BetterTTV Channel Emotes
- FrankerFaceZ Global Emotes
- FrankerFaceZ Channel Emotes
- Twitch Global Badges
- Channel Subscription Badges
- Channel Bit Badges
- Channel Cheer Badges

## Installation / Usage
You will need a Twitch Client ID. Register an application [here](https://dev.twitch.tv/console/apps/create) and copy the Client ID. The OAuth Redirect URL can be any value.


1. Download the [latest release](https://github.com/harry93848bb7/twitch-chat-archive/releases) for your system
2. Extract the files in a folder of your choice
2. Run the CLI: `./chat-archiver -client_id=4ab086186863bffbf81a73d359fc2ec7 -vod_id=951894217`

### CLI Arguments
```
  -client_id string
        Twitch Developer Application Client ID
  -exclude value
        What to exclude from the archival. Possible values "messages", "badges" or "emotes"
  -output_dir string
        Where to output archive files (default ".")
  -remove_unused
        Remove dangling emotes / badges which are not used throughout the entire VOD (default true)
  -vod_id string
        Target VOD ID to archive
```

### Unsupported
- Twitch Turbo variations of the Robot emotes are not archived.
- Channel Subscription Emotes from channels outside of the VOD are not archived.

### TODO:
- Investigate support for archival of subscription emotes from other channels