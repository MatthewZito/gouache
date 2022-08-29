from flask import g
from reporting.repositories.report_repository import ReportRepository
from reporting.repositories.session_repository import SessionRepository


def get_report_ctx():
    if 'report' not in g:
        g.report = ReportRepository("report")

    return g.report


def get_session_ctx():
    if 'session' not in g:
        g.session = SessionRepository()

    return g.session
