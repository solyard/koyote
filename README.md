<p align="center">
  <img src="coyote.png" />
</p>

<p align="center"><b>Koyote - Fast Gitlab Event Notification for Telegram</b></p>

<hr>
<p align="center"><img src="https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white"><img src="https://img.shields.io/badge/GitLab-330F63?style=for-the-badge&logo=gitlab&logoColor=white"><img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"></p>

<p align="center"><a href="https://www.flaticon.com/ru/free-icons/" title="волк иконки">Logo from Freepik - Flaticon</a></p>


# About
 Koyote (no it's not misspelling) it's a simple and fast telegram bot that can integrate with your Gitlab (Cloud or SelfHosted) to notify you about events that happen in your project

# TL;DR
Run Koyote as Binary
```
KOYOTE_API_PORT=8081 ./koyote 
```

Run Koyote with Docker:
```
export KOYOTE_API_PORT=8081
docker run -p $KOYOTE_API_PORT:KOYOTE_API_PORT koyote:v0.1 -e KOYOTE_API_PORT=$KOYOTE_API_PORT
```

# How Koyote works

# Gitlab configuration

# Telegram configuration