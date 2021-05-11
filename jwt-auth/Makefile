SHELL += -eu

BLUE	:= \033[0;34m
GREEN	:= \033[0;32m
RED   	:= \033[0;31m
NC    	:= \033[0m
REGION	:= eu-west-1
STAGE	:= prod
VERSION	:=$$(jq -rM '.version' package.json)
# AWS_ACCESS_KEY_ID    :=$$(./get-aws-profile.sh --profile=$(profile) --key)
# AWS_SECRET_ACCESS_KEY:=$$(./get-aws-profile.sh --profile=$(profile) --secret)
REGIONS="eu-west-1 \
	eu-west-1"

build:
	@echo "${GREEN}✓ building lambda functions... ${NC}\n"
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin/monitor ./src/monitor
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin/export-index ./src/export-index
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin/export-data ./src/export-data
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin/native-export-index ./src/native-export-index
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin/native-snapshot ./src/native-snapshot
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin/trigger-native-snapshot ./src/trigger-native-snapshot

init:
	go mod init

test:
	go test ./src/native-snapshot/cleanup
	@echo "${GREEN}✓ Starting local ES cluster...${NC}\n"
	ID=$$(docker ps | grep -F "elasticsearch" | awk '{ print $$1 }') && \
	if test "$$ID" = ""; then docker run -d --name elasticsearch --net magic -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.1.1; fi  && \
	echo "${RED}✓ Wait for 10s for ES cluster...${NC}\n"
	sleep 10s
	@echo "${GREEN}✓ Testing lambda functions...${NC}\n"
	go test --count=1 -v ./src/export-index
	go test --count=1 -v ./src/export-data
	@echo "${RED}✓ Stopping and removing local ES cluster...${NC}\n"
	docker stop elasticsearch && docker rm elasticsearch
	@echo "${GREEN}✓ Install dependancies...${NC}\n"
	npm install

strip:
	strip ./bin/*

clean:
	rm -r ./bin/

deploy: clean build strip
	@echo "${GREEN}✓ Update version ${NC}\n"
	npm version patch
	@echo "${GREEN}✓ Deploying lambda function and associated resources ${NC}\n"
	npx sls deploy --stage ${STAGE} --region ${REGION} --aws-profile $(profile) --verbose
	@echo "${GREEN}✓ setting the Dockerfile version... ${NC}\n"
	sed -i '/version=/s/".*"/\"'${VERSION}'"/' Dockerfile.blaze
	sed -i '/version=/s/".*"/\"'${VERSION}'"/' Dockerfile.native
	@echo "${GREEN}✓ building docker image for ECS Fargate task... ${NC}\n"
	docker build -t export_data -f Dockerfile.blaze .
	docker build -t native_snapshot -f Dockerfile.native .
	@echo "${GREEN}✓ tag image so that we can push the image to the repository... ${NC}\n"
	# @echo "${GREEN}✓ push this image... ${NC}\n"
	# docker push `aws sts get-caller-identity --output text --query 'Account'`.dkr.ecr.${REGION}.amazonaws.com/export_data:${VERSION}-${STAGE}
	# docker push `aws sts get-caller-identity --output text --query 'Account'`.dkr.ecr.${REGION}.amazonaws.com/native_snapshot:${VERSION}-${STAGE}
	@echo "${GREEN}✓ remove dangling images... ${NC}\n"
	# docker rmi $(docker images --quiet --filter "dangling=true")

logs:
	npx sls logs --stage ${STAGE} --region ${REGION} --aws-profile $(profile) -f exportIndex

push:
	docker tag export_data:latest `aws sts get-caller-identity --output text --query 'Account'`.dkr.ecr.${REGION}.amazonaws.com/export_data:${VERSION}-${STAGE}
	docker tag native_snapshot:latest `aws sts get-caller-identity --output text --query 'Account'`.dkr.ecr.${REGION}.amazonaws.com/native_snapshot:${VERSION}-${STAGE}
	@echo "${GREEN}✓ Login to ECR... ${NC}\n"
	$$(aws ecr get-login --no-include-email --region ${REGION} --profile=$(profile))
	@echo "${GREEN}✓ push this image... ${NC}\n"
	docker push `aws sts get-caller-identity --output text --query 'Account'`.dkr.ecr.${REGION}.amazonaws.com/export_data:${VERSION}-${STAGE}
	docker push `aws sts get-caller-identity --output text --query 'Account'`.dkr.ecr.${REGION}.amazonaws.com/native_snapshot:${VERSION}-${STAGE}

# local: build
#	docker run --rm -it -v $HOME/.aws:/root/.aws -e PROFILE=magic -e LOCAL=1 -e ELASTICSEARCH_ADDRESS=https://search-mw-elk-vhbrgdr5nvml7wggcspvam2i7y.eu-west-1.es.amazonaws.com/ -e INDEXES_TABLE=mw-es-indexes-eu-west-1-prod -e STATUS_GSI=GSI-eu-west-1-prod-status-index -e EXPORTS_BUCKET=mw-mgm-es-backup-prod-exportsbucket-1x3bx5z6e9jn8 -e ES_AWS_ACCESS_KEY=${AWS_ACCESS_KEY_ID} ES_AWS_SECRET_KEY=${AWS_SECRET_ACCESS_KEY} exort_data:latest
remove:
	npx sls remove --stage prod --region eu-west-1 --aws-profile $(profile) --verbose
