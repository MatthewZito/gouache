from analytics.controllers.reporting import reporting
from flask import Flask
from analytics.models.gouache_response import err_response

app = Flask(__name__)


@app.errorhandler(404)
def page_not_found(e):
    return err_response('Route not found', str(e)), 404


@app.errorhandler(401)
def page_not_found(e):
    return err_response('You must be authorized to access this resource', str(e)), 401


app.register_blueprint(reporting, url_prefix='/api')
