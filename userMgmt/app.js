const express = require('express')
const cors = require('cors')
const bodyParser = require('body-parser')
const logger = require('morgan')
const passport  = require('passport')

const routes = require('./routes')

const app = express()

const PORT = process.env.PORT || 5000

require('./config/passport');

app.use(cors());
app.use(bodyParser.urlencoded({extended: false}))
app.use(bodyParser.json())
app.use(logger('dev'))
app.use(passport.initialize())

app.use('/', routes)
app.use('*', (req, res) => res.status(404).json({message: 'route does not exist.'}))

app.listen(PORT, () => console.log(`Listening on port ${PORT}`))


