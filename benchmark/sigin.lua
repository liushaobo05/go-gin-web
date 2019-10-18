wrk.method = "POST"
wrk.path = "/api/v1/signIn"
wrk.body   = [[
  {
    "username": "liushaobo",
    "password": "123456"
  }
]]
wrk.headers["Content-Type"] = "application/json"