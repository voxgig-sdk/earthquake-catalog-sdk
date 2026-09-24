

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { EarthquakeCatalogSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


loadEnvLocal(__dirname + '/../../../.env.local')


describe('EarthquakeDataEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when EARTHQUAKE_CATALOG_TEST_LIVE=TRUE.
  afterEach(liveDelay('EARTHQUAKE_CATALOG_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = EarthquakeCatalogSDK.test()
    const ent = testsdk.EarthquakeData()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.EARTHQUAKE_CATALOG_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'earthquake_data.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"count":{"a":true,"h":"Count","n":"count","r":false,"t":"`$INTEGER`","key$":"count","index$":0},"geometry":{"a":true,"h":"Geometry","n":"geometry","r":false,"t":"`$OBJECT`","key$":"geometry","index$":1},"id":{"a":true,"h":"Id","n":"id","r":false,"t":"`$STRING`","key$":"id","index$":2},"maxAllowed":{"a":true,"h":"Max Allowed","n":"maxAllowed","r":false,"t":"`$INTEGER`","key$":"maxAllowed","index$":3},"properties":{"a":true,"h":"Properties","n":"properties","r":false,"t":"`$OBJECT`","key$":"properties","index$":4},"type":{"a":true,"h":"Type","n":"type","r":false,"t":"`$STRING`","key$":"type","index$":5}},"id":{"field":"id","name":"id"},"name":"earthquake_data","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /query","source":"openapi3","version":2},"g":{"query":[{"a":true,"k":"query","n":"alertlevel","or":"alertlevel","r":false,"t":"`$STRING`","index$":0},{"a":true,"k":"query","n":"callback","or":"callback","r":false,"t":"`$STRING`","index$":1},{"a":true,"k":"query","n":"catalog","or":"catalog","r":false,"t":"`$STRING`","index$":2},{"a":true,"k":"query","n":"contributor","or":"contributor","r":false,"t":"`$STRING`","index$":3},{"a":true,"ex":"2014-01-02","k":"query","n":"endtime","or":"endtime","r":false,"t":"`$STRING`","index$":4},{"a":true,"k":"query","n":"eventid","or":"eventid","r":false,"t":"`$STRING`","index$":5},{"a":true,"ex":"earthquake","k":"query","n":"eventtype","or":"eventtype","r":false,"t":"`$STRING`","index$":6},{"a":true,"ex":"quakeml","k":"query","n":"format","or":"format","r":false,"t":"`$STRING`","index$":7},{"a":true,"ex":false,"k":"query","n":"includeallmagnitude","or":"includeallmagnitude","r":false,"t":"`$BOOLEAN`","index$":8},{"a":true,"ex":false,"k":"query","n":"includeallorigin","or":"includeallorigin","r":false,"t":"`$BOOLEAN`","index$":9},{"a":true,"ex":false,"k":"query","n":"includearrival","or":"includearrival","r":false,"t":"`$BOOLEAN`","index$":10},{"a":true,"ex":"false","k":"query","n":"includedeleted","or":"includedeleted","r":false,"t":"`$STRING`","index$":11},{"a":true,"ex":false,"k":"query","n":"includesuperseded","or":"includesuperseded","r":false,"t":"`$BOOLEAN`","index$":12},{"a":true,"ex":false,"k":"query","n":"jsonerror","or":"jsonerror","r":false,"t":"`$BOOLEAN`","index$":13},{"a":true,"ex":false,"k":"query","n":"kmlanimated","or":"kmlanimated","r":false,"t":"`$BOOLEAN`","index$":14},{"a":true,"ex":"age","k":"query","n":"kmlcolorby","or":"kmlcolorby","r":false,"t":"`$STRING`","index$":15},{"a":true,"k":"query","n":"latitude","or":"latitude","r":false,"t":"`$NUMBER`","index$":16},{"a":true,"k":"query","n":"limit","or":"limit","r":false,"t":"`$INTEGER`","index$":17},{"a":true,"k":"query","n":"longitude","or":"longitude","r":false,"t":"`$NUMBER`","index$":18},{"a":true,"k":"query","n":"magnitudetype","or":"magnitudetype","r":false,"t":"`$STRING`","index$":19},{"a":true,"k":"query","n":"maxcdi","or":"maxcdi","r":false,"t":"`$NUMBER`","index$":20},{"a":true,"ex":1000,"k":"query","n":"maxdepth","or":"maxdepth","r":false,"t":"`$NUMBER`","index$":21},{"a":true,"k":"query","n":"maxgap","or":"maxgap","r":false,"t":"`$NUMBER`","index$":22},{"a":true,"ex":90,"k":"query","n":"maxlatitude","or":"maxlatitude","r":false,"t":"`$NUMBER`","index$":23},{"a":true,"ex":180,"k":"query","n":"maxlongitude","or":"maxlongitude","r":false,"t":"`$NUMBER`","index$":24},{"a":true,"k":"query","n":"maxmagnitude","or":"maxmagnitude","r":false,"t":"`$NUMBER`","index$":25},{"a":true,"k":"query","n":"maxmmi","or":"maxmmi","r":false,"t":"`$NUMBER`","index$":26},{"a":true,"ex":180,"k":"query","n":"maxradius","or":"maxradius","r":false,"t":"`$NUMBER`","index$":27},{"a":true,"ex":20001.6,"k":"query","n":"maxradiuskm","or":"maxradiuskm","r":false,"t":"`$NUMBER`","index$":28},{"a":true,"k":"query","n":"maxsig","or":"maxsig","r":false,"t":"`$INTEGER`","index$":29},{"a":true,"k":"query","n":"mincdi","or":"mincdi","r":false,"t":"`$NUMBER`","index$":30},{"a":true,"ex":-100,"k":"query","n":"mindepth","or":"mindepth","r":false,"t":"`$NUMBER`","index$":31},{"a":true,"k":"query","n":"minfelt","or":"minfelt","r":false,"t":"`$INTEGER`","index$":32},{"a":true,"k":"query","n":"mingap","or":"mingap","r":false,"t":"`$NUMBER`","index$":33},{"a":true,"ex":-90,"k":"query","n":"minlatitude","or":"minlatitude","r":false,"t":"`$NUMBER`","index$":34},{"a":true,"ex":-180,"k":"query","n":"minlongitude","or":"minlongitude","r":false,"t":"`$NUMBER`","index$":35},{"a":true,"ex":5,"k":"query","n":"minmagnitude","or":"minmagnitude","r":false,"t":"`$NUMBER`","index$":36},{"a":true,"k":"query","n":"minsig","or":"minsig","r":false,"t":"`$INTEGER`","index$":37},{"a":true,"ex":204,"k":"query","n":"nodata","or":"nodata","r":false,"t":"`$INTEGER`","index$":38},{"a":true,"ex":1,"k":"query","n":"offset","or":"offset","r":false,"t":"`$INTEGER`","index$":39},{"a":true,"ex":"time","k":"query","n":"orderby","or":"orderby","r":false,"t":"`$STRING`","index$":40},{"a":true,"k":"query","n":"productcode","or":"productcode","r":false,"t":"`$STRING`","index$":41},{"a":true,"ex":"moment-tensor","k":"query","n":"producttype","or":"producttype","r":false,"t":"`$STRING`","index$":42},{"a":true,"ex":"all","k":"query","n":"reviewstatus","or":"reviewstatus","r":false,"t":"`$STRING`","index$":43},{"a":true,"ex":"2014-01-01","k":"query","n":"starttime","or":"starttime","r":false,"t":"`$STRING`","index$":44},{"a":true,"k":"query","n":"updatedafter","or":"updatedafter","r":false,"t":"`$STRING`","index$":45}]},"k":"http","m":"GET","o":"/query","q":{"exist":["alertlevel","callback","catalog","contributor","endtime","eventid","eventtype","format","includeallmagnitude","includeallorigin","includearrival","includedeleted","includesuperseded","jsonerror","kmlanimated","kmlcolorby","latitude","limit","longitude","magnitudetype","maxcdi","maxdepth","maxgap","maxlatitude","maxlongitude","maxmagnitude","maxmmi","maxradius","maxradiuskm","maxsig","mincdi","mindepth","minfelt","mingap","minlatitude","minlongitude","minmagnitude","minsig","nodata","offset","orderby","productcode","producttype","reviewstatus","starttime","updatedafter"]},"r":{},"s":[{"lit":"query"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /count","source":"openapi3","version":2},"g":{"query":[{"a":true,"k":"query","n":"alertlevel","or":"alertlevel","r":false,"t":"`$STRING`","index$":0},{"a":true,"k":"query","n":"catalog","or":"catalog","r":false,"t":"`$STRING`","index$":1},{"a":true,"k":"query","n":"contributor","or":"contributor","r":false,"t":"`$STRING`","index$":2},{"a":true,"k":"query","n":"endtime","or":"endtime","r":false,"t":"`$STRING`","index$":3},{"a":true,"k":"query","n":"eventtype","or":"eventtype","r":false,"t":"`$STRING`","index$":4},{"a":true,"ex":"text","k":"query","n":"format","or":"format","r":false,"t":"`$STRING`","index$":5},{"a":true,"k":"query","n":"latitude","or":"latitude","r":false,"t":"`$NUMBER`","index$":6},{"a":true,"k":"query","n":"longitude","or":"longitude","r":false,"t":"`$NUMBER`","index$":7},{"a":true,"k":"query","n":"magnitudetype","or":"magnitudetype","r":false,"t":"`$STRING`","index$":8},{"a":true,"ex":1000,"k":"query","n":"maxdepth","or":"maxdepth","r":false,"t":"`$NUMBER`","index$":9},{"a":true,"ex":90,"k":"query","n":"maxlatitude","or":"maxlatitude","r":false,"t":"`$NUMBER`","index$":10},{"a":true,"ex":180,"k":"query","n":"maxlongitude","or":"maxlongitude","r":false,"t":"`$NUMBER`","index$":11},{"a":true,"k":"query","n":"maxmagnitude","or":"maxmagnitude","r":false,"t":"`$NUMBER`","index$":12},{"a":true,"ex":180,"k":"query","n":"maxradius","or":"maxradius","r":false,"t":"`$NUMBER`","index$":13},{"a":true,"ex":20001.6,"k":"query","n":"maxradiuskm","or":"maxradiuskm","r":false,"t":"`$NUMBER`","index$":14},{"a":true,"ex":-100,"k":"query","n":"mindepth","or":"mindepth","r":false,"t":"`$NUMBER`","index$":15},{"a":true,"ex":-90,"k":"query","n":"minlatitude","or":"minlatitude","r":false,"t":"`$NUMBER`","index$":16},{"a":true,"ex":-180,"k":"query","n":"minlongitude","or":"minlongitude","r":false,"t":"`$NUMBER`","index$":17},{"a":true,"k":"query","n":"minmagnitude","or":"minmagnitude","r":false,"t":"`$NUMBER`","index$":18},{"a":true,"ex":"all","k":"query","n":"reviewstatus","or":"reviewstatus","r":false,"t":"`$STRING`","index$":19},{"a":true,"k":"query","n":"starttime","or":"starttime","r":false,"t":"`$STRING`","index$":20},{"a":true,"k":"query","n":"updatedafter","or":"updatedafter","r":false,"t":"`$STRING`","index$":21}]},"k":"http","m":"GET","o":"/count","q":{"exist":["alertlevel","catalog","contributor","endtime","eventtype","format","latitude","longitude","magnitudetype","maxdepth","maxlatitude","maxlongitude","maxmagnitude","maxradius","maxradiuskm","mindepth","minlatitude","minlongitude","minmagnitude","reviewstatus","starttime","updatedafter"]},"r":{},"s":[{"lit":"count"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"earthquake_data","name__orig":"earthquake_data","Name":"EarthquakeData","name_":"earthquake_data","name-":"earthquake-data","NAME":"EARTHQUAKE_DATA","index$":0}, {"active":true,"entity":"earthquake_data","key$":"BasicEarthquakeDataFlow","kind":"basic","name":"BasicEarthquakeDataFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"earthquake_data_ref01"}}],"index$":0},{"a":true,"d":{},"i":{"ref":"earthquake_data_ref01","srcdatavar":"earthquake_data_ref01_data","suffix":"_dt0"},"m":{},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-earthquake_data_ref01"}}],"index$":1}]}, 'EarthquakeData', {"GET /query":{"protocol":"http","operationId":"queryEarthquakes","responses":{"200":{"description":"Successful response with earthquake data","content":{"application/json":{"schema":{"type":"object","properties":{"type":{"example":"FeatureCollection","key$":"type","type":"string"},"metadata":{"key$":"metadata","properties":{"api":{"type":"string"},"count":{"type":"integer"},"generated":{"format":"int64","type":"integer"},"status":{"type":"integer"},"title":{"type":"string"},"url":{"type":"string"}},"type":"object"},"features":{"items":{"properties":{"geometry":{"properties":{"coordinates":{"items":{"type":"number"},"maxItems":3,"minItems":3,"type":"array"},"type":{"example":"Point","type":"string"}},"type":"object","key$":"geometry"},"id":{"type":"string","key$":"id"},"properties":{"properties":{"alert":{"type":"string"},"cdi":{"type":"number"},"code":{"type":"string"},"detail":{"type":"string"},"dmin":{"type":"number"},"felt":{"type":"integer"},"gap":{"type":"number"},"ids":{"type":"string"},"mag":{"type":"number"},"magType":{"type":"string"},"mmi":{"type":"number"},"net":{"type":"string"},"nst":{"type":"integer"},"place":{"type":"string"},"rms":{"type":"number"},"sig":{"type":"integer"},"sources":{"type":"string"},"status":{"type":"string"},"time":{"format":"int64","type":"integer"},"title":{"type":"string"},"tsunami":{"type":"integer"},"type":{"type":"string"},"types":{"type":"string"},"tz":{"type":"integer"},"updated":{"format":"int64","type":"integer"},"url":{"type":"string"}},"type":"object","key$":"properties"},"type":{"example":"Feature","type":"string","key$":"type"}},"type":"object","index$":0},"key$":"features","type":"array"},"bbox":{"items":{"type":"number"},"key$":"bbox","type":"array"}}}},"application/xml":{"schema":{"type":"object","description":"QuakeML 1.2 formatted earthquake data"}},"text/csv":{"schema":{"type":"string"}},"application/vnd.google-earth.kml+xml":{"schema":{"type":"object","description":"KML formatted earthquake data"}},"text/plain":{"schema":{"type":"string"}}}},"204":{"description":"No data found (when nodata=204)"},"400":{"description":"Bad request - invalid parameters"},"404":{"description":"No data found (when nodata=404)"},"409":{"description":"Conflict - deleted event (when includedeleted is not specified)"}},"parameters":[{"name":"format","in":"query","description":"Specify the output format","schema":{"type":"string","enum":["csv","geojson","kml","quakeml","text","xml"],"default":"quakeml"},"index$":0},{"name":"starttime","in":"query","description":"Limit to events on or after the specified start time. ISO8601 Date/Time format. UTC assumed if timezone not specified.","schema":{"type":"string","format":"date-time"},"example":"2014-01-01","index$":1},{"name":"endtime","in":"query","description":"Limit to events on or before the specified end time. ISO8601 Date/Time format. UTC assumed if timezone not specified.","schema":{"type":"string","format":"date-time"},"example":"2014-01-02","index$":2},{"name":"updatedafter","in":"query","description":"Limit to events updated after the specified time. ISO8601 Date/Time format. UTC assumed if timezone not specified.","schema":{"type":"string","format":"date-time"},"index$":3},{"name":"minlatitude","in":"query","description":"Limit to events with a latitude larger than the specified minimum. Range: -90 to 90 degrees.","schema":{"type":"number","format":"double","minimum":-90,"maximum":90,"default":-90},"index$":4},{"name":"maxlatitude","in":"query","description":"Limit to events with a latitude smaller than the specified maximum. Range: -90 to 90 degrees.","schema":{"type":"number","format":"double","minimum":-90,"maximum":90,"default":90},"index$":5},{"name":"minlongitude","in":"query","description":"Limit to events with a longitude larger than the specified minimum. Range: -360 to 360 degrees.","schema":{"type":"number","format":"double","minimum":-360,"maximum":360,"default":-180},"index$":6},{"name":"maxlongitude","in":"query","description":"Limit to events with a longitude smaller than the specified maximum. Range: -360 to 360 degrees.","schema":{"type":"number","format":"double","minimum":-360,"maximum":360,"default":180},"index$":7},{"name":"latitude","in":"query","description":"Specify the latitude to be used for a radius search. Range: -90 to 90 degrees.","schema":{"type":"number","format":"double","minimum":-90,"maximum":90},"index$":8},{"name":"longitude","in":"query","description":"Specify the longitude to be used for a radius search. Range: -180 to 180 degrees.","schema":{"type":"number","format":"double","minimum":-180,"maximum":180},"index$":9},{"name":"maxradius","in":"query","description":"Limit to events within the specified maximum number of degrees from the geographic point defined by latitude and longitude parameters. Range: 0 to 180 degrees. Mutually exclusive with maxradiuskm.","schema":{"type":"number","format":"double","minimum":0,"maximum":180,"default":180},"index$":10},{"name":"maxradiuskm","in":"query","description":"Limit to events within the specified maximum number of kilometers from the geographic point defined by latitude and longitude parameters. Range: 0 to 20001.6 km. Mutually exclusive with maxradius.","schema":{"type":"number","format":"double","minimum":0,"maximum":20001.6,"default":20001.6},"index$":11},{"name":"minmagnitude","in":"query","description":"Limit to events with a magnitude larger than the specified minimum.","schema":{"type":"number","format":"double"},"example":5,"index$":12},{"name":"maxmagnitude","in":"query","description":"Limit to events with a magnitude smaller than the specified maximum.","schema":{"type":"number","format":"double"},"index$":13},{"name":"magnitudetype","in":"query","description":"Specify a magnitude type to use for testing the minimum and maximum limits.","schema":{"type":"string"},"index$":14},{"name":"mindepth","in":"query","description":"Limit to events with depth more than the specified minimum. Range: -100 to 1000 km.","schema":{"type":"number","format":"double","minimum":-100,"maximum":1000,"default":-100},"index$":15},{"name":"maxdepth","in":"query","description":"Limit to events with depth less than the specified maximum. Range: -100 to 1000 km.","schema":{"type":"number","format":"double","minimum":-100,"maximum":1000,"default":1000},"index$":16},{"name":"catalog","in":"query","description":"Limit to events from a specified catalog. Use the Catalogs Method to find available catalogs.","schema":{"type":"string"},"index$":17},{"name":"contributor","in":"query","description":"Limit to events contributed by a specified contributor. Use the Contributors Method to find available contributors.","schema":{"type":"string"},"index$":18},{"name":"eventid","in":"query","description":"Select a specific event by ID. Event identifiers are data center specific.","schema":{"type":"string"},"index$":19},{"name":"includeallmagnitudes","in":"query","description":"Specify if all magnitudes for the event should be included.","schema":{"type":"boolean","default":false},"index$":20},{"name":"includeallorigins","in":"query","description":"Specify if all origins for the event should be included.","schema":{"type":"boolean","default":false},"index$":21},{"name":"includearrivals","in":"query","description":"Specify if phase arrivals should be included.","schema":{"type":"boolean","default":false},"index$":22},{"name":"includedeleted","in":"query","description":"Specify if deleted products and events should be included. Value 'only' returns only deleted events.","schema":{"type":"string","enum":["false","true","only"],"default":"false"},"index$":23},{"name":"includesuperseded","in":"query","description":"Specify if superseded products should be included. Only works when specifying eventid parameter.","schema":{"type":"boolean","default":false},"index$":24},{"name":"limit","in":"query","description":"Limit the results to the specified number of events. Maximum: 20000.","schema":{"type":"integer","minimum":1,"maximum":20000},"index$":25},{"name":"offset","in":"query","description":"Return results starting at the event count specified, starting at 1.","schema":{"type":"integer","minimum":1,"default":1},"index$":26},{"name":"orderby","in":"query","description":"Order the results.","schema":{"type":"string","enum":["time","time-asc","magnitude","magnitude-asc"],"default":"time"},"index$":27},{"name":"alertlevel","in":"query","description":"Limit to events with a specific PAGER alert level.","schema":{"type":"string","enum":["green","yellow","orange","red"]},"index$":28},{"name":"callback","in":"query","description":"Convert GeoJSON output to a JSONP response using this callback. Must be used with format=geojson.","schema":{"type":"string","pattern":"^[A-Za-z0-9\\._]+$"},"index$":29},{"name":"eventtype","in":"query","description":"Limit to events of a specific type.","schema":{"type":"string"},"example":"earthquake","index$":30},{"name":"jsonerror","in":"query","description":"Request JSON(P) formatted output even on API error results. Must be used with format=geojson.","schema":{"type":"boolean","default":false},"index$":31},{"name":"kmlanimated","in":"query","description":"Whether to include timestamp in generated kml, for google earth animation support. Must be used with format=kml.","schema":{"type":"boolean","default":false},"index$":32},{"name":"kmlcolorby","in":"query","description":"How earthquakes are colored in KML. Must be used with format=kml.","schema":{"type":"string","enum":["age","depth"],"default":"age"},"index$":33},{"name":"maxcdi","in":"query","description":"Maximum value for Maximum Community Determined Intensity reported by DYFI. Range: 0 to 12.","schema":{"type":"number","format":"double","minimum":0,"maximum":12},"index$":34},{"name":"mincdi","in":"query","description":"Minimum value for Maximum Community Determined Intensity reported by DYFI.","schema":{"type":"number","format":"double"},"index$":35},{"name":"maxgap","in":"query","description":"Limit to events with no more than this azimuthal gap. Range: 0 to 360 degrees.","schema":{"type":"number","format":"double","minimum":0,"maximum":360},"index$":36},{"name":"mingap","in":"query","description":"Limit to events with no less than this azimuthal gap. Range: 0 to 360 degrees.","schema":{"type":"number","format":"double","minimum":0,"maximum":360},"index$":37},{"name":"maxmmi","in":"query","description":"Maximum value for Maximum Modified Mercalli Intensity reported by ShakeMap. Range: 0 to 12.","schema":{"type":"number","format":"double","minimum":0,"maximum":12},"index$":38},{"name":"maxsig","in":"query","description":"Limit to events with no more than this significance.","schema":{"type":"integer"},"index$":39},{"name":"minsig","in":"query","description":"Limit to events with no less than this significance.","schema":{"type":"integer"},"index$":40},{"name":"minfelt","in":"query","description":"Limit to events with this many DYFI responses. Minimum: 1.","schema":{"type":"integer","minimum":1},"index$":41},{"name":"producttype","in":"query","description":"Limit to events that have this type of product associated.","schema":{"type":"string"},"examples":{"momentTensor":{"value":"moment-tensor"},"focalMechanism":{"value":"focal-mechanism"},"shakemap":{"value":"shakemap"}},"index$":42},{"name":"productcode","in":"query","description":"Return the event that is associated with the productcode.","schema":{"type":"string"},"index$":43},{"name":"reviewstatus","in":"query","description":"Limit to events with a specific review status.","schema":{"type":"string","enum":["automatic","reviewed","all"],"default":"all"},"index$":44},{"name":"nodata","in":"query","description":"Define the error code that will be returned when no data is found.","schema":{"type":"integer","enum":[204,404],"default":204},"index$":45}],"securitySource":"unspecified"},"GET /count":{"protocol":"http","operationId":"countEarthquakes","responses":{"200":{"description":"Successful response with event count","content":{"text/plain":{"schema":{"type":"integer","example":42}},"application/json":{"schema":{"type":"object","properties":{"count":{"key$":"count","type":"integer"},"maxAllowed":{"key$":"maxAllowed","type":"integer"}},"index$":0}},"application/xml":{"schema":{"type":"object"}}}},"204":{"description":"No data found"},"400":{"description":"Bad request - invalid parameters"}},"parameters":[{"name":"format","in":"query","description":"Specify the output format","schema":{"type":"string","enum":["text","geojson","xml"],"default":"text"},"index$":0},{"name":"starttime","in":"query","description":"Limit to events on or after the specified start time. ISO8601 Date/Time format. UTC assumed if timezone not specified.","schema":{"type":"string","format":"date-time"},"x-ref":"#/components/parameters/starttime","index$":1},{"name":"endtime","in":"query","description":"Limit to events on or before the specified end time. ISO8601 Date/Time format. UTC assumed if timezone not specified.","schema":{"type":"string","format":"date-time"},"x-ref":"#/components/parameters/endtime","index$":2},{"name":"updatedafter","in":"query","description":"Limit to events updated after the specified time. ISO8601 Date/Time format. UTC assumed if timezone not specified.","schema":{"type":"string","format":"date-time"},"x-ref":"#/components/parameters/updatedafter","index$":3},{"name":"minlatitude","in":"query","description":"Limit to events with a latitude larger than the specified minimum. Range: -90 to 90 degrees.","schema":{"type":"number","format":"double","minimum":-90,"maximum":90,"default":-90},"x-ref":"#/components/parameters/minlatitude","index$":4},{"name":"maxlatitude","in":"query","description":"Limit to events with a latitude smaller than the specified maximum. Range: -90 to 90 degrees.","schema":{"type":"number","format":"double","minimum":-90,"maximum":90,"default":90},"x-ref":"#/components/parameters/maxlatitude","index$":5},{"name":"minlongitude","in":"query","description":"Limit to events with a longitude larger than the specified minimum. Range: -360 to 360 degrees.","schema":{"type":"number","format":"double","minimum":-360,"maximum":360,"default":-180},"x-ref":"#/components/parameters/minlongitude","index$":6},{"name":"maxlongitude","in":"query","description":"Limit to events with a longitude smaller than the specified maximum. Range: -360 to 360 degrees.","schema":{"type":"number","format":"double","minimum":-360,"maximum":360,"default":180},"x-ref":"#/components/parameters/maxlongitude","index$":7},{"name":"latitude","in":"query","description":"Specify the latitude to be used for a radius search. Range: -90 to 90 degrees.","schema":{"type":"number","format":"double","minimum":-90,"maximum":90},"x-ref":"#/components/parameters/latitude","index$":8},{"name":"longitude","in":"query","description":"Specify the longitude to be used for a radius search. Range: -180 to 180 degrees.","schema":{"type":"number","format":"double","minimum":-180,"maximum":180},"x-ref":"#/components/parameters/longitude","index$":9},{"name":"maxradius","in":"query","description":"Limit to events within the specified maximum number of degrees from the geographic point. Range: 0 to 180 degrees.","schema":{"type":"number","format":"double","minimum":0,"maximum":180,"default":180},"x-ref":"#/components/parameters/maxradius","index$":10},{"name":"maxradiuskm","in":"query","description":"Limit to events within the specified maximum number of kilometers from the geographic point. Range: 0 to 20001.6 km.","schema":{"type":"number","format":"double","minimum":0,"maximum":20001.6,"default":20001.6},"x-ref":"#/components/parameters/maxradiuskm","index$":11},{"name":"minmagnitude","in":"query","description":"Limit to events with a magnitude larger than the specified minimum.","schema":{"type":"number","format":"double"},"x-ref":"#/components/parameters/minmagnitude","index$":12},{"name":"maxmagnitude","in":"query","description":"Limit to events with a magnitude smaller than the specified maximum.","schema":{"type":"number","format":"double"},"x-ref":"#/components/parameters/maxmagnitude","index$":13},{"name":"magnitudetype","in":"query","description":"Specify a magnitude type to use for testing the minimum and maximum limits.","schema":{"type":"string"},"x-ref":"#/components/parameters/magnitudetype","index$":14},{"name":"mindepth","in":"query","description":"Limit to events with depth more than the specified minimum. Range: -100 to 1000 km.","schema":{"type":"number","format":"double","minimum":-100,"maximum":1000,"default":-100},"x-ref":"#/components/parameters/mindepth","index$":15},{"name":"maxdepth","in":"query","description":"Limit to events with depth less than the specified maximum. Range: -100 to 1000 km.","schema":{"type":"number","format":"double","minimum":-100,"maximum":1000,"default":1000},"x-ref":"#/components/parameters/maxdepth","index$":16},{"name":"catalog","in":"query","description":"Limit to events from a specified catalog.","schema":{"type":"string"},"x-ref":"#/components/parameters/catalog","index$":17},{"name":"contributor","in":"query","description":"Limit to events contributed by a specified contributor.","schema":{"type":"string"},"x-ref":"#/components/parameters/contributor","index$":18},{"name":"eventtype","in":"query","description":"Limit to events of a specific type.","schema":{"type":"string"},"x-ref":"#/components/parameters/eventtype","index$":19},{"name":"reviewstatus","in":"query","description":"Limit to events with a specific review status.","schema":{"type":"string","enum":["automatic","reviewed","all"],"default":"all"},"x-ref":"#/components/parameters/reviewstatus","index$":20},{"name":"alertlevel","in":"query","description":"Limit to events with a specific PAGER alert level.","schema":{"type":"string","enum":["green","yellow","orange","red"]},"x-ref":"#/components/parameters/alertlevel","index$":21}],"securitySource":"unspecified"}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let earthquake_data_ref01_data = Object.values(setup.data.existing.earthquake_data)[0] as any

    // LIST
    const earthquake_data_ref01_ent = client.EarthquakeData()
    const earthquake_data_ref01_match: any = {}

    const earthquake_data_ref01_list = (await earthquake_data_ref01_ent.list(earthquake_data_ref01_match)).map((e: any) => e.data())


    // LOAD
    const earthquake_data_ref01_match_dt0: any = {}
    earthquake_data_ref01_match_dt0.id = earthquake_data_ref01_data.id
    const earthquake_data_ref01_data_dt0 = (await earthquake_data_ref01_ent.load(earthquake_data_ref01_match_dt0)).data()
    assert(earthquake_data_ref01_data_dt0.id === earthquake_data_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/earthquake_data/EarthquakeDataTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = EarthquakeCatalogSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['earthquake_data01','earthquake_data02','earthquake_data03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'EARTHQUAKE_CATALOG_TEST_EARTHQUAKE_DATA_ENTID': idmap,
    'EARTHQUAKE_CATALOG_TEST_LIVE': 'FALSE',
    'EARTHQUAKE_CATALOG_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['EARTHQUAKE_CATALOG_TEST_EARTHQUAKE_DATA_ENTID']

  const live = 'TRUE' === env.EARTHQUAKE_CATALOG_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['EARTHQUAKE_CATALOG_TEST_EARTHQUAKE_DATA_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new EarthquakeCatalogSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.EARTHQUAKE_CATALOG_TEST_EXPLAIN,
    live,
    transport,
    now: Date.now(),
  }

  return setup
}
  
