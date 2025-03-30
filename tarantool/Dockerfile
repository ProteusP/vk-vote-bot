FROM tarantool/tarantool:2.11.5

COPY db_init.lua /opt/tarantool/

ENV TARANTOOL_USER_NAME=user \
    TARANTOOL_USER_PASSWORD=password \
    TARANTOOL_SCRIPT=/opt/tarantool/db_init.lua

EXPOSE 3301

CMD ["tarantool", "/opt/tarantool/db_init.lua"]
