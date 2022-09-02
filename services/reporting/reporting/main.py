"""Flask app initialization logic.
"""
import os
import socket
from typing import Tuple
from flask import Flask, Response
from flask_cors import CORS

from reporting.services.message_queue_service import (
    MessageQueueService,
)
from reporting.models.gouache_response import ok_response


def do_side_effects(app: Flask):
    """Perform startup side effects, such as initializing adjacent services
    that do not necessarily interact directly with the Flask app instance.

    Args:
        app (Flask): The Flask app instance.
    """
    queue_service = MessageQueueService(
        os.getenv('SQS_QUEUE_NAME', 'report-queue'), app
    )
    queue_service.init()


def create_app():
    """A factory function for creating the main application.

    Returns:
        Flask: A Flask application instance.
    """
    app = Flask(__name__)
    CORS(
        app,
        resources={
            r"/api/*": {
                "origins": "http://localhost:3000",
                "allow_headers": "*",
                "supports_credentials": True,
            }
        },
    )

    from reporting.controllers.global_controller import error_handlers

    app.register_blueprint(error_handlers)

    from reporting.controllers.reporting_controller import reporting

    app.register_blueprint(reporting, url_prefix='/api')

    @app.after_request
    def apply_headers(response: Response) -> Response:
        """Application-wide hook for setting headers on all out-bound responses.

        Args:
            response (Response): Outbound response.

        Returns:
            Response: The forwarded outbound response.
        """
        response.headers["X-Powered-By"] = "gouache/reporting"
        response.headers["Content-Type"] = "application/json"

        return response

    @app.route('/health', methods=['GET'])
    def health() -> Tuple[Response, int]:
        return (ok_response({'server': socket.gethostname()}), 200)

    do_side_effects(app)
    return app
