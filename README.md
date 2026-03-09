# Free Deployment Guide

Your app serves both frontend and backend from a single Go server. The frontend automatically fetches configuration from `/api/config`, so you only need to update the backend to change values.

## Option 1: Railway (Recommended - Easiest)

**Free tier:** 500 hours/month, $5 credit

1. Create account at [railway.app](https://railway.app)
2. Install Railway CLI:
   ```bash
   npm i -g @railway/cli
   # OR
   brew install railway
   ```
3. Login and deploy:
   ```bash
   railway login
   railway init
   railway up
   ```
4. Get your URL from the dashboard

**To update configuration:**

- Edit [constants.go](constants.go)
- Run `railway up` to redeploy
- Frontend automatically gets new values!

## Option 2: Render

**Free tier:** Available with some limitations

1. Create account at [render.com](https://render.com)
2. Click "New +" → "Web Service"
3. Connect your GitHub repo
4. Render auto-detects the [render.yaml](render.yaml) config
5. Click "Create Web Service"

**To update configuration:**

- Edit [constants.go](constants.go)
- Push to GitHub
- Render auto-deploys
- Frontend automatically gets new values!

## Option 3: Fly.io

**Free tier:** 3 shared VMs

1. Install Fly CLI:
   ```bash
   brew install flyctl
   ```
2. Login and launch:
   ```bash
   fly auth login
   fly launch
   ```
3. Deploy:
   ```bash
   fly deploy
   ```

**To update configuration:**

- Edit [constants.go](constants.go)
- Run `fly deploy`
- Frontend automatically gets new values!

## Local Testing Before Deploy

```bash
# Test with Docker
docker build -t terminal-beautified .
docker run -p 8080:8080 terminal-beautified

# Open http://localhost:8080
```

## How It Works

1. **Backend** serves static files (HTML, CSS) and provides `/api/config` endpoint
2. **Frontend** fetches config on page load from `/api/config`
3. **Update process:**
   - Change values in [constants.go](constants.go)
   - Redeploy backend
   - Frontend automatically fetches new values ✨
   - No frontend rebuild needed!

## Environment Variables

Set `PORT` environment variable if needed (default: 8080):

- Railway: Settings → Variables
- Render: Environment tab
- Fly.io: `fly secrets set PORT=8080`
