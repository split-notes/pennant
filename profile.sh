#!/bin/bash


# Use: view/edit documentation
pennant () {
    # if no command given force help page
    local OPTION
	if [[ "$1" != "" ]]; then
        OPTION=$1
    else
        OPTION="help"
    fi
	# handle input options
    case "${OPTION}" in
        'help')
echo "Usage: $ ${FUNCNAME} [option]

Options:
- help: show this menu
- start: launch pennant's docker container(s)
- start-quiet: launch pennant's docker container(s), no output
- stop: stop pennant's docker container(s)"
        ;;
        "start")
            docker-compose up
            echo "STARTING pennant"
        ;;
        "start-quiet")
            docker-compose up -d
            echo "STARTING pennant in background"
        ;;
        'stop')
            echo "STOPPING pennant_app"
#            docker-compose down
            docker stop pennant_app
            docker rm $(docker ps -a -q)
            docker rmi pennant_app
        ;;
        'restart')
            pennant-down
            pennant-up
        ;;
        *)
            echo -e "ERROR: invalid option. Try..\n$ ${FUNCNAME} help"
        ;;
    esac
}
