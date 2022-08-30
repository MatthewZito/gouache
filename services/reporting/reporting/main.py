"""Flask app initialization logic.
"""
from typing import Tuple

from flask import Flask, Response

from reporting.controllers.reporting_controller import reporting
from reporting.meta.const import E_ROUTE_NOT_FOUND, E_UNAUTHORIZED
from reporting.models.gouache_response import err_response


app = Flask(__name__)


@app.errorhandler(404)
def not_found(ex) -> Tuple[Response, int]:
    """Route not found handler.

    Args:
        ex (_type_): @todo

    Returns:
        Tuple[Response, int]: A normalized response object; an HTTP status code.
    """
    return err_response(E_ROUTE_NOT_FOUND, str(ex)), 404


@app.errorhandler(401)
def unauthorized(ex) -> Tuple[Response, int]:
    """Unauthorized request handler.

    Args:
        ex (_type_): @todo

    Returns:
        Tuple[Response, int]: A normalized response object; an HTTP status code.
    """
    return err_response(E_UNAUTHORIZED, str(ex)), 401


app.register_blueprint(reporting, url_prefix='/api')
