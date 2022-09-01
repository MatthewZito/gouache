"""This module houses authorization middleware.
"""
import json
import os

from types import SimpleNamespace
from typing import Callable, TypeVar
from functools import wraps
from datetime import datetime
from dateutil.parser import parse

from flask import abort, request
from werkzeug.local import LocalProxy

from reporting.context.context import get_session_ctx

RT = TypeVar('RT')


def authorize(fn: Callable[..., RT]) -> Callable[..., RT]:
    """Authorization middleware.
       When decorating a request handler, this middleware validates
       the user session.

    Args:
        fn (Callable): The authorization-guarded request.

    Returns:
        Response: A normalized response object.
    """

    @wraps(fn)
    def decorator(*args, **kws) -> RT:
        sid = request.cookies.get(os.getenv('GOUACHE_SESSION_KEY', 'gouache_session'))

        if not sid:
            abort(401)

        db = LocalProxy(get_session_ctx)

        session = db.get(sid)  # type: ignore

        if not session:
            abort(401)

        session_json = json.loads(session, object_hook=lambda d: SimpleNamespace(**d))

        ts = parse(session_json.expiry).timestamp()

        present = datetime.now().timestamp()

        if ts < present:
            abort(401)

        return fn(*args, **kws)

    return decorator
