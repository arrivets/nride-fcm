# FCM Service

FCM Service is an HTTP service to send FCM remote notifications to nride users.
It holds a registry of [NKN address] to [FCM Token], and exposes methods to 
create, delete and query entries in this registry. It also exposes an endpoint
to post notifications to users by NKN address.

## Build and Install

Build directly on the hosting server:

```
ssh -i ~/.ssh/awsfaucet ubuntu@faucet.nride.com
cd go/src/github.com/arrivets/nride-fcm
make local
sudo systemctl restart fcm.service
```

## Config

The app reads a GOOGLE_APPLICATION_CREDENTIALS environment variable that should
point to the secret file containing the credentials.

Use a `.env` file for this in the root repo.

## Use

The faucet runs as a `systemctl` service

### logs:

`journalctl -u fcm.service -f`
