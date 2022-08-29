from datetime import datetime
from functools import wraps
import json
import os
from types import SimpleNamespace
from flask import abort, request
from werkzeug.local import LocalProxy

from reporting.context.context import get_session_ctx


def authorize(f):
    @wraps(f)
    def decorator(*args, **kws):
        sid = request.cookies.get(os.getenv('GOUACHE_SESSION_KEY', 'gouache_session'))

        if not sid:
            abort(401)

        db = LocalProxy(get_session_ctx)

        session = db.get(sid)

        if not session:
            abort(401)

        session_json = json.loads(session, object_hook=lambda d: SimpleNamespace(**d))

        ts = datetime.fromisoformat(session_json.Expiry).timestamp()

        present = datetime.now().timestamp()

        if ts < present:
            abort(401)

        return f(*args, **kws)

    return decorator
