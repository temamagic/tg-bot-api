
<a name="v0.6.9"></a>
## [v0.6.9](https://github.com/temamagic/tgbot/compare/v0.6.8...v0.6.9)

> 0001-01-01


<a name="v0.6.8"></a>
## v0.6.8

> 2023-04-12

### Add

* new sticker types properties 

### Build

* add changelog generator 
* fix module 

### Chore

* go mod bullshit 

### Chores

* remove or moved some files 

### Docs

* fix url 
* changed readme 

### Feat

* add set and get for bot description and short description 
* threads construct helper 
* custom helper for some my purposes in real life project 
* api 6.4 changes from OvyFlash/telegram-bot-api 
* faster json lib 
* add IsTextMention to MessageEntity 
* add premium animation into message struct 
* add is premium user status 
* Add CustomTitle field to ChatMember 
* add sendDice configs [Add sendDice config to use in Send method
as a Chattable interface.
Add NewDice and NewDiceWithEmoji helpers
Add tests

https://core.telegram.org/bots/api#senddice]
* Replaces *http.Client with an interface [Signed-off-by: Andrii Soluk <isoluchok[@gmail](https://github.com/gmail).com>]
* **bot:** implement unban chat member for channels 
* **bot:** implement deleteMessage method 
* **types:** add language code to user struct 

### Fix

* UserShared & ChatShared json 
* missed sentfrom & fromchat actions 
* omitempty for optional fields 

### README

* link to godoc 
* fix typo 

### Reverts

* Remove a broken check for GetUpdatesChan.

