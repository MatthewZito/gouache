from flask import request


def deserialize(class_):
    def wrap(f):
        def decorator(*args):
            try:
                obj = class_(**request.get_json())
                return f(obj)
            except Exception as e:
                return f(None)

        return decorator

    return wrap
