# 更新项目
update:
	hz update -idl idl/oyasumi.thrift

# 新建项目
new_project:
	hz new -module soladxy/oyasumi -idl idl/oyasumi.thrift

build:
	sh build.sh