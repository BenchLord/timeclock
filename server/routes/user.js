'use strict'
let ctl = require('../controllers/user.js');
let mw = require('./middleware.js');
let express = require('express');
let router = express.Router();

router.get('/:id', ctl.getUser);
router.put('/:id', ctl.update);
router.post('/signup', ctl.signup);
router.post('/login', ctl.login);

module.exports = router;
