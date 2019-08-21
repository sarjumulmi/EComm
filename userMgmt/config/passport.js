const bcrypt = require('bcrypt')
const passport = require('passport')
const localStrategy = require('passport-local').Strategy
const JWTStrategy = require('passport-jwt').Strategy
const ExtractJwt = require('passport-jwt').ExtractJwt

const {
  User
} = require('../sequelize')

const JWT_SECRET = 'myjwtsecret'

const BCRYPT_SALT = 12

passport.use(
  'register',
  new localStrategy({
      usernameField: 'email',
      passwordField: 'password',
      session: false,
      passReqToCallback: true
    },
    (req, email, password, done) => {
      console.log("Before finding user");
      try {
        User.findOne({
          where: {
            email: email
          }
        }).then(user => {
          if (user !== null) {
            console.log('user email already exists')
            return done(null, false, {
              message: 'User with email already exists'
            })
          } else {
            bcrypt.hash(password, BCRYPT_SALT).then(passwordHash => {
              User.create({
                email,
                passwordHash,
                firstName: req.body.firstName,
                lastName: req.body.lastName,
                address: req.body.email,
                userType: req.body.userType
              }).then(user => {
                console.log('User created..')
                return done(null, user)
              })

            })
          }
        })
      } catch (error) {
        console.log(error);
        done(error)
      }
    }
  )
)

passport.use(
  'login',
  new localStrategy({
      usernameField: 'email',
      passwordField: 'password',
      session: false
    },
    (email, password, done) => {
      try {
        User.findOne({
          where: {
            email
          }}).then(user => {
            if (user === null) {
              return done(null, false, {
                message: 'Incorrect Username'
              })
            } else {
              bcrypt.compare(password, user.passwordHash).then(res => {
                if (res !== true) {
                  console.log('Email/Password mismatch')
                  return done(null, false, {
                    message: 'Email/Password do not match'
                  })
                }
                console.log('user authenticated')
                return done(null, user)
              })
            }
          })
      } catch (error) {
        done(error)
      }
    })
)

const opts = {
  jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
  secretOrKey: JWT_SECRET
}

passport.use(
  'jwt',
  new JWTStrategy(opts, (jwtPayload, done) => {
    try {
      User.findOne({
        where: {
          email: jwtPayload.email
        }
      }).then(user => {
        if (user) {
          console.log('user found')
          done(null, user)
        } else {
          console.log('user not found')
          done(null, false)
        }
      })
    } catch (error) {
      done(error)
    }
  })
)


module.exports = {
  JWT_SECRET
}