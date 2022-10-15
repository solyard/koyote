<p align="center">
  <img src="coyote.png" />
</p>

<p align="center"><b>Koyote - Fast Gitlab Event Notification for Telegram</b></p>

<hr>
<p align="center"><img src="https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white"><img src="https://img.shields.io/badge/GitLab-330F63?style=for-the-badge&logo=gitlab&logoColor=white"><img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"></p>


# FYI
 🦊 Koyote (no it's not misspelling) it's a simple and fast telegram bot that can integrate with your Gitlab (Cloud or SelfHosted) to notify you about events that happen in your project

# TL;DR
Run Koyote as Binary
```
KOYOTE_API_PORT=8081 KOYOTE_TELEGRAM_BOT_TOKEN=abc:11223344 ./koyote 
```

Run Koyote with Docker:
```
docker run -p 8081:8081 koyote:v0.1 -e KOYOTE_API_PORT=8081 -e KOYOTE_TELEGRAM_BOT_TOKEN=abc:11223344
```

# How Koyote works

1. Koyote receive the event from Gitlab
2. Koyote trying to parse responce from API to known models
3. Koyote template message for telegram notification
4. Koyote send the message to telegram chat or channel depends on ID that received from WebHook URL


# Koyote params list
|Parameter|Description|Default Value|
|--|--|--|
|`KOYOTE_API_PORT`|PORT for Koyote to receive the events from Gitlab| 8081|
|`KOYOTE_TELEGRAM_BOT_TOKEN`| Bot token for telegram from @BotFather | empty (required)|
|`KOYOTE_ENABLE_JOB_NOTIFICATION`|Enable notification to telegram for JOB event |false|
|`KOYOTE_ENABLE_MR_NOTIFICATION`|Enable notification to telegram for MERGE REQUEST event |true|
|`KOYOTE_ENABLE_NOTE_NOTIFICATION`|Enable notification to telegram for NOTE event |false|
|`KOYOTE_ENABLE_PIPELINE_NOTIFICATION`|Enable notification to telegram for PIPELINE event |true|
|`KOYOTE_ENABLE_PUSH_NOTIFICATION`|Enable notification to telegram for PUSH event |false|
|`KOYOTE_ENABLE_TAG_PUSH_NOTIFICATION`|Enable notification to telegram for TAG PUSH event |false|
|`KOYOTE_REDIS_ENABLED`|Enable Redis for Event pooling if Telegram cannot be reached/whatever|false|
|`KOYOTE_REDIS_CHECK_UNSENDED_EVENTS_INTEVAL`|Interval (in seconds) to check for unsended event to Telegram|empty|
|`KOYOTE_REDIS_UNSENDED_TASK_TTL`|Time in seconds to expiration for unsended event to Telegram|empty|
|`KOYOTE_REDIS_INSTANCE_URI`|Redis instance host|empty|
|`KOYOTE_REDIS_INSTANCE_PORT`|Redis instance port|empty|
|`KOYOTE_REDIS_USERNAME`|Username for Redis (if not default 'root')|empty|
|`KOYOTE_REDIS_PASSWORD`|Password for Redis|empty|
# Gitlab configuration

1. Open your project in Gitlab and go to Settings -> Webhooks
2. Check the triggers that you need to receive in Telegram
3. Insert URL for example: http://koyote/notify/<chat_id>
4. Press the button Add Webhook at the bottom of the page
5. At the bottom of the page, you can see a new webhook. Try to send a test event with the button "Test" and select event that you want to receive


# Telegram configuration
1. Go to the @BotFather
2. To create the Bot follow the instructions
3. Get Bot API TOKEN and forward it to Koyote with ENV variable KOYOTE_TELEGRAM_BOT_TOKEN
4. Enjoy :)

# ROADMAP
- [ ] Improve stability and fix the codestyle
- [ ] Implement a normal logic for taskpooler
- [ ] ...


# For what this bot?

We faced the problem from developers that spends a lot of time notifying the team about new MR requests or notify about failed builds and pipelines in the development chat. 
So I created this bot to make notifications more quickly and automated. This bot makes helping our developers not spend time notifying others manually and resending asks about new MR's to other developers. Thats all. 

While testing this bot in my projects for a month and not founded a normal realization of bot like mine i decided to create another community version for developers that looking for something to automate notifying process.

<p align="center"><a href="https://www.flaticon.com/ru/free-icons/" title="волк иконки">Logo from Freepik - Flaticon</a></p>

