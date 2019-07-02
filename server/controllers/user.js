'use strict'
let jwt = require('jsonwebtoken');
let mockUser = 'user';
let mockPass = 'pass';

module.exports = {
  signup: (req, res) => {
    return res.json("Signed up!");
  },
  login: (req, res) => {
    let username = req.body.username;
    let password = req.body.password;
    if (username == mockUser && password == mockPass) {
      // get the users id from db
      let token = jwt.sign({id: 1}, process.env.JWTSECRET);
      return res.json({
        sucess: true,
        token: token
      });
    }
    return res.status(400).json({
      success: false,
      message: 'Incorrect username and password'
    });
  },
  update: (req, res) => {
    return res.json('update user: ' + req.params.id);
  },
  getUser: (req, res) => {
    return res.json('get user: ' + req.params.id);
  }
}
