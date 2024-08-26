import { expect, test } from 'vitest'
import { hello } from '../src/hello-world'

test('says hello world', () => {
  expect(hello()).toBe('Hello, World!')
})