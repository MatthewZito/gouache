"""Data normalization utilities"""


def normalize_dynamo_report(result_item: dict):
    """Normalize a dict `result_item` as returned by
    DynamoDB into a lower-case-keyed object.

    Args:
        result_item (dict): A report result as returned by DynamoDB.

    Returns:
        _type_: @todo
    """
    return {
        'id': result_item.get('Id'),
        'data': result_item.get('Data'),
        'caller': result_item.get('Caller'),
        'name': result_item.get('Name'),
        'ts': result_item.get('TS'),
    }
