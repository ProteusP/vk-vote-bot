box.cfg { listen = 3301,
    log_level = 5,
}

s = box.schema.space.create('votes', {
    if_not_exists = true
})

s:format({
    { name = 'id',       type = 'string' },
    { name = 'creator',  type = 'string' },
    { name = 'question', type = 'string' },
    { name = 'options',  type = 'map' },
    { name = 'votes',    type = 'map' },
    { name = 'status',   type = 'string' }
})

s:create_index('primary', {
    parts = { 'id' },
    if_not_exists = true
})


-- [DEBUG START]

-- local function format(data)
--     local format_order = { 'id', 'creator', 'question', 'options', 'votes', 'status' }
--     local tuple = {}
--     for _, field in ipairs(format_order) do
--         table.insert(tuple, data[field])
--     end
--     return tuple
-- end

-- local vote_data = {
--     id = '1',
--     creator = 'tester',
--     question = 'who?',
--     options = { ["me"] = 0, ["you"] = 0 },
--     votes = { ["user"] = 1 },
--     status = "active"
-- }
box.schema.user.grant('guest', 'read,write,execute', 'universe')
-- s:insert(format(vote_data))
-- [DEBUG END]
print("Таблица 'votes' успешно создана!")
