<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Redirecting...</title>

  <!-- GA4 -->
  <script async src="https://www.googletagmanager.com/gtag/js?id=G-KPZSVGQCBF"></script>
  <script>
    window.dataLayer = window.dataLayer || [];
    function gtag(){ dataLayer.push(arguments); }
    gtag('js', new Date());
    gtag('config', 'G-KPZSVGQCBF');
  </script>

  <style>
    body {
      margin: 0;
      display: flex;
      align-items: center;
      justify-content: center;
      min-height: 100vh;
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
      background: #f5f5f5;
      color: #333;
    }
    .msg { text-align: center; }
    .msg p { font-size: 15px; color: #666; margin-top: 8px; }
  </style>
</head>
<body>
<div class="msg">
  <p>Redirecting to Google Play...</p>
</div>

<script>
  // App configs
  const APPS = {
    bedrock: { id: 'com.mcmodshub.app', medium: 'short' },
    java:    { id: 'com.mcmodshub.jav', medium: 'short' },
    pixwall: { id: 'com.pixwall.app',   medium: 'short' }
  };

  // Source → medium mapping
  const MEDIUM = {
    tiktok:  'short',
    youtube: 'short',
    reddit:  'post'
  };

  // Read params: ?app=bedrock&src=tiktok
  const params = new URLSearchParams(window.location.search);
  const appKey = (params.get('app') || '').toLowerCase();
  const src    = (params.get('src') || 'unknown').toLowerCase();
  const app    = APPS[appKey];

  if (!app) {
    // Unknown app — redirect to main GitHub page
    window.location.href = 'https://lashkulwork-ux.github.io';
  } else {
    const medium   = MEDIUM[src] || 'link';
    const campaign = 'organic_2026';

    const playUrl = `https://play.google.com/store/apps/details`
      + `?id=${app.id}`
      + `&utm_source=${src}`
      + `&utm_medium=${medium}`
      + `&utm_campaign=${campaign}`;

    // Fire GA4 event BEFORE redirect
    gtag('event', 'click_to_store', {
      app_name:     appKey,
      traffic_source: src,
      medium:       medium,
      campaign:     campaign
    });

    // Small delay so GA4 event has time to send
    setTimeout(() => {
      window.location.href = playUrl;
    }, 300);
  }
</script>
</body>
</html>
