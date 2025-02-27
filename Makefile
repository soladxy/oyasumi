gen_orm:
	go run ./biz/gen/gorm

gen_service:
	go run ./biz/gen/service -name=user -outdir=./biz/service

# 更新项目
update:
	hz update -idl idl/oyasumi.thrift --handler_by_method --customize_package=hz_tmpl/package.yaml

# 新建项目
new_project:
	hz new -module github.com/soladxy/oyasumi -idl idl/oyasumi.thrift --handler_by_method --customize_layout=hz_tmpl/layout.yaml --customize_package=hz_tmpl/package.yaml

build:
	sh build.sh

wire:
	cd biz/container/ && wire && cd ../../

TAG := $(shell git rev-parse --short HEAD)-$(shell TZ=Asia/Shanghai date +%Y%m%d)

# 构建镜像
build_image:
	@echo "Building image: 971181317/oyasumi:$(TAG)"
	docker build -t 971181317/oyasumi:$(TAG) -t 971181317/oyasumi:latest .

image_push:
	docker push 971181317/oyasumi:$(TAG)
	docker push 971181317/oyasumi:latest

# 重新运行，之前已经已经运行过，会先删除再运行
re_run:
	docker stop oyasumi && \
	docker rm oyasumi && \
	docker run -dit -p 8890:8888 --name oyasumi --network dxytoll 971181317/oyasumi:latest

# 运行镜像
run:
	docker run -dit -p 8890:8888 --name oyasumi --network dxytoll 971181317/oyasumi:latest
