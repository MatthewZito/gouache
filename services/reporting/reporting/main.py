"""Flask app initialization logic.
"""
from flask import Flask, Response


def create_app():
    """A factory function for creating the main application.

    Returns:
        Flask: A Flask application instance.
    """
    app = Flask(__name__)

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
        return response

    return app
