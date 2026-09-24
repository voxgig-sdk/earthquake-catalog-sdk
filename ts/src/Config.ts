
import { BaseFeature } from './feature/base/BaseFeature'
import { RatelimitFeature } from './feature/ratelimit/RatelimitFeature'
import { RetryFeature } from './feature/retry/RetryFeature'
import { TestFeature } from './feature/test/TestFeature'
import { TimeoutFeature } from './feature/timeout/TimeoutFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   ratelimit: RatelimitFeature,
 retry: RetryFeature,
 test: TestFeature,
 timeout: TimeoutFeature,

}


const FEATURE_PLUGINS: Record<string, any[]> = {
  
}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    return fi
  }

  // False for a feature added at runtime via options.extend (station's
  // adopt path) - the constructor uses this to skip makeFeature for names
  // no generated class backs.
  hasFeature(this: any, fn: string) {
    return null != FEATURE_CLASS[fn]
  }


  main = {
    name: 'EarthquakeCatalog',
        slug: "earthquake-catalog",
    version: "0.0.1",
    target: "ts",

  }


  feature = {
     ratelimit:     {
      "options": {
        "active": false,
        "burst": 5,
        "rate": 5
      },
      "optspec": {
        "now": "`$FUNCTION`",
        "sleep": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },
 retry:     {
      "options": {
        "active": false,
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
          504
        ]
      },
      "optspec": {
        "jitter": "`$BOOLEAN`",
        "sleep": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },
 test:     {
      "options": {
        "active": false
      },
      "optspec": {
        "entity": "`$MAP`",
        "net": "`$MAP`"
      },
      "strict": false,
      "transport": "base"
    },
 timeout:     {
      "options": {
        "active": false,
        "ms": 30000
      },
      "optspec": {
        "clearTimer": "`$FUNCTION`",
        "setTimer": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },

  }


  options = {
    base: "https://earthquake.usgs.gov/fdsnws/event/1",

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
        earthquake_data: {
        },
  
        service_information: {
        },
  
    }
  }


  entity = {
    "earthquake_data": {
      "fields": [
        {
          "name": "count",
          "title": "Count",
          "type": "`$INTEGER`"
        },
        {
          "name": "geometry",
          "title": "Geometry",
          "type": "`$OBJECT`"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`"
        },
        {
          "name": "maxAllowed",
          "title": "Max Allowed",
          "type": "`$INTEGER`"
        },
        {
          "name": "properties",
          "title": "Properties",
          "type": "`$OBJECT`"
        },
        {
          "name": "type",
          "title": "Type",
          "type": "`$STRING`"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
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
                  "lit": "query"
                }
              ],
              "parts": [
                "query"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "alertlevel",
                    "orig": "alertlevel",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "callback",
                    "orig": "callback",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "catalog",
                    "orig": "catalog",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "contributor",
                    "orig": "contributor",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "endtime",
                    "orig": "endtime",
                    "type": "`$STRING`",
                    "kind": "query",
                    "example": "2014-01-02"
                  },
                  {
                    "name": "eventid",
                    "orig": "eventid",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "eventtype",
                    "orig": "eventtype",
                    "type": "`$STRING`",
                    "kind": "query",
                    "example": "earthquake"
                  },
                  {
                    "name": "format",
                    "orig": "format",
                    "type": "`$STRING`",
                    "kind": "query",
                    "example": "quakeml"
                  },
                  {
                    "name": "includeallmagnitude",
                    "orig": "includeallmagnitude",
                    "type": "`$BOOLEAN`",
                    "kind": "query",
                    "example": false
                  },
                  {
                    "name": "includeallorigin",
                    "orig": "includeallorigin",
                    "type": "`$BOOLEAN`",
                    "kind": "query",
                    "example": false
                  },
                  {
                    "name": "includearrival",
                    "orig": "includearrival",
                    "type": "`$BOOLEAN`",
                    "kind": "query",
                    "example": false
                  },
                  {
                    "name": "includedeleted",
                    "orig": "includedeleted",
                    "type": "`$STRING`",
                    "kind": "query",
                    "example": "false"
                  },
                  {
                    "name": "includesuperseded",
                    "orig": "includesuperseded",
                    "type": "`$BOOLEAN`",
                    "kind": "query",
                    "example": false
                  },
                  {
                    "name": "jsonerror",
                    "orig": "jsonerror",
                    "type": "`$BOOLEAN`",
                    "kind": "query",
                    "example": false
                  },
                  {
                    "name": "kmlanimated",
                    "orig": "kmlanimated",
                    "type": "`$BOOLEAN`",
                    "kind": "query",
                    "example": false
                  },
                  {
                    "name": "kmlcolorby",
                    "orig": "kmlcolorby",
                    "type": "`$STRING`",
                    "kind": "query",
                    "example": "age"
                  },
                  {
                    "name": "latitude",
                    "orig": "latitude",
                    "type": "`$NUMBER`",
                    "kind": "query"
                  },
                  {
                    "name": "limit",
                    "orig": "limit",
                    "type": "`$INTEGER`",
                    "kind": "query"
                  },
                  {
                    "name": "longitude",
                    "orig": "longitude",
                    "type": "`$NUMBER`",
                    "kind": "query"
                  },
                  {
                    "name": "magnitudetype",
                    "orig": "magnitudetype",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "maxcdi",
                    "orig": "maxcdi",
                    "type": "`$NUMBER`",
                    "kind": "query"
                  },
                  {
                    "name": "maxdepth",
                    "orig": "maxdepth",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": 1000
                  },
                  {
                    "name": "maxgap",
                    "orig": "maxgap",
                    "type": "`$NUMBER`",
                    "kind": "query"
                  },
                  {
                    "name": "maxlatitude",
                    "orig": "maxlatitude",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": 90
                  },
                  {
                    "name": "maxlongitude",
                    "orig": "maxlongitude",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": 180
                  },
                  {
                    "name": "maxmagnitude",
                    "orig": "maxmagnitude",
                    "type": "`$NUMBER`",
                    "kind": "query"
                  },
                  {
                    "name": "maxmmi",
                    "orig": "maxmmi",
                    "type": "`$NUMBER`",
                    "kind": "query"
                  },
                  {
                    "name": "maxradius",
                    "orig": "maxradius",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": 180
                  },
                  {
                    "name": "maxradiuskm",
                    "orig": "maxradiuskm",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": 20001.6
                  },
                  {
                    "name": "maxsig",
                    "orig": "maxsig",
                    "type": "`$INTEGER`",
                    "kind": "query"
                  },
                  {
                    "name": "mincdi",
                    "orig": "mincdi",
                    "type": "`$NUMBER`",
                    "kind": "query"
                  },
                  {
                    "name": "mindepth",
                    "orig": "mindepth",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": -100
                  },
                  {
                    "name": "minfelt",
                    "orig": "minfelt",
                    "type": "`$INTEGER`",
                    "kind": "query"
                  },
                  {
                    "name": "mingap",
                    "orig": "mingap",
                    "type": "`$NUMBER`",
                    "kind": "query"
                  },
                  {
                    "name": "minlatitude",
                    "orig": "minlatitude",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": -90
                  },
                  {
                    "name": "minlongitude",
                    "orig": "minlongitude",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": -180
                  },
                  {
                    "name": "minmagnitude",
                    "orig": "minmagnitude",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": 5
                  },
                  {
                    "name": "minsig",
                    "orig": "minsig",
                    "type": "`$INTEGER`",
                    "kind": "query"
                  },
                  {
                    "name": "nodata",
                    "orig": "nodata",
                    "type": "`$INTEGER`",
                    "kind": "query",
                    "example": 204
                  },
                  {
                    "name": "offset",
                    "orig": "offset",
                    "type": "`$INTEGER`",
                    "kind": "query",
                    "example": 1
                  },
                  {
                    "name": "orderby",
                    "orig": "orderby",
                    "type": "`$STRING`",
                    "kind": "query",
                    "example": "time"
                  },
                  {
                    "name": "productcode",
                    "orig": "productcode",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "producttype",
                    "orig": "producttype",
                    "type": "`$STRING`",
                    "kind": "query",
                    "example": "moment-tensor"
                  },
                  {
                    "name": "reviewstatus",
                    "orig": "reviewstatus",
                    "type": "`$STRING`",
                    "kind": "query",
                    "example": "all"
                  },
                  {
                    "name": "starttime",
                    "orig": "starttime",
                    "type": "`$STRING`",
                    "kind": "query",
                    "example": "2014-01-01"
                  },
                  {
                    "name": "updatedafter",
                    "orig": "updatedafter",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
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
                  "updatedafter"
                ]
              }
            }
          ]
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
                  "lit": "count"
                }
              ],
              "parts": [
                "count"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "alertlevel",
                    "orig": "alertlevel",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "catalog",
                    "orig": "catalog",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "contributor",
                    "orig": "contributor",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "endtime",
                    "orig": "endtime",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "eventtype",
                    "orig": "eventtype",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "format",
                    "orig": "format",
                    "type": "`$STRING`",
                    "kind": "query",
                    "example": "text"
                  },
                  {
                    "name": "latitude",
                    "orig": "latitude",
                    "type": "`$NUMBER`",
                    "kind": "query"
                  },
                  {
                    "name": "longitude",
                    "orig": "longitude",
                    "type": "`$NUMBER`",
                    "kind": "query"
                  },
                  {
                    "name": "magnitudetype",
                    "orig": "magnitudetype",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "maxdepth",
                    "orig": "maxdepth",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": 1000
                  },
                  {
                    "name": "maxlatitude",
                    "orig": "maxlatitude",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": 90
                  },
                  {
                    "name": "maxlongitude",
                    "orig": "maxlongitude",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": 180
                  },
                  {
                    "name": "maxmagnitude",
                    "orig": "maxmagnitude",
                    "type": "`$NUMBER`",
                    "kind": "query"
                  },
                  {
                    "name": "maxradius",
                    "orig": "maxradius",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": 180
                  },
                  {
                    "name": "maxradiuskm",
                    "orig": "maxradiuskm",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": 20001.6
                  },
                  {
                    "name": "mindepth",
                    "orig": "mindepth",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": -100
                  },
                  {
                    "name": "minlatitude",
                    "orig": "minlatitude",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": -90
                  },
                  {
                    "name": "minlongitude",
                    "orig": "minlongitude",
                    "type": "`$NUMBER`",
                    "kind": "query",
                    "example": -180
                  },
                  {
                    "name": "minmagnitude",
                    "orig": "minmagnitude",
                    "type": "`$NUMBER`",
                    "kind": "query"
                  },
                  {
                    "name": "reviewstatus",
                    "orig": "reviewstatus",
                    "type": "`$STRING`",
                    "kind": "query",
                    "example": "all"
                  },
                  {
                    "name": "starttime",
                    "orig": "starttime",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "updatedafter",
                    "orig": "updatedafter",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
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
                  "updatedafter"
                ]
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
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
                  "lit": "catalogs"
                }
              ],
              "parts": [
                "catalogs"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/contributors",
              "segments": [
                {
                  "lit": "contributors"
                }
              ],
              "parts": [
                "contributors"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
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
                  "lit": "application.json"
                }
              ],
              "parts": [
                "application.json"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/application.wadl",
              "segments": [
                {
                  "lit": "application.wadl"
                }
              ],
              "parts": [
                "application.wadl"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/version",
              "segments": [
                {
                  "lit": "version"
                }
              ],
              "parts": [
                "version"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    }
  }
}


const config = new Config()

export {
  config,
  FEATURE_PLUGINS,
}

