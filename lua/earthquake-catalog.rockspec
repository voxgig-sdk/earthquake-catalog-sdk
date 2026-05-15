package = "voxgig-sdk-earthquake-catalog"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/earthquake-catalog-sdk.git"
}
description = {
  summary = "EarthquakeCatalog SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["earthquake-catalog_sdk"] = "earthquake-catalog_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
