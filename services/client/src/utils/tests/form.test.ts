import {
  generateValidator,
  listRequired,
  maxLength,
  minLength,
  optionalWithRule,
  required,
} from '../form'

describe('form validation utilities', () => {
  describe('required', () => {
    it('requires a valid string input with a length greater than zero', () => {
      const output = 'test'
      const validate = required(output)

      const inputs = [null, '']

      inputs.forEach(input => {
        expect(validate(input)).toEqual(output)
      })

      expect(validate('input')).toEqual(true)
    })
  })

  describe('listRequired', () => {
    it('requires a valid list with a length greater than zero', () => {
      const output = 'test'
      const validate = listRequired(output)

      const inputs = [null, []]

      inputs.forEach(input => {
        expect(validate(input)).toEqual(output)
      })

      expect(validate(['input'])).toEqual(true)
    })
  })

  describe('minLength', () => {
    it('enforces a minimum length', () => {
      const output = 'test'
      const validate = minLength(output, 3)

      const inputs = [null, '', '12']

      inputs.forEach(input => {
        expect(validate(input)).toEqual(output)
      })

      expect(validate('1234')).toEqual(true)
    })
  })

  describe('maxLength', () => {
    it('enforces a maximum length', () => {
      const output = 'test'
      const validate = maxLength(output, 10)

      const inputs = ['12345678901']

      inputs.forEach(input => {
        expect(validate(input)).toEqual(output)
      })

      expect(validate('1234567890')).toEqual(true)
    })
  })

  describe('optionalWithRule', () => {
    it('passes through the given rule if provided no value', () => {
      const output = 'test'
      const validate = optionalWithRule(minLength(output, 3))

      expect(validate(null)).toEqual(true)
      expect(validate('')).toEqual(output)
      expect(validate('123')).toEqual(true)
    })
  })

  describe('generateValidator', () => {
    it('generates a validator that validates a given input against all provided rules', () => {
      const inputs = [null, '', '12', '123456']

      const validate = generateValidator([
        required(''),
        minLength('', 3),
        maxLength('', 5),
      ])

      inputs.forEach(input => {
        expect(validate(input)).toEqual(false)
      })

      expect(validate('1234')).toEqual(true)
    })
  })
})
