# 设置变量
SERVICES := user publish favorite message comment relation

.PHONY: all $(SERVICES)

all: $(SERVICES)

$(SERVICES):
	$(MAKE) -C cmd/$@

