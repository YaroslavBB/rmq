export ROOT=..
export DOCKER_PATH=$ROOT/docker
export REST_SERVICE=$ROOT/rest_service/cmd

cd $DOCKER_PATH && docker compose up -d
cd $REST_SERVICE && go run .