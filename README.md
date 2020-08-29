# AnomalyBot
This is a disord bot written in Golang that is used in 3 different servers 2 of which are the official development servers. The bot provides some utility like playing music into a voice chat, things like moderation to kick and ban people, integration with the Project Anomaly API to give you stats and more
## Part of the Anomaly Software Package™
As part of the Anomaly Software Package it comes with integration for Project Anomaly, Project Apex, the Project Anomaly itch.io page and the Anomaly website.
## Installation
1. Install Golang
2. Download the bot files 
3. Create a discord application and create a bot in it
4. Open the terminal, cd into the folder with the bot files and do `go run . -t <your bot token here>` or do `go build .` and `./bot -t <your token here>`
5. OPTIONAL: If you are running this on a server you may want to keep the bot running even after closing the ssh connection for this you can use `screen` then skip through the notices and run you application with `./bot -t <your token here>`
6. You are ready
