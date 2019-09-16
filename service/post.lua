-- example HTTP POST script which demonstrates setting the
-- HTTP method, body, and adding a header
-- local socket = require('socket')


-- function Random(length)
--     math.randomseed(socket.gettime()*10000)
-- 	local res = ""
-- 	for i = 1, length do
-- 		res = res .. string.char(math.random(97, 122))
-- 	end
-- 	return res
-- end


wrk.method = "POST"
-- wrk.body   = "{\"login\":\"" .. Random(math.random(5, 9)) .. "\", \"password\": \"" .. Random(math.random(5, 9)) .. "\"}"
wrk.body   = "{\"login\":\"Kola\", \"pass\": \"123456789\", \"age\": 12}"
wrk.headers["Content-Type"] = "application/json"

done = function(summary, latency, requests)
	io.write("------------------------------\n")
	for _, p in pairs({ 50, 90, 99, 99.999 }) do
	   n = latency:percentile(p)
	   io.write(string.format("%g%%,%d\n", p, n))
	end
 end