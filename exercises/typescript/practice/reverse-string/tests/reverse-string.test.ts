import { expect, test } from 'vitest'
import { reverse } from '../src/reverse-string'

test('empty string', () => {
  expect(reverse('')).toBe('')
})

test('word', () => {
  expect(reverse('robot')).toBe('tobor')
})

test('capitalized word', () => {
  expect(reverse('Ramen')).toBe('nemaR')
})

test('sentence with punctuation', () => {
  expect(reverse('I\'m hungry!')).toBe('!yrgnuh m\'I')
})

test('palindrome', () => {
  expect(reverse('racecar')).toBe('racecar')
})