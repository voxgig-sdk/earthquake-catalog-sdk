
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { EarthquakeCatalogSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = EarthquakeCatalogSDK.test()
    equal(testsdk instanceof EarthquakeCatalogSDK, true,
      'EarthquakeCatalogSDK.test() must return a client synchronously')
  })

})
