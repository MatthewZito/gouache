from flask import jsonify


def serialize():
    def wrap(f):
        def dd(*args, **kwargs):
            result = f(*args, **kwargs)

            if result is None:
                return jsonify('{}')

            return jsonify(str(result))

        return dd

    return wrap
