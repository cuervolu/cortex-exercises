import { defineConfig } from 'vitest/config'

export default defineConfig({
  test: {
    include: ['practice/**/test/*.test.ts', 'practice/**/tests/*.test.ts'],
    globals: true,
  },
})