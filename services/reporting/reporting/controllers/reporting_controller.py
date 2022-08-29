from reporting.context.context import get_report_ctx
from reporting.meta.const import (
    E_REPORT_CREATE,
    E_REPORT_CREATE_INVALID_INPUT,
    E_REPORT_GET,
)
from reporting.middleware.authorize import authorize
from reporting.models.gouache_response import err_response, ok_response
from reporting.utils.deserialize import deserialize
from flask import Blueprint, request
from reporting.entities.report import Report

from werkzeug.local import LocalProxy

from reporting.utils.normalize import normalize_dynamo_report


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
            err_response(E_REPORT_GET, result),
            400,
        )

    if 'Item' in result:
        return ok_response(normalize_dynamo_report(result.get('Item'))), 200

    return ok_response(None), 404


@reporting.route('/report', methods=['POST'])
@deserialize(Report)
@authorize
def create_report(report: Report):
    db = LocalProxy(get_report_ctx)

    if report is None:
        return (
            err_response(
                E_REPORT_CREATE_INVALID_INPUT,
                f'report value: {request.get_json(silent=True)}',
            ),
            400,
        )

    result = db.put(
        caller=report.caller,
        data=report.data,
        id=report.id,
        name=report.name,
    )

    if type(result) == str:
        return (
            err_response(E_REPORT_CREATE, result),
            400,
        )

    if (
        'ResponseMetadata' in result
        and result['ResponseMetadata'].get('HTTPStatusCode') == 200
    ):
        return ok_response({'id': report.id}), 201

    return err_response(E_REPORT_CREATE, result), 400
