type NormalizedResponse<T> = Promise<
  NormalizedOkResponse<T> | NormalizedErrorResponse
>

interface NormalizedOkResponse<T> extends GouacheResponse<T> {
  ok: true
  data: T
}

interface NormalizedErrorResponse extends GouacheResponse<null> {
  ok: false
  data: null
}

interface GouacheResponse<T> {
  internal: string | null
  friendly: string | null
  data: T | null
  flags: number | null
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

    return normalize(response, [200])
  }

  async post<T, D>(url = '/', payload?: D): NormalizedResponse<T> {
    const response = await this.request(url, {
      method: 'POST',
      ...(payload ? { body: JSON.stringify(payload) } : {}),
    })

    return normalize(response, [200, 201])
  }

  async patch<T, D>(url = '/', payload: D): NormalizedResponse<T> {
    const response = await this.request(url, {
      method: 'PATCH',
      body: JSON.stringify(payload),
    })

    return normalize(response, [200])
  }
}

async function normalize<T>(
  response: Response,
  successCodes: number[],
): NormalizedResponse<T> {
  try {
    const { data, internal, friendly, flags } = await response.json()

    if (!successCodes.includes(response.status)) {
      return {
        ok: false,
        data: null,
        internal,
        friendly,
        flags,
      }
    }

    return {
      ok: true,
      data,
      internal: null,
      friendly: null,
      flags: null,
    }
  } catch (ex) {
    return {
      ok: false,
      data: null,
      internal: ex instanceof Error ? ex.message : JSON.stringify(ex),
      friendly: 'Something went wrong while processing the request.',
      flags: 0,
    }
  }
}
