function random(n, m)
    math.randomseed(os.clock()*math.random(1000000,90000000)+math.random(1000000,90000000))
    return math.random(n, m)
end

function randomNumber(len)
    local rt = ""
    for i=1,len,1 do
        if i == 1 then
            rt = rt..random(1,9)
        else
            rt = rt..random(0,9)
        end
    end
    return rt
end

function randomString(len)
    local rt = ""
    for i = 1, len, 1 do
        rt = rt..string.char(random(97,122))
    end
    return rt
end

wrk.method = "POST"
wrk.path = "/api/v1/signUp"
wrk.body = [[
  {
	"username": "test123",
    "name": "test123",
    "password": "test123",
    "email": "test123",
    "phone": "test"
  }
]]
wrk.headers["Content-Type"] = "application/json"
