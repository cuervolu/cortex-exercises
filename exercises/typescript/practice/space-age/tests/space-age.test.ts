import { describe, it, expect } from 'vitest';
import { SpaceAge } from '../src/space-age';

function assertInDelta(actual: number, expected: number) {
  const delta = 0.01;
  expect(Math.abs(actual - expected)).toBeLessThanOrEqual(delta);
}

describe('Space Age', () => {
  it('age on Earth', () => {
    const age = new SpaceAge(1000000000);
    assertInDelta(age.onEarth(), 31.69);
  });

  it('age on Mercury', () => {
    const age = new SpaceAge(2134835688);
    assertInDelta(age.onMercury(), 280.88);
  });

  it('age on Venus', () => {
    const age = new SpaceAge(189839836);
    assertInDelta(age.onVenus(), 9.78);
  });

  it('age on Mars', () => {
    const age = new SpaceAge(2329871239);
    assertInDelta(age.onMars(), 39.25);
  });

  it('age on Jupiter', () => {
    const age = new SpaceAge(901876382);
    assertInDelta(age.onJupiter(), 2.41);
  });

  it('age on Saturn', () => {
    const age = new SpaceAge(3000000000);
    assertInDelta(age.onSaturn(), 3.23);
  });

  it('age on Uranus', () => {
    const age = new SpaceAge(3210123456);
    assertInDelta(age.onUranus(), 1.21);
  });

  it('age on Neptune', () => {
    const age = new SpaceAge(8210123456);
    assertInDelta(age.onNeptune(), 1.58);
  });
});