"""Request deserialization utilities."""
from typing import Type, Callable
from flask import request


def deserialize(class_: Type):
    """Wrap a Flask request handler and deserialize its
    JSON request body into class `class_`.

    Args:
        class_ (Type[T]): A class constructor into which the
        request body will be deserialized.
    """

    def wrap(fn: Callable):
        def decorator(*args):  # pylint: disable=W0613
            try:

                obj = class_(**request.get_json())
                return fn(obj)
            except Exception:
                return fn(None)

        return decorator

    return wrap
