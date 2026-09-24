# EarthquakeCatalog SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "EarthquakeCatalog",
            "slug": "earthquake-catalog",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "ratelimit": {
        "options": {
          "active": False,
          "burst": 5,
          "rate": 5,
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "retry": {
        "options": {
          "active": False,
          "factor": 2,
          "maxDelay": 2000,
          "minDelay": 50,
          "retries": 2,
          "statuses": [
            408,
            425,
            429,
            500,
            502,
            503,
            504,
          ],
        },
        "optspec": {
          "jitter": "`$BOOLEAN`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "test": {
        "options": {
          "active": False,
        },
        "optspec": {
          "entity": "`$MAP`",
          "net": "`$MAP`",
        },
        "strict": False,
        "transport": "base",
      },
            "timeout": {
        "options": {
          "active": False,
          "ms": 30000,
        },
        "optspec": {
          "clearTimer": "`$FUNCTION`",
          "setTimer": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
        },
        "options": {
            "base": "https://earthquake.usgs.gov/fdsnws/event/1",
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "earthquake_data": {},
                "service_information": {},
            },
        },
        "entity": {
      "earthquake_data": {
        "fields": [
          {
            "name": "count",
            "title": "Count",
            "type": "`$INTEGER`",
          },
          {
            "name": "geometry",
            "title": "Geometry",
            "type": "`$OBJECT`",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
          },
          {
            "name": "maxAllowed",
            "title": "Max Allowed",
            "type": "`$INTEGER`",
          },
          {
            "name": "properties",
            "title": "Properties",
            "type": "`$OBJECT`",
          },
          {
            "name": "type",
            "title": "Type",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "earthquake_data",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/query",
                "segments": [
                  {
                    "lit": "query",
                  },
                ],
                "parts": [
                  "query",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "query": [
                    {
                      "name": "alertlevel",
                      "orig": "alertlevel",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "callback",
                      "orig": "callback",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "catalog",
                      "orig": "catalog",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "contributor",
                      "orig": "contributor",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "endtime",
                      "orig": "endtime",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2014-01-02",
                    },
                    {
                      "name": "eventid",
                      "orig": "eventid",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "eventtype",
                      "orig": "eventtype",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "earthquake",
                    },
                    {
                      "name": "format",
                      "orig": "format",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "quakeml",
                    },
                    {
                      "name": "includeallmagnitude",
                      "orig": "includeallmagnitude",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                      "example": False,
                    },
                    {
                      "name": "includeallorigin",
                      "orig": "includeallorigin",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                      "example": False,
                    },
                    {
                      "name": "includearrival",
                      "orig": "includearrival",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                      "example": False,
                    },
                    {
                      "name": "includedeleted",
                      "orig": "includedeleted",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "false",
                    },
                    {
                      "name": "includesuperseded",
                      "orig": "includesuperseded",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                      "example": False,
                    },
                    {
                      "name": "jsonerror",
                      "orig": "jsonerror",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                      "example": False,
                    },
                    {
                      "name": "kmlanimated",
                      "orig": "kmlanimated",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                      "example": False,
                    },
                    {
                      "name": "kmlcolorby",
                      "orig": "kmlcolorby",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "age",
                    },
                    {
                      "name": "latitude",
                      "orig": "latitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                    },
                    {
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "longitude",
                      "orig": "longitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                    },
                    {
                      "name": "magnitudetype",
                      "orig": "magnitudetype",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "maxcdi",
                      "orig": "maxcdi",
                      "type": "`$NUMBER`",
                      "kind": "query",
                    },
                    {
                      "name": "maxdepth",
                      "orig": "maxdepth",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": 1000,
                    },
                    {
                      "name": "maxgap",
                      "orig": "maxgap",
                      "type": "`$NUMBER`",
                      "kind": "query",
                    },
                    {
                      "name": "maxlatitude",
                      "orig": "maxlatitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": 90,
                    },
                    {
                      "name": "maxlongitude",
                      "orig": "maxlongitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": 180,
                    },
                    {
                      "name": "maxmagnitude",
                      "orig": "maxmagnitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                    },
                    {
                      "name": "maxmmi",
                      "orig": "maxmmi",
                      "type": "`$NUMBER`",
                      "kind": "query",
                    },
                    {
                      "name": "maxradius",
                      "orig": "maxradius",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": 180,
                    },
                    {
                      "name": "maxradiuskm",
                      "orig": "maxradiuskm",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": 20001.6,
                    },
                    {
                      "name": "maxsig",
                      "orig": "maxsig",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "mincdi",
                      "orig": "mincdi",
                      "type": "`$NUMBER`",
                      "kind": "query",
                    },
                    {
                      "name": "mindepth",
                      "orig": "mindepth",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": -100,
                    },
                    {
                      "name": "minfelt",
                      "orig": "minfelt",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "mingap",
                      "orig": "mingap",
                      "type": "`$NUMBER`",
                      "kind": "query",
                    },
                    {
                      "name": "minlatitude",
                      "orig": "minlatitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": -90,
                    },
                    {
                      "name": "minlongitude",
                      "orig": "minlongitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": -180,
                    },
                    {
                      "name": "minmagnitude",
                      "orig": "minmagnitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": 5,
                    },
                    {
                      "name": "minsig",
                      "orig": "minsig",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "nodata",
                      "orig": "nodata",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 204,
                    },
                    {
                      "name": "offset",
                      "orig": "offset",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 1,
                    },
                    {
                      "name": "orderby",
                      "orig": "orderby",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "time",
                    },
                    {
                      "name": "productcode",
                      "orig": "productcode",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "producttype",
                      "orig": "producttype",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "moment-tensor",
                    },
                    {
                      "name": "reviewstatus",
                      "orig": "reviewstatus",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "all",
                    },
                    {
                      "name": "starttime",
                      "orig": "starttime",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2014-01-01",
                    },
                    {
                      "name": "updatedafter",
                      "orig": "updatedafter",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
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
                  ],
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/count",
                "segments": [
                  {
                    "lit": "count",
                  },
                ],
                "parts": [
                  "count",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "query": [
                    {
                      "name": "alertlevel",
                      "orig": "alertlevel",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "catalog",
                      "orig": "catalog",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "contributor",
                      "orig": "contributor",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "endtime",
                      "orig": "endtime",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "eventtype",
                      "orig": "eventtype",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "format",
                      "orig": "format",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "text",
                    },
                    {
                      "name": "latitude",
                      "orig": "latitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                    },
                    {
                      "name": "longitude",
                      "orig": "longitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                    },
                    {
                      "name": "magnitudetype",
                      "orig": "magnitudetype",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "maxdepth",
                      "orig": "maxdepth",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": 1000,
                    },
                    {
                      "name": "maxlatitude",
                      "orig": "maxlatitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": 90,
                    },
                    {
                      "name": "maxlongitude",
                      "orig": "maxlongitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": 180,
                    },
                    {
                      "name": "maxmagnitude",
                      "orig": "maxmagnitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                    },
                    {
                      "name": "maxradius",
                      "orig": "maxradius",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": 180,
                    },
                    {
                      "name": "maxradiuskm",
                      "orig": "maxradiuskm",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": 20001.6,
                    },
                    {
                      "name": "mindepth",
                      "orig": "mindepth",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": -100,
                    },
                    {
                      "name": "minlatitude",
                      "orig": "minlatitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": -90,
                    },
                    {
                      "name": "minlongitude",
                      "orig": "minlongitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                      "example": -180,
                    },
                    {
                      "name": "minmagnitude",
                      "orig": "minmagnitude",
                      "type": "`$NUMBER`",
                      "kind": "query",
                    },
                    {
                      "name": "reviewstatus",
                      "orig": "reviewstatus",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "all",
                    },
                    {
                      "name": "starttime",
                      "orig": "starttime",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updatedafter",
                      "orig": "updatedafter",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
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
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "service_information": {
        "fields": [],
        "name": "service_information",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/catalogs",
                "segments": [
                  {
                    "lit": "catalogs",
                  },
                ],
                "parts": [
                  "catalogs",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/contributors",
                "segments": [
                  {
                    "lit": "contributors",
                  },
                ],
                "parts": [
                  "contributors",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/application.json",
                "segments": [
                  {
                    "lit": "application.json",
                  },
                ],
                "parts": [
                  "application.json",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/application.wadl",
                "segments": [
                  {
                    "lit": "application.wadl",
                  },
                ],
                "parts": [
                  "application.wadl",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/version",
                "segments": [
                  {
                    "lit": "version",
                  },
                ],
                "parts": [
                  "version",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
