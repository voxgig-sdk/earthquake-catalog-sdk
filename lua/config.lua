-- EarthquakeCatalog SDK configuration

-- Build a fresh, fully materialised config table. Every call rebuilds the
-- whole structure, so prefer require("config_shared") unless you need a
-- private copy you intend to mutate.
local function make_config()
  return {
    main = {
      name = "EarthquakeCatalog",
      slug = "earthquake-catalog",
      version = "0.0.1",
      target = "lua",
    },
    feature = {
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
      },
    },
    options = {
      base = "https://earthquake.usgs.gov/fdsnws/event/1",
      headers = {
        ["content-type"] = "application/json",
      },
      entity = {
        ["earthquake_data"] = {},
        ["service_information"] = {},
      },
    },
    entity = {
      ["earthquake_data"] = {
        ["fields"] = {
          {
            ["name"] = "count",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "geometry",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "maxAllowed",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "properties",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "type",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "earthquake_data",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "alertlevel",
                      ["orig"] = "alertlevel",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "callback",
                      ["orig"] = "callback",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "catalog",
                      ["orig"] = "catalog",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "contributor",
                      ["orig"] = "contributor",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "2014-01-02",
                      ["kind"] = "query",
                      ["name"] = "endtime",
                      ["orig"] = "endtime",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "eventid",
                      ["orig"] = "eventid",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "earthquake",
                      ["kind"] = "query",
                      ["name"] = "eventtype",
                      ["orig"] = "eventtype",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "quakeml",
                      ["kind"] = "query",
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "includeallmagnitude",
                      ["orig"] = "includeallmagnitude",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "includeallorigin",
                      ["orig"] = "includeallorigin",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "includearrival",
                      ["orig"] = "includearrival",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = "false",
                      ["kind"] = "query",
                      ["name"] = "includedeleted",
                      ["orig"] = "includedeleted",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "includesuperseded",
                      ["orig"] = "includesuperseded",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "jsonerror",
                      ["orig"] = "jsonerror",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "kmlanimated",
                      ["orig"] = "kmlanimated",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = "age",
                      ["kind"] = "query",
                      ["name"] = "kmlcolorby",
                      ["orig"] = "kmlcolorby",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "latitude",
                      ["orig"] = "latitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "limit",
                      ["orig"] = "limit",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "longitude",
                      ["orig"] = "longitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "magnitudetype",
                      ["orig"] = "magnitudetype",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "maxcdi",
                      ["orig"] = "maxcdi",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 1000,
                      ["kind"] = "query",
                      ["name"] = "maxdepth",
                      ["orig"] = "maxdepth",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "maxgap",
                      ["orig"] = "maxgap",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 90,
                      ["kind"] = "query",
                      ["name"] = "maxlatitude",
                      ["orig"] = "maxlatitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 180,
                      ["kind"] = "query",
                      ["name"] = "maxlongitude",
                      ["orig"] = "maxlongitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "maxmagnitude",
                      ["orig"] = "maxmagnitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "maxmmi",
                      ["orig"] = "maxmmi",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 180,
                      ["kind"] = "query",
                      ["name"] = "maxradius",
                      ["orig"] = "maxradius",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 20001.6,
                      ["kind"] = "query",
                      ["name"] = "maxradiuskm",
                      ["orig"] = "maxradiuskm",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "maxsig",
                      ["orig"] = "maxsig",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "mincdi",
                      ["orig"] = "mincdi",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = -100,
                      ["kind"] = "query",
                      ["name"] = "mindepth",
                      ["orig"] = "mindepth",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "minfelt",
                      ["orig"] = "minfelt",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "mingap",
                      ["orig"] = "mingap",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = -90,
                      ["kind"] = "query",
                      ["name"] = "minlatitude",
                      ["orig"] = "minlatitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = -180,
                      ["kind"] = "query",
                      ["name"] = "minlongitude",
                      ["orig"] = "minlongitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 5,
                      ["kind"] = "query",
                      ["name"] = "minmagnitude",
                      ["orig"] = "minmagnitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "minsig",
                      ["orig"] = "minsig",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 204,
                      ["kind"] = "query",
                      ["name"] = "nodata",
                      ["orig"] = "nodata",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "offset",
                      ["orig"] = "offset",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = "time",
                      ["kind"] = "query",
                      ["name"] = "orderby",
                      ["orig"] = "orderby",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "productcode",
                      ["orig"] = "productcode",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "moment-tensor",
                      ["kind"] = "query",
                      ["name"] = "producttype",
                      ["orig"] = "producttype",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "all",
                      ["kind"] = "query",
                      ["name"] = "reviewstatus",
                      ["orig"] = "reviewstatus",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "2014-01-01",
                      ["kind"] = "query",
                      ["name"] = "starttime",
                      ["orig"] = "starttime",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "updatedafter",
                      ["orig"] = "updatedafter",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/query",
                ["parts"] = {
                  "query",
                },
                ["select"] = {
                  ["exist"] = {
                    "alertlevel",
                    "callback",
                    "catalog",
                    "contributor",
                    "endtime",
                    "eventid",
                    "eventtype",
                    "format",
                    "includeallmagnitude",
                    "includeallorigin",
                    "includearrival",
                    "includedeleted",
                    "includesuperseded",
                    "jsonerror",
                    "kmlanimated",
                    "kmlcolorby",
                    "latitude",
                    "limit",
                    "longitude",
                    "magnitudetype",
                    "maxcdi",
                    "maxdepth",
                    "maxgap",
                    "maxlatitude",
                    "maxlongitude",
                    "maxmagnitude",
                    "maxmmi",
                    "maxradius",
                    "maxradiuskm",
                    "maxsig",
                    "mincdi",
                    "mindepth",
                    "minfelt",
                    "mingap",
                    "minlatitude",
                    "minlongitude",
                    "minmagnitude",
                    "minsig",
                    "nodata",
                    "offset",
                    "orderby",
                    "productcode",
                    "producttype",
                    "reviewstatus",
                    "starttime",
                    "updatedafter",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "alertlevel",
                      ["orig"] = "alertlevel",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "catalog",
                      ["orig"] = "catalog",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "contributor",
                      ["orig"] = "contributor",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "endtime",
                      ["orig"] = "endtime",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "eventtype",
                      ["orig"] = "eventtype",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "text",
                      ["kind"] = "query",
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "latitude",
                      ["orig"] = "latitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "longitude",
                      ["orig"] = "longitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "magnitudetype",
                      ["orig"] = "magnitudetype",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = 1000,
                      ["kind"] = "query",
                      ["name"] = "maxdepth",
                      ["orig"] = "maxdepth",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 90,
                      ["kind"] = "query",
                      ["name"] = "maxlatitude",
                      ["orig"] = "maxlatitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 180,
                      ["kind"] = "query",
                      ["name"] = "maxlongitude",
                      ["orig"] = "maxlongitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "maxmagnitude",
                      ["orig"] = "maxmagnitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 180,
                      ["kind"] = "query",
                      ["name"] = "maxradius",
                      ["orig"] = "maxradius",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 20001.6,
                      ["kind"] = "query",
                      ["name"] = "maxradiuskm",
                      ["orig"] = "maxradiuskm",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = -100,
                      ["kind"] = "query",
                      ["name"] = "mindepth",
                      ["orig"] = "mindepth",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = -90,
                      ["kind"] = "query",
                      ["name"] = "minlatitude",
                      ["orig"] = "minlatitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = -180,
                      ["kind"] = "query",
                      ["name"] = "minlongitude",
                      ["orig"] = "minlongitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "minmagnitude",
                      ["orig"] = "minmagnitude",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = "all",
                      ["kind"] = "query",
                      ["name"] = "reviewstatus",
                      ["orig"] = "reviewstatus",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "starttime",
                      ["orig"] = "starttime",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "updatedafter",
                      ["orig"] = "updatedafter",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/count",
                ["parts"] = {
                  "count",
                },
                ["select"] = {
                  ["exist"] = {
                    "alertlevel",
                    "catalog",
                    "contributor",
                    "endtime",
                    "eventtype",
                    "format",
                    "latitude",
                    "longitude",
                    "magnitudetype",
                    "maxdepth",
                    "maxlatitude",
                    "maxlongitude",
                    "maxmagnitude",
                    "maxradius",
                    "maxradiuskm",
                    "mindepth",
                    "minlatitude",
                    "minlongitude",
                    "minmagnitude",
                    "reviewstatus",
                    "starttime",
                    "updatedafter",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["service_information"] = {
        ["fields"] = {},
        ["name"] = "service_information",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/catalogs",
                ["parts"] = {
                  "catalogs",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/contributors",
                ["parts"] = {
                  "contributors",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/application.json",
                ["parts"] = {
                  "application.json",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/application.wadl",
                ["parts"] = {
                  "application.wadl",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/version",
                ["parts"] = {
                  "version",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
    },
  }
end


local function make_feature(name)
  local features = require("features")
  local factory = features[name]
  if factory ~= nil then
    return factory()
  end
  return features.base()
end


-- Attach make_feature to the SDK class
local function setup_sdk(SDK)
  SDK._make_feature = make_feature
end


return make_config
