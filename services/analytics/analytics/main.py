from analytics.controllers.reporting import reporting
from flask import Flask

app = Flask(__name__)


app.register_blueprint(reporting, url_prefix='/api')
