from flask import appcontext_pushed
from contextlib import contextmanager


@contextmanager
def test_context_set(
    app, session_repository: BaseRepository, report_repository: BaseRepository
):
    def handler(sender, **kwargs):
        g.session = session_repository
        g.report = report_repository

    with appcontext_pushed.connected_to(handler, app):
        yield
