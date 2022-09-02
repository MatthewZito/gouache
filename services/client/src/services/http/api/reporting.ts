import type { MutableReport, Report } from '@/types'
import type { UUID } from '@/types/scalar'

import { HttpClient } from '../client'

export interface GetReportsResponse {
  items: Report[]
  next: string
}

const client = new HttpClient({
  baseUrl: `${import.meta.env.VITE_REPORTING_API}/api/report`,
})

export const reportingApi = {
  async getReport(id: UUID) {
    return client.get<Report>(`/${id}`)
  },

  async getReports() {
    return client.get<GetReportsResponse>('')
  },

  async createReport(payload: Omit<MutableReport, 'caller'>) {
    const blob = new Blob([JSON.stringify(payload)], {
      type: 'application/json; charset=UTF-8',
    })

    ;(payload as MutableReport).caller = 'gouache/client'

    window.navigator.sendBeacon(
      `${import.meta.env.VITE_REPORTING_API}/api/report`,
      blob,
    )
  },
}
