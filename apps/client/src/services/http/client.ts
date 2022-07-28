type NormalizedResponse<T> = Promise<
  NormalizedOkResponse<T> | NormalizedErrorResponse
>

interface NormalizedOkResponse<T> {
  ok: true
  data: T
  error: null
}

interface NormalizedErrorResponse {
  ok: false
  data: null
  error: string
}

export class HttpClient {
  constructor(public baseUrl: string) {}

  private request(url = '/', opts: RequestInit) {
    return fetch(this.baseUrl + url, {
      mode: 'cors',
      ...opts,
    })
  }

  async get<T>(url = '/'): NormalizedResponse<T> {
    const response = await this.request(url, {
      method: 'GET',
    })

    if (response.status !== 200) {
      return {
        ok: false,
        data: null,
        error: 'Something went wrong while making the request.',
      }
    }

    const { ok, data, error } = await response.json()

    return {
      ok,
      data,
      error,
    }
  }

  async post<T, D>(url = '/', payload: D): NormalizedResponse<T> {
    const response = await this.request(url, {
      method: 'POST',
      body: JSON.stringify(payload),
    })

    if (response.status !== 200) {
      return {
        ok: false,
        data: null,
        error: 'Something went wrong while making the request.',
      }
    }

    const { ok, data, error } = await response.json()

    return {
      ok,
      data,
      error,
    }
  }

  async patch<T, D>(url = '/', payload: D): NormalizedResponse<T> {
    const response = await this.request(url, {
      method: 'PATCH',
      body: JSON.stringify(payload),
    })

    if (response.status !== 200) {
      return {
        ok: false,
        data: null,
        error: 'Something went wrong while making the request.',
      }
    }

    const { ok, data, error } = await response.json()

    return {
      ok,
      data,
      error,
    }
  }
}
