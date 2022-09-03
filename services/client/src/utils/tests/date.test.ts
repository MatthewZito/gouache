import {
  epochToReadableTime,
  nowPlusNSeconds,
  secondsToMinutes,
  toReadableDate,
} from '../date'

describe('date utilities', () => {
  beforeAll(() => {
    vi.useFakeTimers({
      now: new Date('2022, 1, 1'),
    })
  })
  describe('toReadableDate', () => {
    it('formats a given timestamp into a readable date', () => {
      const input = '2022-09-03T11:54:05.392Z'

      expect(toReadableDate(input)).toEqual('09/03/2022')
    })

    it('formats a given date into a readable date', () => {
      const input = new Date('2022-09-03T11:54:05.392Z')

      expect(toReadableDate(input)).toEqual('09/03/2022')
    })

    it('no-ops on bad input', () => {
      const inputs = ['', void 0]

      inputs.forEach(input => {
        expect(toReadableDate(input)).toEqual(null)
      })
    })
  })

  describe('epochToReadableTime', () => {
    it('converts a UNIX epoch into a readable time', () => {
      const input = 1662206370608

      expect(epochToReadableTime(input)).toEqual('10:50 AM')
    })

    it('no-ops on bad input', () => {
      const inputs = [0, null]

      inputs.forEach(input => {
        expect(epochToReadableTime(input)).toEqual(null)
      })
    })
  })

  describe('secondsToMinutes', () => {
    it('converts seconds to minutes', () => {
      const input = 60000

      expect(secondsToMinutes(input)).toEqual('1000.0')
    })

    it('no-ops on bad input', () => {
      const inputs = [0, null]

      inputs.forEach(input => {
        expect(secondsToMinutes(input)).toEqual(null)
      })
    })
  })

  describe('nowPlusNSeconds', () => {
    it('calculates "now" plus n seconds', () => {
      const nSeconds = 3600

      expect(nowPlusNSeconds(nSeconds)).toEqual('01:00 AM')
    })

    it('no-ops on bad input', () => {
      const inputs = [0, null]

      inputs.forEach(input => {
        expect(nowPlusNSeconds(input)).toEqual(null)
      })
    })
  })
})
