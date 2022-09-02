/* eslint-disable @typescript-eslint/no-unsafe-assignment */
type NormalizedResponse<T> = Promise<
  NormalizedErrorResponse | NormalizedOkResponse<T>
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

  cors: boolean

  constructor({
    baseUrl = '/',
    withCredentials = true,
    cors = true,
  }: {
    baseUrl: string
    withCredentials?: boolean
    cors?: boolean
  }) {
    this.baseUrl = baseUrl
    this.withCredentials = withCredentials
    this.cors = cors
  }

  async get<T>(url: string): NormalizedResponse<T> {
    const response = await this.request(url, {
      method: 'GET',
    })

    return normalize(response, [200])
  }

  async post<T, D>(
    url: string,
    payload?: D,
    json = false,
  ): NormalizedResponse<T> {
    const response = await this.request(url, {
      method: 'POST',
      ...(json ? { headers: { 'Content-Type': 'application/json' } } : {}),
      ...(payload ? { body: JSON.stringify(payload) } : {}),
    })

    return normalize(response, [200, 201])
  }

  async patch<T, D>(
    url: string,
    payload: D,
    json = false,
  ): NormalizedResponse<T> {
    const response = await this.request(url, {
      method: 'PATCH',
      ...(json ? { headers: { 'Content-Type': 'application/json' } } : {}),
      body: JSON.stringify(payload),
    })

    return normalize(response, [200])
  }

  async delete<T>(url: string): NormalizedResponse<T> {
    const response = await this.request(url, {
      method: 'DELETE',
    })

    return normalize(response, [200])
  }

  private async request(url: string, opts: RequestInit) {
    return fetch(this.baseUrl + url, {
      mode: this.cors ? 'cors' : 'no-cors',
      ...(this.withCredentials ? { credentials: 'include' } : {}),
      ...opts,
    })
  }
}

async function normalize<T>(
  response: Response,
  successCodes: number[],
): NormalizedResponse<T> {
  try {
    const d = await response.json()

    // @todo validate
    const { data, internal, friendly, flags } = d

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
