# Redmine Telegram Notification Bot

## Project Overview
Redmine Telegram Notification is a Telegram bot designed to notify users about updates and status changes of issues in Redmine. Developed for personal use, this project can be utilized by others, and feedback is greatly appreciated. <br>
For any inquiries or suggestions, reach out via Telegram: [@danek_kulikov](https://t.me/danek_kulikov)

## Table of Contents
- [Configuration](#configuration)
- [Functionality](#functionality)


## Configuration
Create a `.env` file in the root directory of the project with the following environment variables:
- `REDMINE_BASE_URL`: Base URL for your Redmine instance like https://redmine.your-company.com
- `REDMINE_API_KEY`: Your API key for Redmine.
- `TELEGRAM_BOT_TOKEN`: The token for your Telegram bot.
- `CHAT_ID`: The chat ID where the bot will send notifications.
- `PROJECTS_LIST`: Here is a list of project IDs for which you want to receive notifications. You can get the list of project IDs and their names using the GetProjectsList() method in the redmine package. This method is called by default in main and logs the list of projects to the console.
- `GOOGLE_DEV_API_KEY`: Create an API app in the [google developer account](https://console.developers.google.com/), from the "Credentials" tab you can create an API key, you get something like this AIzaSyBcOT_DpEQysiwFmmmZXupKpnrOdJYAhhM

**ATTENTION, CUSTOM RUSSIAN HOLIDAY CALENDAR IN USE!** <br>
If you are not from Russia, you can find an API for your country's calendar. I don't think there would be any difficulties with that.

Example `.env` file:
```dotenv
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
CHAT_ID=your_chat_id
REDMINE_API_KEY=your_redmine_api_key
BASE_URL=https://redmine.your-company.com
PROJECTS_LIST=1,2,3
GOOGLE_DEV_API_KEY=your_google_dev_api_key
```

## Functionality

The bot retrieves a list of tasks through the Redmine API, maps the tasks, and adds comments to each task. <br>
After a specified time interval, it fetches a new list of tasks with the same request, maps the new list of tasks, and adds comments to each new task. <br>
Then, the old map and the new map are compared by iterating through each task in a for loop. <br>
If a new task appears in the new map that was not in the old map, a notification about the new task is immediately created. <br>
If the task is not new but a change of interest is found, this change is added to the message about task updates. <br>
Once all relevant changes for the task are added, a template notification message is generated and sent to the specified Telegram chat. <br>

### Notifications
Generates a notification message that will be sent to a Telegram chat:
- **New Issues:** Notifies about newly added issues.
- **Status change:** Аdds a message indicating that the task status has changed.
- **Priority change:** Аdds a message indicating that the task priority has changed.
- **Assignee change:** Аdds a message indicating that the task assignee has changed.
- **New comment:** Аdds a message indicating that a comment has been added, including the comment text.
- **Task emoji tagging:** tags tasks based on status and tracker.


### Time Management
- **Work Time Check:** Verifies if the current time falls within working hours.
- **Holidays:** Integrates with Google Calendar to check for holidays.

Feel free to contribute, fork, or suggest improvements. Your feedback is valuable in making this bot more efficient and user-friendly. For further assistance or to report issues, please contact via Telegram: [@danek_kulikov](https://t.me/danek_kulikov)

For a full explanation of the code and more detailed instructions, please refer to each file's comments and functions within the codebase.
