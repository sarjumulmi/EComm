/* jshint indent: 2 */

module.exports = function(sequelize, DataTypes) {
  return sequelize.define('User', {
    userId: {
      type: DataTypes.INTEGER(11),
      autoIncrement: true,
      allowNull: false,
      primaryKey: true
    },
    firstName: {
      type: DataTypes.STRING(45),
      allowNull: false
    },
    lastName: {
      type: DataTypes.STRING(45),
      allowNull: false
    },
    email: {
      type: DataTypes.STRING(100),
      allowNull: false,
      unique: true
    },
    address: {
      type: DataTypes.STRING(100),
      allowNull: true
    },
    passwordHash: {
      type: DataTypes.STRING(100),
      allowNull: false
    },
    userType: {
      type: DataTypes.STRING(45),
      allowNull: false
    }
  }, {
    tableName: 'User',
    timestamps: false
  });
};
