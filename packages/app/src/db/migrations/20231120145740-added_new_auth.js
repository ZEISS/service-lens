'use strict'

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    await queryInterface.createTable('verification_token', {
      token: {
        type: Sequelize.STRING,
        primaryKey: true
      },
      identifier: {
        type: Sequelize.STRING,
        allowNull: false
      },
      expires: {
        type: Sequelize.DATE,
        allowNull: false
      }
    })

    await queryInterface.createTable('accounts', {
      id: {
        type: Sequelize.UUID,
        defaultValue: Sequelize.UUIDV4,
        primaryKey: true
      },
      type: {
        type: Sequelize.STRING,
        allowNull: false
      },
      provider: {
        type: Sequelize.STRING,
        allowNull: false
      },
      provider_account_id: {
        type: Sequelize.STRING,
        allowNull: false
      },
      refresh_token: {
        type: Sequelize.STRING
      },
      access_token: {
        type: Sequelize.STRING
      },
      expires_at: {
        type: Sequelize.INTEGER
      },

      token_type: {
        type: Sequelize.STRING
      },
      scope: {
        type: Sequelize.STRING
      },
      id_token: {
        type: Sequelize.TEXT
      },
      session_state: {
        type: Sequelize.STRING
      },
      user_id: {
        type: Sequelize.UUID
      }
    })

    await queryInterface.createTable('sessions', {
      id: {
        type: Sequelize.UUID,
        defaultValue: Sequelize.UUIDV4,
        primaryKey: true
      },
      expires: {
        type: Sequelize.DATE,
        allowNull: false
      },
      session_token: {
        type: Sequelize.STRING,
        unique: 'session_token',
        allowNull: false
      },
      user_id: {
        type: Sequelize.UUID
      }
    })

    await queryInterface.createTable('users', {
      id: {
        type: Sequelize.UUID,
        defaultValue: Sequelize.UUIDV4,
        primaryKey: true
      },
      name: {
        type: Sequelize.STRING
      },
      email: {
        type: Sequelize.STRING,
        unique: 'email'
      },
      email_verified: {
        type: Sequelize.DATE
      },
      image: {
        type: Sequelize.STRING
      }
    })
  },

  async down(queryInterface, Sequelize) {
    await queryInterface.dropTable('verification_token')
    await queryInterface.dropTable('accounts')
    await queryInterface.dropTable('sessions')
    await queryInterface.dropTable('users', { cascade: true })
  }
}
