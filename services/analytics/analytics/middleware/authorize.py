from datetime import datetime
from functools import wraps
import json
import os
from types import SimpleNamespace
from flask import abort, request

from analytics.context.context import get_session_ctx


def authorize(f):
    @wraps(f)
    def decorator(*args, **kws):
        sid = request.cookies.get(os.getenv('GOUACHE_SESSION_KEY', 'gouache_session'))

        if not sid:
            abort(401)

        redis = get_session_ctx()
        session = redis.get(sid)

        if not session:
            abort(401)

        session_json = json.loads(session, object_hook=lambda d: SimpleNamespace(**d))

        ts = datetime.utcfromtimestamp(session_json.Expiry)

        present = datetime.now()

        if ts < present.date():
            abort(401)

        return f(*args, **kws)

    return decorator
