'use strict'

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    await queryInterface.createTable(
      'ownership',
      {
        id: {
          type: Sequelize.BIGINT,
          autoIncrement: true,
          allowNull: false,
          primaryKey: true
        },
        resourceId: {
          type: Sequelize.UUID
        },
        ownerId: {
          type: Sequelize.UUID
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
      },
      {
        uniqueKeys: {
          Items_unique: {
            fields: ['resourceId', 'ownerId']
          }
        }
      }
    )
  },

  async down(queryInterface, Sequelize) {
    await queryInterface.dropTable('ownership')
  }
}
