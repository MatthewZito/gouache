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
  baseUrl: string
  withCredentials: boolean
  constructor({
    baseUrl = '/',
    withCredentials = false,
  }: {
    baseUrl: string
    withCredentials: boolean
  }) {
    this.baseUrl = baseUrl
    this.withCredentials = withCredentials
  }

  private request(url = '/', opts: RequestInit) {
    return fetch(this.baseUrl + url, {
      mode: 'cors',
      ...(this.withCredentials ? { credentials: 'include' } : {}),
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

  async post<T, D>(url = '/', payload?: D): NormalizedResponse<T> {
    const response = await this.request(url, {
      method: 'POST',
      ...(payload ? { body: JSON.stringify(payload) } : {}),
    })

    if (![200, 201].includes(response.status)) {
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
