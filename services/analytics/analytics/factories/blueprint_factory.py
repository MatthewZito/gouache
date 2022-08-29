def construct_blueprint(name: str):

    myblueprint = Blueprint(name, __name__)

    @myblueprint.route('/route', methods=['GET'])
    def route():
        database = database

    return myblueprint
