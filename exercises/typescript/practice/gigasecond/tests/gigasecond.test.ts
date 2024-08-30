import { describe, it, expect } from 'vitest';
import { after } from '../src/gigasecond';

describe('Gigasecond', () => {
  it('tells the date a gigasecond after start date', () => {
    const start = new Date(Date.UTC(2011, 3, 25, 0, 0, 0));
    const expected = new Date(Date.UTC(2043, 0, 1, 1, 46, 40));
    expect(after(start)).toEqual(expected);
  });

  it('tells the date a gigasecond after start date with hours, minutes, and seconds', () => {
    const start = new Date(Date.UTC(2015, 0, 24, 22, 0, 0));
    const expected = new Date(Date.UTC(2046, 9, 2, 23, 46, 40));
    expect(after(start)).toEqual(expected);
  });

  it('does not mutate the input', () => {
    const start = new Date(Date.UTC(2020, 0, 1, 0, 0, 0));
    after(start);
    expect(start).toEqual(new Date(Date.UTC(2020, 0, 1, 0, 0, 0)));
  });

  it('handles dates before 1970-01-01', () => {
    const start = new Date(Date.UTC(1959, 6, 19, 0, 0, 0));
    const expected = new Date(Date.UTC(1991, 2, 27, 1, 46, 40));
    expect(after(start)).toEqual(expected);
  });
});
