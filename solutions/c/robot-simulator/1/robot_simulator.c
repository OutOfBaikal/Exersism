#include "robot_simulator.h"
#include <string.h>

robot_status_t robot_create(robot_direction_t direction, int x, int y) {
    robot_position_t new_position = { x, y };
    robot_status_t robot = { direction, new_position };
    return robot;
}
void robot_move(robot_status_t *robot, const char *commands) {
    size_t len = strlen(commands);
    for (size_t i = 0; i < len; ++i) {
        switch(commands[i]) {
            case 'R':
                switch(robot->direction) {
                    case DIRECTION_NORTH:
                        robot->direction = DIRECTION_EAST;
                        break;
                    case DIRECTION_EAST:
                        robot->direction = DIRECTION_SOUTH;
                        break;
                    case DIRECTION_SOUTH:
                        robot->direction = DIRECTION_WEST;
                        break;
                    case DIRECTION_WEST:
                        robot->direction = DIRECTION_NORTH;
                        break;
                    default:
                        break;
                }
                break;
            case 'L':
                switch(robot->direction) {
                    case DIRECTION_NORTH:
                        robot->direction = DIRECTION_WEST;
                        break;
                    case DIRECTION_EAST:
                        robot->direction = DIRECTION_NORTH;
                        break;
                    case DIRECTION_SOUTH:
                        robot->direction = DIRECTION_EAST;
                        break;
                    case DIRECTION_WEST:
                        robot->direction = DIRECTION_SOUTH;
                        break;
                    default:
                        break;
                }
                break;
            case 'A':
                switch(robot->direction) {
                    case DIRECTION_NORTH:
                        ++robot->position.y;
                        break;
                    case DIRECTION_EAST:
                        ++robot->position.x;
                        break;
                    case DIRECTION_SOUTH:
                        --robot->position.y;
                        break;
                    case DIRECTION_WEST:
                        --robot->position.x;
                        break;
                    default:
                        break;
                }
                break;
            default:
                break;
        }
    }
}
