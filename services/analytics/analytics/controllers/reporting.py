from analytics.utils.serialize import serialize
from analytics.utils.deserialize import deserialize
from flask import Blueprint
from analytics.entities.report import Report

reports: list[Report] = []

reporting = Blueprint(
    'reporting',
    __name__,
)


@reporting.route('/report/<id>')
@serialize()
def get_report(id: str):
    print(int(id))
    report = next((r for r in reports if r.id == int(id)), None)
    # report = next(r for r in reports if r.id == id)
    print(f"here {report}")
    return str(report)


@reporting.route('/report', methods=['POST'])
@deserialize(Report)
def create_report(report):

    print(report)
    reports.append(report)

    return f'ok'
