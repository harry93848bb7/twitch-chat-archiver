# Twitch Chat Archiver

A tool to archive an entire VOD chat. The goal of this tool is provide all relavent data for a complete offline replay of a VOD chat.

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

### Unsupported
- Twitch Turbo variations of the Robot emotes are not archived.
- Channel Subscription Emotes from channels outside of the VOD are not archived.

### TODO:
- Improve logging when archiving message chunks (progress logging)
- Investigate support for archival of subscription emotes from other channels