import { defineConfig } from 'vitest/config'

export default defineConfig({
  test: {
    include: ['practice/**/*.test.ts'],
    globals: true,
  },
})