import { describe, it, expect } from 'vitest';
import { anagrams } from '../src/anagram';

describe('Anagram', () => {
  it('no matches', () => {
    expect(anagrams('diaper', ['hello', 'world', 'zombies', 'pants'])).toEqual(
      []
    );
  });

  it('detects two anagrams', () => {
    expect(anagrams('master', ['stream', 'pigeon', 'maters'])).toEqual([
      'stream',
      'maters',
    ]);
  });

  it('does not detect anagram subsets', () => {
    expect(anagrams('good', ['dog', 'goody'])).toEqual([]);
  });

  it('detects anagram', () => {
    expect(
      anagrams('listen', ['enlists', 'google', 'inlets', 'banana'])
    ).toEqual(['inlets']);
  });

  it('detects three anagrams', () => {
    expect(
      anagrams('allergy', [
        'gallery',
        'ballerina',
        'regally',
        'clergy',
        'largely',
        'leading',
      ])
    ).toEqual(['gallery', 'regally', 'largely']);
  });

  it('detects multiple anagrams with different case', () => {
    expect(anagrams('nose', ['Eons', 'ONES'])).toEqual(['Eons', 'ONES']);
  });

  it('does not detect non-anagrams with identical checksum', () => {
    expect(anagrams('mass', ['last'])).toEqual([]);
  });

  it('detects anagrams case-insensitively', () => {
    expect(
      anagrams('Orchestra', ['cashregister', 'Carthorse', 'radishes'])
    ).toEqual(['Carthorse']);
  });

  it('detects anagrams using case-insensitive subject', () => {
    expect(
      anagrams('Orchestra', ['cashregister', 'carthorse', 'radishes'])
    ).toEqual(['carthorse']);
  });

  it('detects anagrams using case-insensitive possible matches', () => {
    expect(
      anagrams('orchestra', ['cashregister', 'Carthorse', 'radishes'])
    ).toEqual(['Carthorse']);
  });

  it('does not detect an anagram if the original word is repeated', () => {
    expect(anagrams('go', ['go Go GO'])).toEqual([]);
  });

  it('anagrams must use all letters exactly once', () => {
    expect(anagrams('tapper', ['patter'])).toEqual([]);
  });

  it('words are not anagrams of themselves (case-insensitive)', () => {
    expect(anagrams('BANANA', ['BANANA', 'Banana', 'banana'])).toEqual([]);
  });

  it('words other than themselves can be anagrams', () => {
    expect(anagrams('LISTEN', ['Listen', 'Silent', 'LISTEN'])).toEqual([
      'Silent',
    ]);
  });
});
