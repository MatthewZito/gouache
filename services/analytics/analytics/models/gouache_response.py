from flask import Response, jsonify


def gouache_response(data=None, friendly='', internal='', flags=0) -> Response:
    return jsonify(
        {'data': data, 'friendly': friendly, 'internal': internal, 'flags': flags}
    )


def ok_response(data: object) -> Response:
    return gouache_response(data)


def err_response(friendly: str, internal: str, flags=0) -> Response:
    return gouache_response(None, friendly, internal, flags)
