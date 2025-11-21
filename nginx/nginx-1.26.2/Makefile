
default:	build

clean:
	rm -rf Makefile objs

.PHONY:	default clean

build:
	$(MAKE) -f objs/Makefile

install:
	$(MAKE) -f objs/Makefile install

modules:
	$(MAKE) -f objs/Makefile modules

upgrade:
	/home/azdhar/evergon/nginx/build/sbin/nginx -t

	kill -USR2 `cat /home/azdhar/evergon/nginx/build/logs/nginx.pid`
	sleep 1
	test -f /home/azdhar/evergon/nginx/build/logs/nginx.pid.oldbin

	kill -QUIT `cat /home/azdhar/evergon/nginx/build/logs/nginx.pid.oldbin`

.PHONY:	build install modules upgrade
