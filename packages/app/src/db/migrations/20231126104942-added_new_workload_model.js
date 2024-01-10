'use strict'

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    await queryInterface.createTable('workloads', {
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
      profilesId: {
        type: Sequelize.UUID,
        references: {
          model: 'profiles',
          key: 'id'
        },
        allowNull: true,
        onDelete: 'CASCADE'
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

    await queryInterface.createTable('workloads-lenses', {
      id: {
        type: Sequelize.BIGINT,
        autoIncrement: true,
        allowNull: false,
        primaryKey: true
      },
      workloadId: {
        type: Sequelize.UUID,
        references: {
          model: 'workloads',
          key: 'id'
        },
        onUpdate: 'CASCADE',
        onDelete: 'CASCADE'
      },
      lensId: {
        type: Sequelize.UUID,
        references: {
          model: 'lenses',
          key: 'id'
        },
        onUpdate: 'CASCADE',
        onDelete: 'CASCADE'
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

    await queryInterface.createTable(
      'workloads-lenses-answers',
      {
        id: {
          type: Sequelize.BIGINT,
          autoIncrement: true,
          allowNull: false,
          primaryKey: true
        },
        workloadId: {
          type: Sequelize.UUID,
          references: {
            model: 'workloads',
            key: 'id'
          },
          onUpdate: 'CASCADE',
          onDelete: 'CASCADE'
        },
        lensPillarQuestionId: {
          type: Sequelize.BIGINT,
          references: {
            model: 'lenses-pillars-questions',
            key: 'id'
          },
          onUpdate: 'CASCADE',
          onDelete: 'CASCADE'
        },
        notes: {
          type: Sequelize.TEXT
        },
        doesNotApply: {
          type: Sequelize.BOOLEAN,
          defaultValue: false
        },
        doesNotApplyReason: {
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
      },
      {
        uniqueKeys: {
          Items_unique: {
            fields: ['workloadId', 'lensPillarQuestionId']
          }
        }
      }
    )

    await queryInterface.createTable(
      'workloads-lenses-answers-choices',
      {
        id: {
          type: Sequelize.BIGINT,
          autoIncrement: true,
          allowNull: false,
          primaryKey: true
        },
        answerId: {
          type: Sequelize.BIGINT,
          references: {
            model: 'workloads-lenses-answers',
            key: 'id'
          },
          onUpdate: 'CASCADE',
          onDelete: 'CASCADE'
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
            fields: ['answerId', 'choiceId']
          }
        }
      }
    )

    await queryInterface.createTable('workloads-environment', {
      id: {
        type: Sequelize.BIGINT,
        autoIncrement: true,
        allowNull: false,
        primaryKey: true
      },
      workloadId: {
        type: Sequelize.UUID
      },
      environmentId: {
        type: Sequelize.BIGINT,
        references: {
          model: 'environments',
          key: 'id'
        },
        onUpdate: 'CASCADE',
        onDelete: 'CASCADE'
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
    await queryInterface.dropTable('workloads-lenses-answers-choices')
    await queryInterface.dropTable('workloads-lenses-answers')
    await queryInterface.dropTable('workloads-environment')
    await queryInterface.dropTable('workloads-lenses', { cascade: true })
    await queryInterface.dropTable('workloads')
  }
}
