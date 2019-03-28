var express = require('express');
var router = express.Router();

/* GET users listing. */
router.get('/', function(req, res, next) {
  res.send('respond with a resource');
});

/* GET users cool listing. */
router.get('/cool', function(req, res, next) {
  res.send('You are so Cool');
});

/* GET users login 
router.get('/login', function(req, res, next) {
  res.render('login', { title: 'Login Screen' });
});
*/

module.exports = router;
