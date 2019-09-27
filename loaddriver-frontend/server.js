const express = require('express');
const path = require('path');
const app = express();

app.use(express.static(path.join(__dirname, 'build')));

const WINDOW_ENV =
  "window.env={'API_ROOT':'" +
  process.env.API_ROOT +
  "', 'CLIENT_ID':'" +
  process.env.CLIENT_ID +
  "', 'REDIRECT':'" +
  process.env.REDIRECT +
  "'}";

app.get('/env.js', function(req, res) {
  res.setHeader('Cache-Control', 'public, max-age=300');
  res.send(WINDOW_ENV);
});

app.get('/*', function(req, res) {
  res.sendFile(path.join(__dirname, 'build', 'index.html'));
});

app.listen(process.env.PORT || 9000);
