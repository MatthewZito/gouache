"""This module houses authorization middleware.
"""
import json
import os

from datetime import datetime
from functools import wraps
from types import SimpleNamespace
from typing import Callable

from flask import abort, request
from werkzeug.local import LocalProxy

from reporting.context.context import get_session_ctx


def authorize(fn: Callable):
    """Authorization middleware.
       When decorating a request handler, this middleware validates
       the user session.

    Args:
        fn (Callable): The authorization-guarded request.

    Returns:
        _type_: @todo
    """

    @wraps(fn)
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

        return fn(*args, **kws)

    return decorator
