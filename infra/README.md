# Deployment

## Caddy

### Install

```bash
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https curl
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy
```

### Configuration

In /etc/caddy/ add the following configuration in a file named Caddyfile:

```
<API_URL> {

	reverse_proxy localhost:8080

}
```

This sets up a reverse proxy for a specific API URL to the service running on localhost:8080.

```bash
sudo systemctl start caddy
```

## Systemd

### Configuration

```bash
sudo ln -s /yab-explorer/infra/yab-explorer-api-sepolia.service /etc/systemd/system/yab-explorer-api-sepolia.service
sudo systemctl enable yab-explorer-api-sepolia.service
sudo systemctl start yab-explorer-api-sepolia
```
