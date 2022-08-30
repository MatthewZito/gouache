"""Response formatters and normalizers."""
from flask import Response, jsonify


def gouache_response(
    data: object = None, friendly='', internal='', flags=0
) -> Response:
    """A system-wide standard response format.

    Args:
        data (object, optional): The response data. Defaults to None.
        friendly (str, optional): The user-facing error message. Defaults to ''.
        internal (str, optional): The internal error message. Defaults to ''.
        flags (int, optional): A bit-field of metadata flags. Defaults to 0.

    Returns:
        Response: A JSON response object.
    """
    return jsonify(
        {'data': data, 'friendly': friendly, 'internal': internal, 'flags': flags}
    )


def ok_response(data: object) -> Response:
    """Build a successful gouache_response.

    Args:
        data (object): The response data.

    Returns:
        Response: A JSON response object.
    """
    return gouache_response(data)


def err_response(friendly: str, internal: str, flags=0) -> Response:
    """Build an erroneous gouache_response.

    Args:
        friendly (str): The user-facing error message.
        internal (str): The internal error message.
        flags (int, optional): A bit-field of metadata flags. Defaults to 0.

    Returns:
        Response: A JSON response object.
    """
    return gouache_response(None, friendly, internal, flags)
