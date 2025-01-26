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

pkg_container:
	docker build -t sola_api_go .

re_run:
	docker stop sola_api_go && \
	docker rm sola_api_go && \
	docker run -dit -p 8890:8888 --name sola_api_go --network dxytoll sola_api_go

run:
	docker run -dit -p 8890:8888 --name sola_api_go --network dxytoll sola_api_go
