box.cfg { listen = 3301 }

local votes = box.schema.space.create('votes',{
    format = {
        {name = 'id', type = 'string'},
        {name = 'crator', type = 'string'},
        {name = 'question', type = 'string'},
        {name = 'options', type = 'map'},
        {name = 'votes', type = 'map'},
        {name = 'status', type = 'string'}
    }
})

box.schema.user.create('user', {password = 'password'})
box.schema.user.grant('user', 'read,write', 'space','votes')
