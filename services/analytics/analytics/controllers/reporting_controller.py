from analytics.context.context import get_report_ctx
from analytics.repositories.report_repository import ReportRepository
from analytics.middleware.authorize import authorize
from analytics.models.gouache_response import err_response, ok_response
from analytics.utils.deserialize import deserialize
from flask import Blueprint
from analytics.entities.report import Report

from werkzeug.local import LocalProxy


reporting = Blueprint(
    'reporting',
    __name__,
)


@reporting.route('/report/<id>', methods=['GET'])
@authorize
def get_report(id: str):
    db = LocalProxy(get_report_ctx)
    result = db.get(id)

    if type(result) == str:
        return (
            err_response('An exception occurred while retrieving the report', result),
            400,
        )

    # {'Item': {'Data': 'some data', 'Id': 'eddde4d5-90b7-4cc3-91bb-b4561b4137b4', 'Caller': 'gouache-test', 'Name': 'the report name', 'TS': '1661786507886.366'}, 'ResponseMetadata': {'RequestId': 'bda454b7-8c8d-4abc-ab04-a650f1095196', 'HTTPStatusCode': 200, 'HTTPHeaders': {'date': 'Mon, 29 Aug 2022 15:22:02 GMT', 'content-type': 'application/x-amz-json-1.0', 'x-amz-crc32': '198904004', 'x-amzn-requestid': 'bda454b7-8c8d-4abc-ab04-a650f1095196', 'content-length': '177', 'server': 'Jetty(9.4.43.v20210629)'}, 'RetryAttempts': 0}}
    if 'Item' in result:
        # {'Data': 'some data', 'Id': 'eddde4d5-90b7-4cc3-91bb-b4561b4137b4', 'Caller': 'gouache-test', 'Name': 'the report name', 'TS': '1661786507886.366'}
        return ok_response(result['Item']), 200

    return ok_response(None), 404


@reporting.route('/report', methods=['POST'])
@deserialize(Report)
@authorize
def create_report(report: Report):
    db = LocalProxy(get_report_ctx)

    result = db.put(
        caller=report.caller,
        data=report.data,
        id=report.id,
        name=report.name,
    )

    if type(result) == str:
        return (
            err_response('An exception occurred while creating the report', result),
            400,
        )

    if (
        'ResponseMetadata' in result
        and result['ResponseMetadata']['HTTPStatusCode'] == 200
    ):
        # {'ResponseMetadata': {'RequestId': '3e5e2b4b-da84-4b48-aade-baeab12c20e9', 'HTTPStatusCode': 200, 'HTTPHeaders': {'date': 'Mon, 29 Aug 2022 15:21:20 GMT', 'content-type': 'application/x-amz-json-1.0', 'x-amz-crc32': '2745614147', 'x-amzn-requestid': '3e5e2b4b-da84-4b48-aade-baeab12c20e9', 'content-length': '2', 'server': 'Jetty(9.4.43.v20210629)'}, 'RetryAttempts': 0}}
        return ok_response({'id': report.id}), 201

    return err_response('An exception occurred while creating the report', result), 400
