'use strict'

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    await queryInterface.createTable('tags', {
      id: {
        type: Sequelize.BIGINT,
        autoIncrement: true,
        allowNull: false,
        primaryKey: true
      },
      name: {
        type: Sequelize.STRING
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

    await queryInterface.createTable('tags-taggable', {
      tagId: {
        type: Sequelize.BIGINT,
        unique: 'tt_unique_constraint'
      },
      taggableId: {
        type: Sequelize.UUID,
        unique: 'tt_unique_constraint',
        references: null
      },
      taggableType: {
        type: Sequelize.STRING,
        unique: 'tt_unique_constraint'
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
    await queryInterface.dropTable('tags-taggable')
    await queryInterface.dropTable('tags')
  }
}
