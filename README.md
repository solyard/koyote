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


<p align="center"><a href="https://www.flaticon.com/ru/free-icons/" title="волк иконки">Logo from Freepik - Flaticon</a></p>

