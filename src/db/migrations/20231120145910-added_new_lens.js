'use strict'

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    await queryInterface.createTable('lenses', {
      id: {
        type: Sequelize.UUID,
        defaultValue: Sequelize.UUIDV4,
        allowNull: false,
        primaryKey: true
      },
      name: {
        type: Sequelize.STRING
      },
      description: {
        type: Sequelize.TEXT
      },
      version: {
        type: Sequelize.INTEGER
      },
      spec: {
        type: Sequelize.JSONB
      },
      isDraft: {
        type: Sequelize.BOOLEAN,
        defaultValue: true
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

    await queryInterface.createTable('lenses-pillars-questions', {
      id: {
        type: Sequelize.BIGINT,
        autoIncrement: true,
        allowNull: false,
        primaryKey: true
      },
      ref: {
        type: Sequelize.STRING,
        allowNull: false
      },
      name: {
        type: Sequelize.STRING
      },
      description: {
        type: Sequelize.TEXT
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

    await queryInterface.createTable('lenses-pillars-questions-resources', {
      id: {
        type: Sequelize.BIGINT,
        autoIncrement: true,
        allowNull: false,
        primaryKey: true
      },
      questionId: {
        type: Sequelize.BIGINT,
        references: {
          model: 'lenses-pillars-questions',
          key: 'id'
        },
        onUpdate: 'CASCADE',
        onDelete: 'CASCADE'
      },
      url: {
        type: Sequelize.STRING
      },
      description: {
        type: Sequelize.TEXT
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

    await queryInterface.createTable('lenses-pillars-risks', {
      id: {
        type: Sequelize.BIGINT,
        autoIncrement: true,
        allowNull: false,
        primaryKey: true
      },
      risk: {
        type: Sequelize.STRING,
        allowNull: false
      },
      condition: {
        type: Sequelize.STRING,
        allowNull: false
      },
      description: {
        type: Sequelize.TEXT
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

    await queryInterface.createTable('lenses-pillars-choices', {
      id: {
        type: Sequelize.BIGINT,
        autoIncrement: true,
        allowNull: false,
        primaryKey: true
      },
      ref: {
        type: Sequelize.STRING,
        allowNull: false
      },
      noneOfThese: {
        type: Sequelize.BOOLEAN,
        allowNull: true,
        defaultValue: false
      },
      name: {
        type: Sequelize.STRING
      },
      description: {
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

    await queryInterface.createTable('lenses-pillars', {
      id: {
        type: Sequelize.BIGINT,
        autoIncrement: true,
        allowNull: false,
        primaryKey: true
      },
      ref: {
        type: Sequelize.STRING,
        allowNull: false
      },
      name: {
        type: Sequelize.STRING
      },
      description: {
        type: Sequelize.TEXT
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

    await queryInterface.createTable('lenses-pillars-resources', {
      id: {
        type: Sequelize.BIGINT,
        autoIncrement: true,
        allowNull: false,
        primaryKey: true
      },
      pillarId: {
        type: Sequelize.BIGINT,
        references: {
          model: 'lenses-pillars',
          key: 'id'
        },
        onUpdate: 'CASCADE',
        onDelete: 'CASCADE'
      },
      url: {
        type: Sequelize.STRING
      },
      description: {
        type: Sequelize.TEXT
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

    await queryInterface.createTable('lenses-pillars-choices-resources', {
      id: {
        type: Sequelize.BIGINT,
        autoIncrement: true,
        allowNull: false,
        primaryKey: true
      },
      choiceId: {
        type: Sequelize.BIGINT,
        references: {
          model: 'lenses-pillars-choices',
          key: 'id'
        },
        onUpdate: 'CASCADE',
        onDelete: 'CASCADE'
      },
      url: {
        type: Sequelize.STRING
      },
      description: {
        type: Sequelize.TEXT
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

    await queryInterface.addColumn('lenses-pillars-choices', 'questionId', {
      type: Sequelize.BIGINT,
      references: {
        model: 'lenses-pillars-questions',
        key: 'id'
      },
      allowNull: true,
      onDelete: 'CASCADE'
    })

    await queryInterface.addColumn('lenses-pillars-risks', 'questionId', {
      type: Sequelize.BIGINT,
      references: {
        model: 'lenses-pillars-questions',
        key: 'id'
      },
      allowNull: true,
      onDelete: 'CASCADE',
      onUpdate: 'CASCADE'
    })

    await queryInterface.addColumn('lenses-pillars-questions', 'pillarId', {
      type: Sequelize.BIGINT,
      references: {
        model: 'lenses-pillars',
        key: 'id'
      },
      allowNull: true,
      onDelete: 'CASCADE',
      onUpdate: 'CASCADE'
    })

    await queryInterface.addColumn('lenses-pillars', 'lensId', {
      type: Sequelize.UUID,
      references: {
        model: 'lenses',
        key: 'id'
      },
      allowNull: true,
      onDelete: 'CASCADE',
      onUpdate: 'CASCADE'
    })
  },

  async down(queryInterface, Sequelize) {
    await queryInterface.dropTable('lenses-pillars-resources')
    await queryInterface.dropTable('lenses-pillars-questions-resources')
    await queryInterface.dropTable('lenses-pillars-choices-resources')
    await queryInterface.dropTable('lenses-pillars-risks', { cascade: true })
    await queryInterface.dropTable('lenses-pillars-choices', { cascade: true })
    await queryInterface.dropTable('lenses-pillars-questions', {
      cascade: true
    })
    await queryInterface.dropTable('lenses-pillars', { cascade: true })
    await queryInterface.dropTable('lenses', { cascade: true })
  }
}
