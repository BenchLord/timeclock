'use strict'
let jwt = require('jsonwebtoken');

module.exports = {
  checkToken: (req, res, next) => {
    let token = req.headers['x-access-token'] || req.headers['authorization'];
    if (token && token.startsWith('Bearer ')) {
      token = token.slice(7, token.length);
    }
    if (token) {
      jwt.verify(token, process.env.JWTSECRET, (err, token) => {
        if (err) {
          return res.json({
            error: err.name + ": " + err.message,
            success: false
          })
        } else {
          req.token = token;
          next();
        }
      });
    } else {
      return res.json({
        error: "No jwt in header",
        success: false
      });
    }
  }
}
