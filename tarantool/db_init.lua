box.cfg {
    listen = 3301,
    log_level = 5,
    memtx_dir = '/var/lib/tarantool/data'
}

local is_initialized = box.space._schema:get('INIT_DONE') ~= nil

if not is_initialized then
    box.once('init_votes_space', function()
        local s = box.schema.space.create('votes', {
            if_not_exists = true,
            format = {
                { name = 'id',       type = 'string' },
                { name = 'creator',  type = 'string' },
                { name = 'question', type = 'string' },
                { name = 'options',  type = 'map' },
                { name = 'votes',    type = 'map' },
                { name = 'status',   type = 'string' }
            }
        })

        s:create_index('primary', {
            parts = { 'id' },
            if_not_exists = true
        })
    end)

    box.once('grant_guest_access', function()
        box.schema.user.grant('guest', 'read,write,execute', 'universe')
    end)

    box.space._schema:replace { 'INIT_DONE', true }

    print("Инициализация БД выполнена!")
else
    print("БД уже инициализирована, пропускаем setup")
end
