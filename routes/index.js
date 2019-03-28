'use strict';

var express = require('express');
var router = express.Router();
var userSchema = require('../models/user');
var bankSchema = require('../models/bankinfo');
var qresult = {};
var loginID ='';

// block chain const from here      
const fs = require('fs');
var reguser = require('../network/registerUser');
var invoketran = require('../network/invoke');
var queryuser = require('../network/query');
// block chain const end here

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('login', { title: 'Profile' });
});

router.post('/login', async function(req, res, next) {
  if (typeof req.body.loginid === 'undefined') 
      req.body.loginid=[];
  else {
    loginID = req.body.loginid;
    var querytype = 'user';
    var qresult = await queryuser.qresult(loginID, querytype);
    res.render('userinfo', {title: 'User Registration',bag:qresult});
  }
});

router.post('/viewbank', async function(req, res, next) {
  if (typeof req.body.loginid === 'undefined') 
      req.body.loginid=[];
  else {
    loginID = req.body.loginid;
    var querytype = 'bank';
    var qresult = await queryuser.qresult(loginID, querytype);
    res.render('bankinfo', {title: 'Bank Information',bag:qresult});
  }
});

//adding a new user
router.post('/register', function(req, res, next) {
  res.render('register', {title: 'User Registration'});
  if (typeof req.body.loginid === 'undefined') 
      req.body.loginid=[];
  else {
    var user = new userSchema (
    { loginid: req.body.loginid,
      password: req.body.password,
      first_name: req.body.first_name,
      last_name: req.body.last_name,
      address1: req.body.address1,
      address2: req.body.address2,
      city: req.body.city,
      state: req.body.state,
      zipcode: req.body.zipcode,
      country: req.body.country,
      user_type: req.body.user_type,
      email: req.body.email,
      phone: req.body.phone
      });
      user.save(function (err) {
        if (err) console.log(err);
      });
      
      // The blockchain code from here
      console.log('calling register user');
      loginID = req.body.loginid;
      let response = reguser(loginID, user);
    };
});

//adding user's bank account info
router.post('/addbank', function(req, res, next) {
  res.render('bankadd', {title: 'User Bank Information'});
  if (typeof req.body.loginid === 'undefined') 
      req.body.loginid=[];
  else {
    var bank = new bankSchema (
    { objectype: 'mypro_bank',
      loginid: req.body.loginid,
      bank_profile_name: req.body.Profilename,
      account_type: req.body.Accountype,
      bank_name: req.body.Bankname,
      routing_num: req.body.Routing,
      account_num: req.body.Account,
      verification_result: req.body.Validated,
      shared_with_login: req.body.Sharedlogin
      });
      bank.save(function (err) {
        if (err) console.log(err);
      });
      
      // The blockchain code from here
      console.log('calling add bank');
      loginID = req.body.loginid;
      let response = invoketran(loginID, bank);
    };
});
  /*        
  // invoke the main function, can catch any error that might escape
  main().then(() => {
    console.log('done');
  }).catch((e) => {
      console.log('Final error checking.......');
      console.log(e);
      console.log(e.stack);
      process.exit(-1);
  }); */
  //

module.exports = router;