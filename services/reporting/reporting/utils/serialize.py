"""Response serialization utilities."""
from typing import Callable
from flask import jsonify


def serialize():
    """Wrap a Flask request handler and serialize its JSON return value."""

    def wrap(fn: Callable):
        def decorator(*args, **kwargs):
            result = fn(*args, **kwargs)

            if result is None:
                return jsonify('{}')

            return jsonify(str(result))

        return decorator

    return wrap
