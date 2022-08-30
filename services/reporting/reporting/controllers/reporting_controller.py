"""
API endpoints for the reporting service.
"""
from typing import Tuple

from flask import Blueprint, Response, request
from werkzeug.local import LocalProxy

from reporting.context.context import get_report_ctx
from reporting.meta.const import (
    E_REPORT_CREATE,
    E_REPORT_CREATE_INVALID_INPUT,
    E_REPORT_GET,
    E_REPORT_GET_ALL,
)
from reporting.middleware.authorize import authorize
from reporting.models.gouache_response import err_response, ok_response
from reporting.utils.deserialize import deserialize
from reporting.utils.normalize import normalize_dynamo_report
from reporting.entities.report import Report


reporting = Blueprint(
    'reporting',
    __name__,
)


@reporting.route('/report', methods=['GET'])
@authorize
def get_all_reports() -> Tuple[Response, int]:
    """Retrieve all reports.

    Returns:
        Tuple[Response, int]: A normalized response object potentially
        containing the retrieved Reports or page thereof; an HTTP status code.
    """
    last_page_key: str | None = request.args.get('last_page_key')
    db = LocalProxy(get_report_ctx)

    result = db.get_all(last_page_key)  # type: ignore

    if isinstance(result, str):
        return (
            err_response(E_REPORT_GET_ALL, result),
            400,
        )

    if (
        'ResponseMetadata' in result
        and result['ResponseMetadata'].get('HTTPStatusCode') == 200
    ):
    # LastEvaluatedKey
    print(result)
    return ok_response(result), 200


@reporting.route('/report/<report_id>', methods=['GET'])
@authorize
def get_report(report_id: str) -> Tuple[Response, int]:
    """Retrieve a Report by its id `report_id`.

    Args:
        report_id (str): The UUID corresponding to the desired report.

    Returns:
        Tuple[Response, int]: A normalized response object potentially
        containing the retrieved Report; an HTTP status code.
    """
    db = LocalProxy(get_report_ctx)
    result = db.get(report_id)  # type: ignore

    if isinstance(result, str):
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
def create_report(report: Report) -> Tuple[Response, int]:
    """Create report endpoint. Creates a new Report in
    persistent storage and returns its system-generated UUID.

    Args:
        report (Report): The deserialized Report.
        This object is auto-generated using the request body,
        which should fulfill the public contract of the Report constructor.

    Returns:
        Tuple[Response, int]: A normalized response object potentially
        containing the id of the newly-created Report; an HTTP status code.
    """
    db = LocalProxy(get_report_ctx)

    if report is None:
        return (
            err_response(
                E_REPORT_CREATE_INVALID_INPUT,
                f'report value: {request.get_json(silent=True)}',
            ),
            400,
        )

    result = db.put(  # type: ignore
        caller=report.caller,
        data=report.data,
        report_id=report.id,
        name=report.name,
    )

    if isinstance(result, str):
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
