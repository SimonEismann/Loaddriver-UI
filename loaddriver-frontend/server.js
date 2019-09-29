const express = require('express');
const path = require('path');
const app = express();

app.use(express.static(path.join(__dirname, 'public')));

const WINDOW_ENV =
{
  API_ROOT: process.env.API_ROOT,
  CONSOLE_URI: process.env.CONSOLE_URI,
}

app.get('/env.js', function (req, res) {
  res.setHeader('Cache-Control', 'public, max-age=300');
  res.send("window.env = " + JSON.stringify(WINDOW_ENV));
});

app.get('/*', function (req, res) {
  res.sendFile(path.join(__dirname, 'public', 'index.html'));
});

app.listen(process.env.PORT || 9000);
