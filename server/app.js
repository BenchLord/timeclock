'use strict'
let express = require('express');
let bodyParser = require('body-parser');
let app = express();
if (process.env.NODE_ENV != 'production') {
  require('dotenv').config();
}

app.use(bodyParser.urlencoded({
  extended: true
}));

app.use('/user', require('./routes/user.js'));

let mw = require('./routes/middleware.js');
app.get('/test', mw.checkToken, (req, res) => {
  res.json(req.token);
})

let port = process.env.PORT;
app.listen(port, () => {
  console.log("Running on port " + port);
})
