# основные пути
GOPATH			:=	${shell pwd}
BINPATH			=	$(GOPATH)/bin

# основные команды
GOCMD			=	go

# параметры команд
GODEP			=	$(GOCMD) get
GOBUILD			=	$(GOCMD) build

GOTEST			=	$(GOCMD) test
GOINSTALL		=	$(GOCMD) install

# установка пути
export GOPATH

# настройка пакетов для сборки
MAIN_PKGS 		:=	topcloud.ru/api

LIBS_PKGS		:=

DEPS_PKGS 		:=	git.topcloud.ru/topcloud.ru/api/pkg \
					git.topcloud.ru/topcloud.ru/api/pkg/config \
					git.topcloud.ru/topcloud.ru/api/pkg/mongo \
					git.topcloud.ru/topcloud.ru/api/pkg/server \
					git.topcloud.ru/topcloud.ru/api/pkg/mock \
					github.com/gorilla/handlers \
					github.com/gorilla/mux \
					github.com/dgrijalva/jwt-go \
					github.com/BurntSushi/toml \
					github.com/google/uuid \
					golang.org/x/crypto/bcrypt \
					gopkg.in/mgo.v2 \
					gopkg.in/mgo.v2/bson

TEST_PKGS		:=	$(LIBS_PKGS) $(MAIN_PKGS)

# списки для сборки
DEPS_LIST		=	$(foreach int, $(DEPS_PKGS), $(int)_deps)
BUILD_LIST		=	$(foreach int, $(MAIN_PKGS), $(int)_build)

TEST_LIST		=	$(foreach int, $(TEST_PKGS), $(int)_test)
INSTALL_LIST	=	$(foreach int, $(MAIN_PKGS), $(int)_install)

# пишем все в .PHONY
.PHONY:			$(DEPS_LIST) $(TEST_LIST) $(BUILD_LIST) $(INSTALL_LIST)

all:			build
deps:			$(DEPS_LIST)
test:			$(TEST_LIST)
build:			$(BUILD_LIST)
install:		$(INSTALL_LIST)

$(DEPS_LIST): %_deps:
	$(GODEP) $*
$(BUILD_LIST): %_build:
	$(GOBUILD) -o $(BINPATH)/$(shell basename $*) $*

$(TEST_LIST): %_test:
	$(GOTEST) $*
$(INSTALL_LIST): %_install:
	$(GOINSTALL) $*
