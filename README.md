# **Pavel-dis-tele-bot**
A bot that helps to transfer messages with some preconfigured prefix from Discord to the telegram channel

To use the bot you have to preconfigure it before start:

1. Create a channel in Telegram and get this channel's id.
2. You should create a Telegram bot and add it to channel.
3. Create a Discord bot and add it to your channel with access to read messages
4. Add the following values to config file _config.json_

Example:

`
{
    "telegram_token" : "5517450103:AAHaG7Emh7FtF69XkeT6va6bJ4AUrEKAwZY",
    "telegram_chat_id" : "-1001617785435",
    "discord_token" : "OTkzNDYzMzQ2NzE4MzB7ODI1.GKbXDh.4pmIxmJdiCmLehHCmjCJc_Y7YHKg7jLHJ_8b2w",
    "prefix" : "@Ждун"
}
`