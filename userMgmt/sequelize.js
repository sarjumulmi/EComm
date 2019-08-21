const Sequelize = require('sequelize');
const UserModel = require('./models/User');

const sequelize = new Sequelize('EComm', 'root', 'Oracle123', {
  host: 'localhost',
  dialect: 'mysql',
  pool: {
    max: 10,
    min: 0,
    acquire: 30000,
    idle: 10000
  }
})

const User = UserModel(sequelize, Sequelize)

const email = 's@m.com'
sequelize.sync({
    force: false
  })
  .then(() => {
    console.log(`Database & tables created!`)
    // User.findOne({where: {email}}).then(user => console.log(user));
  })

module.exports = {
  User
}