# Hydra

## Create a OAuth2 Client

```bash
client=$(hydra create client \
    --endpoint http://hydra:4445/ \
    --format json \
    --grant-type client_credentials)

client_id=$(echo $client | jq -r '.client_id')
client_secret=$(echo $client | jq -r '.client_secret')
```

## Perform the client credentials flow

```bash
hydra perform client-credentials \
  --endpoint http://hydra:4444/ \
  --client-id "$client_id" \
  --client-secret "$client_secret"
```

## Perform token introspection

```bash
hydra introspect token \
  --format json-pretty \
  --endpoint http://hydra:4445/ \
  ${your_token}
```

## Create a OAuth2 Client for Authorization Code Flow

```bash
code_client=$(hydra create client \
    --endpoint http://hydra:4445 \
    --grant-type authorization_code,refresh_token \
    --response-type code,id_token \
    --format json \
    --scope openid --scope offline \
    --redirect-uri http://127.0.0.1:5173/callback)

code_client_id=$(echo $code_client | jq -r '.client_id')
code_client_secret=$(echo $code_client | jq -r '.client_secret')
```

## Add 'hydra' to `hosts` file

open /etc/hosts and add `hydra` at the end of the following line:

```
127.0.0.1 localhost hydra
```

## Start an example web app

```bash
hydra perform authorization-code \
    --client-id $code_client_id \
    --client-secret $code_client_secret \
    --endpoint http://hydra:4444/ \
    --port 5173 \
    --scope openid --scope offline
```