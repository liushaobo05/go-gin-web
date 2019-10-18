
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

function randomLetter(len)
    local rt = ""
    for i = 1, len, 1 do
        rt = rt..string.char(random(97,122))
    end
    return rt
end

local str = "hello world" .. randomLetter(12)

body = {
    username =  randomLetter(12),
    name = randomLetter(12),
    password =  randomLetter(12),
    email = randomLetter(12),
    phone = randomLetter(12),
  }

for k,v in ipairs(body) do
        print(k)
end



