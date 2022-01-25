# fireboy

## Setup

Add the bot discord token to `fireboy/.env`:

```
echo 'DISCORD_TOKEN=**********' > .env
```

Set trigger word for your test bot in `fireboy/.env` (production default is ;;fb):

```
# Example using ';;fbat' as the trigger string
echo 'TRIGGER_STRING=;;fbat' >> .env
```

Run the bot (builds first and then runs):
```
go run fireboy
```
