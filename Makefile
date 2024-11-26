# 定义变量
SERVER_DIR := yg_server
UI_DIR     := yg_desktop_ui

# 执行 make go 时的行为
go:
	(cd $(SERVER_DIR); fresh)

# 执行 make dev 时的行为
dev:
	(cd $(UI_DIR); pnpm run dev)

# 执行 make tidy 时的行为
tidy:
	(cd $(SERVER_DIR); go mod tidy)

# 执行 make install 时的行为
install:
	(cd $(UI_DIR); pnpm install)