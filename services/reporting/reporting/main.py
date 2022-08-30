"""Flask app initialization logic.
"""
from flask import Flask

from reporting.controllers.reporting_controller import reporting
from reporting.meta.const import E_ROUTE_NOT_FOUND, E_UNAUTHORIZED
from reporting.models.gouache_response import err_response


app = Flask(__name__)


@app.errorhandler(404)
def not_found(ex):
    """Route not found handler.

    Args:
        ex (_type_): @todo

    Returns:
        _type_: @todo
    """
    return err_response(E_ROUTE_NOT_FOUND, str(ex)), 404


@app.errorhandler(401)
def unauthorized(ex):
    """Unauthorized request handler.

    Args:
        ex (_type_): @todo

    Returns:
        _type_: @todo
    """
    return err_response(E_UNAUTHORIZED, str(ex)), 401


app.register_blueprint(reporting, url_prefix='/api')
