"""Application-wide error handlers"""

from typing import Tuple

from flask import Blueprint, Response

from reporting.meta.const import (
    E_INTERNAL_SERVER_ERROR,
    E_METHOD_NOT_ALLOWED,
    E_ROUTE_NOT_FOUND,
    E_UNAUTHORIZED,
)
from reporting.models.gouache_response import err_response

error_handlers = Blueprint(
    'error_handlers',
    __name__,
)


@error_handlers.app_errorhandler(401)
def unauthorized(ex: Exception) -> Tuple[Response, int]:
    """Unauthorized request handler.

    Args:
        ex (_type_): @todo

    Returns:
        Tuple[Response, int]: A normalized response object; an HTTP status code.
    """
    return err_response(E_UNAUTHORIZED, str(ex)), 401


@error_handlers.app_errorhandler(404)
def not_found(ex: Exception) -> Tuple[Response, int]:
    """Route not found handler.

    Args:
        ex (_type_): @todo

    Returns:
        Tuple[Response, int]: A normalized response object; an HTTP status code.
    """
    return err_response(E_ROUTE_NOT_FOUND, str(ex)), 404


@error_handlers.app_errorhandler(405)
def method_not_allowed(ex: Exception) -> Tuple[Response, int]:
    """Method not allowed handler.

    Args:
        ex (_type_): @todo

    Returns:
        Tuple[Response, int]: A normalized response object; an HTTP status code.
    """
    return err_response(E_METHOD_NOT_ALLOWED, str(ex)), 405


@error_handlers.app_errorhandler(500)
def server_error(ex: Exception) -> Tuple[Response, int]:
    """Server error handler.

    Args:
        ex (_type_): @todo

    Returns:
        Tuple[Response, int]: A normalized response object; an HTTP status code.
    """
    return err_response(E_INTERNAL_SERVER_ERROR, str(ex)), 500
