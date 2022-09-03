import { normalizeNullish } from '../normalize'

describe('normalization utilities', () => {
  describe('normalizeNullish', () => {
    it('normalizes nullish values', () => {
      const nullishValues = [null, void 0]

      nullishValues.forEach(nullishValue => {
        expect(normalizeNullish(nullishValue)).toEqual('N/A')
      })
    })

    it('returns non-nullish values as-is', () => {
      const nonNullishValues = [-1, '', [1], 'v']

      nonNullishValues.forEach(nonNullishValue => {
        expect(normalizeNullish(nonNullishValue)).toEqual(nonNullishValue)
      })
    })
  })
})
