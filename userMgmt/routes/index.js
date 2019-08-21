const express = require('express')
const passport = require('passport')
const jwt = require('jsonwebtoken')
const {User} = require('../sequelize')

const {JWT_SECRET} = require('../config/passport');

const router = express.Router()

router.post('/registerUser', (req, res, next) => {
  passport.authenticate('register', (err, user, info) => {
    if (err) {
      console.log(err)
      return next(err)
    }
    if (info !== undefined) {
      console.log(info.message);
      return res.send(info.message)
    } else {
      const token  = jwt.sign({sub: 'user', user}, JWT_SECRET)
      console.log(user)
      return res.status(201).send({auth: true, token, message: 'user created'})   
    }
  })(req, res, next)
})

router.post('/login', (req, res, next) => {
  passport.authenticate('login', (err, user, info) => {
    if (err) {
      console.log(err)
      retures.send(err)
    }
    if (info !== undefined) {
      console.log(info.message);
      res.send(info.message)
    } else {
      const token = jwt.sign({sub: 'user', user}, JWT_SECRET)
      console.log(user);
      return res.status(200).json({auth: true, token, message: 'user authenticated'})
    }
  })(req, res, next)
})

module.exports = router