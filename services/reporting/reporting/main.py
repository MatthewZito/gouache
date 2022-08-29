from reporting.controllers.reporting_controller import reporting
from flask import Flask
from reporting.meta.const import E_ROUTE_NOT_FOUND, E_UNAUTHORIZED
from reporting.models.gouache_response import err_response

app = Flask(__name__)


@app.errorhandler(404)
def not_found(e):
    return err_response(E_ROUTE_NOT_FOUND, str(e)), 404


@app.errorhandler(401)
def unauthorized(e):
    return err_response(E_UNAUTHORIZED, str(e)), 401


app.register_blueprint(reporting, url_prefix='/api')
