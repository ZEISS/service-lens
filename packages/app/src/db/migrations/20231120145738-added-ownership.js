'use strict'

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    await queryInterface.createTable('ownership', {
      resourceId: {
        type: Sequelize.UUID,
        unique: 'ownership_unique_constraint'
      },
      resourceType: {
        type: Sequelize.STRING,
        unique: 'ownership_unique_constraint'
      },
      ownerId: {
        type: Sequelize.UUID,
        unique: 'ownership_unique_constraint'
      },
      createdAt: {
        type: Sequelize.DATE
      },
      updatedAt: {
        type: Sequelize.DATE
      },
      deletedAt: {
        type: Sequelize.DATE
      }
    })
  },

  async down(queryInterface, Sequelize) {
    await queryInterface.dropTable('ownership')
  }
}
