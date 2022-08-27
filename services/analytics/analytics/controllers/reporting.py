from analytics.context.context import get_report_ctx
from analytics.repositories.report_repository import ReportRepository
from analytics.middleware.authorize import authorize
from analytics.models.gouache_response import err_response, ok_response
from analytics.utils.deserialize import deserialize
from flask import Blueprint
from analytics.entities.report import Report


reporting = Blueprint(
    'reporting',
    __name__,
)


@reporting.route('/report/<id>', methods=['GET'])
@authorize
def get_report(id: str):
    result = get_report_ctx().get(id)

    if type(result) == str:
        return (
            err_response('An exception occurred while retrieving the report', result),
            400,
        )

    if 'Item' in result:
        return ok_response(result['Item']), 201

    return ok_response(None), 404


@reporting.route('/report', methods=['POST'])
@deserialize(Report)
@authorize
def create_report(report: Report):

    result = get_report_ctx().put(
        caller=report.caller,
        data=report.data,
        id=report.id,
        name=report.name,
    )

    if (
        'ResponseMetadata' in result
        and result['ResponseMetadata']['HTTPStatusCode'] == 200
    ):
        return ok_response(None), 201

    return err_response('An exception occurred while creating the report', result), 400
