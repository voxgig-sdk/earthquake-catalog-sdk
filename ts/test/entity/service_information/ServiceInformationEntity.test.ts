

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


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
loadEnvLocal(__dirname + '/../../../.env.local')


describe('ServiceInformationEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when EARTHQUAKE_CATALOG_TEST_LIVE=TRUE.
  afterEach(liveDelay('EARTHQUAKE_CATALOG_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = EarthquakeCatalogSDK.test()
    const ent = testsdk.ServiceInformation()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.EARTHQUAKE_CATALOG_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'service_information.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[],"name":"service_information","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{},"contract":{"id":"GET /catalogs","json":"{\"operationId\":\"getCatalogs\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"type\":\"string\"},\"type\":\"array\"}}},\"description\":\"List of available catalogs\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/catalogs","segments":[{"lit":"catalogs"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0},{"active":true,"args":{},"contract":{"id":"GET /contributors","json":"{\"operationId\":\"getContributors\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"type\":\"string\"},\"type\":\"array\"}}},\"description\":\"List of available contributors\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/contributors","segments":[{"lit":"contributors"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"active":true,"args":{},"contract":{"id":"GET /application.json","json":"{\"operationId\":\"getApplicationJson\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"type\":\"object\"}}},\"description\":\"Enumerated parameter values\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/application.json","segments":[{"lit":"application.json"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0},{"active":true,"args":{},"contract":{"id":"GET /application.wadl","json":"{\"operationId\":\"getWadl\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/xml\":{\"schema\":{\"type\":\"object\"}}},\"description\":\"WADL document\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/application.wadl","segments":[{"lit":"application.wadl"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":1},{"active":true,"args":{},"contract":{"id":"GET /version","json":"{\"operationId\":\"getVersion\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"text/plain\":{\"schema\":{\"example\":\"1.13.6\",\"type\":\"string\"}}},\"description\":\"Service version number\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/version","segments":[{"lit":"version"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":2}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"service_information","name__orig":"service_information","Name":"ServiceInformation","name_":"service_information","name-":"service-information","NAME":"SERVICE_INFORMATION","index$":1}, {"active":true,"entity":"service_information","key$":"BasicServiceInformationFlow","kind":"basic","name":"BasicServiceInformationFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"service_information_ref01"}}],"index$":0},{"active":true,"data":{},"input":{"ref":"service_information_ref01","srcdatavar":"service_information_ref01_data","suffix":"_dt0"},"match":{},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-service_information_ref01"}}],"index$":1}]}, 'ServiceInformation')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let service_information_ref01_data = Object.values(setup.data.existing.service_information)[0] as any

    // LIST
    const service_information_ref01_ent = client.ServiceInformation()
    const service_information_ref01_match: any = {}

    const service_information_ref01_list = (await service_information_ref01_ent.list(service_information_ref01_match)).map((e: any) => e.data())


    // LOAD
    const service_information_ref01_match_dt0: any = {}
    const service_information_ref01_data_dt0 = (await service_information_ref01_ent.load(service_information_ref01_match_dt0)).data()
    assert(null != service_information_ref01_data_dt0)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/service_information/ServiceInformationTestData.json')

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
    ['service_information01','service_information02','service_information03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'EARTHQUAKE_CATALOG_TEST_SERVICE_INFORMATION_ENTID': idmap,
    'EARTHQUAKE_CATALOG_TEST_LIVE': 'FALSE',
    'EARTHQUAKE_CATALOG_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['EARTHQUAKE_CATALOG_TEST_SERVICE_INFORMATION_ENTID']

  const live = 'TRUE' === env.EARTHQUAKE_CATALOG_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['EARTHQUAKE_CATALOG_TEST_SERVICE_INFORMATION_ENTID']
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
  
