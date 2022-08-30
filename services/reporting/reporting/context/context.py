"""Flask global context getters.
"""
from flask import g
from reporting.repositories.report_repository import ReportRepository
from reporting.repositories.session_repository import SessionRepository


def get_report_ctx() -> ReportRepository:
    """A getter method for global report context.

    Returns:
        ReportRepository: An initialized report repository.
    """
    if 'report' not in g:
        g.report = ReportRepository("report")

    return g.report


def get_session_ctx() -> SessionRepository:
    """A getter method for global session context.

    Returns:
        SessionRepository: An initialized session repository.
    """
    if 'session' not in g:
        g.session = SessionRepository()

    return g.session
