FROM tarantool/tarantool:2.11.5

COPY db_init.lua /opt/tarantool/
RUN mkdir -p /var/lib/tarantool/data
EXPOSE 3301

# Запускаем Tarantool с нашим скриптом
CMD ["tarantool", "/opt/tarantool/db_init.lua"]
