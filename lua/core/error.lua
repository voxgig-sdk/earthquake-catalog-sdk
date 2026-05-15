-- EarthquakeCatalog SDK error

local EarthquakeCatalogError = {}
EarthquakeCatalogError.__index = EarthquakeCatalogError


function EarthquakeCatalogError.new(code, msg, ctx)
  local self = setmetatable({}, EarthquakeCatalogError)
  self.is_sdk_error = true
  self.sdk = "EarthquakeCatalog"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function EarthquakeCatalogError:error()
  return self.msg
end


function EarthquakeCatalogError:__tostring()
  return self.msg
end


return EarthquakeCatalogError
