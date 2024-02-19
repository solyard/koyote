<p align="center">
  <img src="koyote.jpg" width="400" height="400"/>
</p>

<p align="center"><b>Koyote - Fast GitLab Event Notifications for Telegram</b></p>

<hr>
<p align="center"><img src="https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white"><img src="https://img.shields.io/badge/GitLab-330F63?style=for-the-badge&logo=gitlab&logoColor=white"><img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"></p>


# FYI
🦊 Koyote (no, it's not a misspelling) is a simple and fast Telegram bot that integrates with your GitLab (Cloud or Self-Hosted) to notify you about events that happen in your project.

# TL;DR
Run Koyote as Binary
```
KOYOTE_API_PORT=8081 KOYOTE_TELEGRAM_BOT_TOKEN=abc:11223344 ./koyote 
```

Run Koyote with Docker:
```
docker run -p 8081:8081 koyote:v0.1 -e KOYOTE_API_PORT=8081 -e KOYOTE_TELEGRAM_BOT_TOKEN=abc:11223344
```

# Chat Utils

Bot now supports some commands:
```
/chatID - Return current chatID
/threadID - Return current threadID (If 0 then you are in General Thread or chat without Threads support)
```

# How Koyote works

1. Koyote receives an event from GitLab
2. Koyote tries to parse the response from the API to known models
3. Koyote templates a message for a Telegram notification
4. Koyote sends the message to a Telegram chat or channel depending on the ID that is received from the WebHook URL


# Koyote parameters
|Parameter|Description|Default Value|
|--|--|--|
|`KOYOTE_API_PORT`|Koyote web-server port| 8081|
|`KOYOTE_TELEGRAM_BOT_TOKEN`|Telegram bot token from @BotFather| empty (required)|
|`KOYOTE_ENABLE_JOB_NOTIFICATION`|Enable Telegram notification for JOB event|false|
|`KOYOTE_ENABLE_MR_NOTIFICATION`|Enable Telegram notification for MERGE REQUEST event|true|
|`KOYOTE_ENABLE_NOTE_NOTIFICATION`|Enable Telegram notification for NOTE event|false|
|`KOYOTE_ENABLE_PIPELINE_NOTIFICATION`|Enable Telegram notification for PIPELINE event|true|
|`KOYOTE_ENABLE_PUSH_NOTIFICATION`|Enable Telegram notification for PUSH event|false|
|`KOYOTE_ENABLE_TAG_PUSH_NOTIFICATION`|Enable Telegram notification for TAG PUSH event|false|

# GitLab configuration

1. Open your project in GitLab and go to Settings -> Webhooks
2. Check the triggers that you want to receive in Telegram
3. Insert the URL (e.g. `http://koyote/notify/<chat_id>/<topic_id>`) (topic_id optional field if you are using Telegram Topics)
4. Press the "Add Webhook" button at the bottom of the page
5. At the bottom of the page, you can see the new webhook. Try sending a test event with the "Test" button and select the event you want to receive

# Telegram configuration
1. Go to the @BotFather
2. Follow the instructions to create the bot
3. Get the Bot API token and forward it to Koyote with the environment variable `KOYOTE_TELEGRAM_BOT_TOKEN`
4. Enjoy :)

# Roadmap
- [ ] Improve stability and fix the codestyle
- [ ] Implement a normal logic for taskpooler
- [ ] ...

# What is this bot for?

We faced the problem of developers spending a lot of time notifying the team about new MR requests or failed builds and pipelines in the development chat. So I created this bot to make notifications faster and more automated. This bot helps our developers save time by not having to notify others manually and resend requests for new MRs to other developers. That's all.

After testing this bot in my projects for a month and not finding a normal realization of a bot like mine, I decided to create a community version for developers who are looking for something to automate the notification process.

<p align="center">Logo by MJ 5.2 :)</p>

