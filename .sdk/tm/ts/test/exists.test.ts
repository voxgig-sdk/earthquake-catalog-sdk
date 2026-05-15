
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { EarthquakeCatalogSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await EarthquakeCatalogSDK.test()
    equal(null !== testsdk, true)
  })

})
